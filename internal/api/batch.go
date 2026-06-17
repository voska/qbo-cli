package api

import (
	"context"
	"encoding/json"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// Batch posts a QBO batch request. items are the raw BatchItemRequest entries
// (each in QBO's native batch format, e.g. {"bId":"1","operation":"create","Invoice":{...}}
// or {"bId":"2","Query":"select * from Invoice"}), passed through verbatim
// under the BatchItemRequest envelope.
func (c *Client) Batch(ctx context.Context, items []json.RawMessage) (map[string]any, error) {
	payload := map[string]any{
		"BatchItemRequest": items,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot marshal batch request", err)
	}
	return c.post(ctx, c.addMinorVersion(c.url("batch")), body)
}
