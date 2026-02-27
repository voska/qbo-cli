package api

import (
	"context"
	"encoding/json"

	"github.com/voska/qbo-cli/internal/errfmt"
)

type BatchItem struct {
	BId       string `json:"bId"`
	Operation string `json:"operation,omitempty"`
	Query     string `json:"Query,omitempty"`
	Entity    string `json:"-"`
	Body      any    `json:",omitempty"`
}

func (c *Client) Batch(ctx context.Context, items []BatchItem) (map[string]any, error) {
	payload := map[string]any{
		"BatchItemRequest": items,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot marshal batch request", err)
	}
	return c.post(ctx, c.addMinorVersion(c.url("batch")), body)
}
