package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// VoidBody builds the minimal request for the entity's void operation.
func (e EntityInfo) VoidBody(id, syncToken string) ([]byte, error) {
	if e.VoidStyle == "" {
		return nil, errfmt.Usage(e.Name + " cannot be voided")
	}
	body := map[string]any{"Id": id, "SyncToken": syncToken}
	if e.VoidStyle == VoidSparseUpdate {
		body["sparse"] = true
	}
	return json.Marshal(body)
}

// Void reads the current token before voiding, optionally enforcing an expected
// token. The boolean reports an already-void record, for which no POST is made.
func (c *Client) Void(ctx context.Context, entity EntityInfo, id, expectedToken string) (map[string]any, bool, error) {
	if entity.VoidStyle == "" {
		return nil, false, errfmt.Usage(entity.Name + " cannot be voided")
	}
	current, err := c.Read(ctx, entity.Endpoint, id)
	if err != nil {
		return nil, false, err
	}
	record, ok := current[entity.Name].(map[string]any)
	if !ok || record["Id"] != id {
		return nil, false, errfmt.New(errfmt.ExitError, "read response does not contain "+entity.Name+" "+id)
	}
	syncToken, ok := record["SyncToken"].(string)
	if !ok || syncToken == "" {
		return nil, false, errfmt.New(errfmt.ExitError, "read response is missing a valid SyncToken")
	}
	if expectedToken != "" && syncToken != expectedToken {
		return nil, false, errfmt.New(errfmt.ExitError, fmt.Sprintf("SyncToken mismatch: expected %s, got %s; no void submitted", expectedToken, syncToken))
	}
	if isVoided(record) {
		return current, true, nil
	}
	body, err := entity.VoidBody(id, syncToken)
	if err != nil {
		return nil, false, err
	}
	endpoint := c.addMinorVersion(c.url(entity.Endpoint + "?" + string(entity.VoidStyle)))
	result, err := c.post(ctx, endpoint, body)
	if err != nil {
		return nil, false, err
	}
	voided, ok := result[entity.Name].(map[string]any)
	if !ok || voided["Id"] != id || !isVoided(voided) {
		return nil, false, errfmt.New(errfmt.ExitError, "void submitted, unverified: response does not show "+entity.Name+" "+id+" as voided; read the record before retrying")
	}
	return result, false, nil
}

func isVoided(record map[string]any) bool {
	total, ok := record["TotalAmt"].(float64)
	note, _ := record["PrivateNote"].(string)
	return ok && total == 0 && strings.HasPrefix(note, "Voided")
}
