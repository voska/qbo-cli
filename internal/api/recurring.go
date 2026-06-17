package api

import (
	"context"
	"encoding/json"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// DeleteRecurringTransaction deletes a recurring transaction template.
//
// Unlike flat entities, QBO's recurringtransaction delete does not accept a
// simple {Id, SyncToken} body. It requires the full type-wrapped object (e.g.
// {"Invoice": {Id, SyncToken, RecurringInfo, ...}}) echoed back, matching the
// shape returned by a read — so we read the current template and POST it back
// with ?operation=delete. This mirrors Intuit's own SDKs.
func (c *Client) DeleteRecurringTransaction(ctx context.Context, id string) (map[string]any, error) {
	current, err := c.Read(ctx, "recurringtransaction", id)
	if err != nil {
		return nil, err
	}
	body, err := recurringTxnBody(current)
	if err != nil {
		return nil, err
	}
	return c.Delete(ctx, "recurringtransaction", body)
}

// recurringTxnBody extracts the bare type wrapper ({"<TxnType>": {...}}) from a
// recurringtransaction read response, dropping the response envelope (the
// "time" field and an optional "RecurringTransaction" wrapper) so the result
// can be POSTed back as the delete body.
func recurringTxnBody(read map[string]any) ([]byte, error) {
	obj := make(map[string]any, len(read))
	for k, v := range read {
		if k == "time" {
			continue
		}
		obj[k] = v
	}
	if rt, ok := obj["RecurringTransaction"].(map[string]any); ok && len(obj) == 1 {
		obj = rt
	}
	if len(obj) == 0 {
		return nil, errfmt.New(errfmt.ExitError, "could not extract recurring transaction body from read response")
	}
	return json.Marshal(obj)
}
