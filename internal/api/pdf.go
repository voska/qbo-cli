package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// maxPDFSize bounds a rendered document. Real invoices are tens of kilobytes;
// the cap only stops a runaway response from filling memory.
const maxPDFSize = 50 << 20

// pdfEntities are the transaction types QuickBooks renders through the
// /{entity}/{id}/pdf endpoint. Anything else answers with an error page.
var pdfEntities = map[string]bool{
	"invoice":       true,
	"estimate":      true,
	"salesreceipt":  true,
	"creditmemo":    true,
	"purchaseorder": true,
}

// SupportsPDF reports whether QuickBooks can render the entity endpoint as a PDF.
func SupportsPDF(endpoint string) bool { return pdfEntities[endpoint] }

// PDF fetches a transaction as the PDF QuickBooks renders for it — the same
// document the customer receives when it is emailed from QuickBooks.
//
// The endpoint must be asked for application/pdf explicitly: with the
// application/json Accept header every other call sends, QuickBooks answers
// 500 "No match for accept header".
func (c *Client) PDF(ctx context.Context, endpoint, id string) ([]byte, error) {
	u := c.addMinorVersion(c.url(endpoint + "/" + id + "/pdf"))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}
	req.Header.Set("Accept", "application/pdf")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitRetryable, "request failed", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxPDFSize+1))
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot read response", err)
	}
	if resp.StatusCode >= 400 {
		return nil, mapHTTPError(resp.StatusCode, body)
	}
	if len(body) > maxPDFSize {
		return nil, errfmt.New(errfmt.ExitError, fmt.Sprintf("PDF exceeds %d byte limit", maxPDFSize))
	}
	// A 200 that is not a PDF (an XML or JSON fault) must not be saved under a
	// .pdf name, where it would look like a document that fails to open.
	if !bytes.HasPrefix(body, []byte("%PDF")) {
		snippet := body
		if len(snippet) > 200 {
			snippet = snippet[:200]
		}
		return nil, &errfmt.Error{Code: errfmt.ExitError, Message: "response is not a PDF", Detail: string(snippet)}
	}
	return body, nil
}
