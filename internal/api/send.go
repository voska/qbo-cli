package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// SendInvoice emails an invoice once. QuickBooks does not deduplicate sends,
// so neither redirects nor retryable error codes are safe for this operation.
func (c *Client) SendInvoice(ctx context.Context, id, to string) (map[string]any, error) {
	endpoint := c.addMinorVersion(c.url("invoice/" + url.PathEscape(id) + "/send"))
	if to != "" {
		separator := "?"
		if c.minorVersion > 0 {
			separator = "&"
		}
		endpoint += separator + "sendTo=" + url.QueryEscape(to)
	}
	// A nil/NoBody POST can be replayed by Go 1.25/1.26's HTTP/2 transport.
	// This empty, non-rewindable reader sends no bytes but prevents that replay.
	// Do not set GetBody or an idempotency key on a send.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, io.NopCloser(strings.NewReader("")))
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build send request", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/octet-stream")

	// Copy the client so other commands keep their existing redirect policy.
	client := *c.httpClient
	client.CheckRedirect = func(_ *http.Request, _ []*http.Request) error { return http.ErrUseLastResponse }
	resp, err := client.Do(req)
	if err != nil {
		return nil, unknownSend(id, err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, readErr := io.ReadAll(resp.Body)
	switch resp.StatusCode {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return nil, mapHTTPError(resp.StatusCode, body)
	}
	if readErr != nil {
		return nil, unknownSend(id, readErr)
	}
	if resp.StatusCode == http.StatusRequestTimeout || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
		return nil, unknownSend(id, fmt.Errorf("HTTP %d: %s", resp.StatusCode, body))
	}
	if resp.StatusCode >= 400 {
		return nil, mapHTTPError(resp.StatusCode, body)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, unknownSend(id, fmt.Errorf("unexpected HTTP %d (redirects are not followed)", resp.StatusCode))
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, unknownSend(id, err)
	}
	invoice, _ := result["Invoice"].(map[string]any)
	delivery, _ := invoice["DeliveryInfo"].(map[string]any)
	deliveryTime, _ := delivery["DeliveryTime"].(string)
	email, _ := invoice["BillEmail"].(map[string]any)
	address, _ := email["Address"].(string)
	if invoice["Id"] != id || invoice["EmailStatus"] != "EmailSent" || deliveryTime == "" || address == "" || (to != "" && address != to) {
		return nil, unknownSend(id, fmt.Errorf("response does not verify the invoice ID, EmailStatus, BillEmail.Address and DeliveryInfo.DeliveryTime"))
	}
	return result, nil
}

func unknownSend(id string, cause error) error {
	return errfmt.Wrap(errfmt.ExitError, "send outcome unknown: read invoice "+id+" (EmailStatus, DeliveryInfo) before sending again", cause)
}
