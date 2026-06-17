package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"github.com/voska/qbo-cli/internal/errfmt"
)

// downloadTimeout bounds the fetch of a pre-signed attachment URL. It is
// generous because attachments may be up to the 100 MB upload limit.
const downloadTimeout = 10 * time.Minute

func (c *Client) Upload(ctx context.Context, metadata []byte, filename, contentType string, fileContent io.Reader) (map[string]any, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	metaHeader := make(textproto.MIMEHeader)
	metaHeader.Set("Content-Disposition", `form-data; name="file_metadata_01"; filename="metadata.json"`)
	metaHeader.Set("Content-Type", "application/json")
	metaPart, err := writer.CreatePart(metaHeader)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot create metadata part", err)
	}
	if _, err := metaPart.Write(metadata); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot write metadata", err)
	}

	fileHeader := make(textproto.MIMEHeader)
	fileHeader.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file_content_01"; filename="%s"`, filename))
	if contentType != "" {
		fileHeader.Set("Content-Type", contentType)
	}
	filePart, err := writer.CreatePart(fileHeader)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot create file part", err)
	}
	if _, err := io.Copy(filePart, fileContent); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot write file content", err)
	}

	if err := writer.Close(); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot close multipart writer", err)
	}

	endpoint := c.url("upload")
	endpoint = c.addMinorVersion(endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, &body)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	result, err := c.do(req)
	if err != nil {
		return nil, err
	}

	ar, _ := result["AttachableResponse"].([]any)
	if len(ar) == 0 {
		return nil, errfmt.New(errfmt.ExitError, "unexpected upload response: missing AttachableResponse")
	}
	first, _ := ar[0].(map[string]any)
	att, _ := first["Attachable"].(map[string]any)
	if att == nil {
		return nil, errfmt.New(errfmt.ExitError, "unexpected upload response: missing Attachable")
	}
	return att, nil
}

func (c *Client) FetchDownloadURL(ctx context.Context, attachableID string) (string, error) {
	endpoint := c.url("download/" + attachableID)
	endpoint = c.addMinorVersion(endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", errfmt.Wrap(errfmt.ExitRetryable, "request failed", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errfmt.Wrap(errfmt.ExitError, "cannot read response", err)
	}

	if resp.StatusCode >= 400 {
		return "", mapHTTPError(resp.StatusCode, respBody)
	}

	// The download endpoint returns the URL as a JSON-quoted string literal,
	// e.g. "https://...?Expires=...". Unwrap it so callers get a usable URL
	// rather than one wrapped in stray double quotes.
	trimmed := strings.TrimSpace(string(respBody))
	var url string
	if err := json.Unmarshal([]byte(trimmed), &url); err != nil {
		url = strings.Trim(trimmed, `"`)
	}
	return url, nil
}

func (c *Client) Download(ctx context.Context, attachableID string) (io.ReadCloser, error) {
	presignedURL, err := c.FetchDownloadURL(ctx, attachableID)
	if err != nil {
		return nil, err
	}
	// The pre-signed URL points at external storage (not the QBO API host) and
	// carries its own signature, so use a plain context-aware client — never the
	// OAuth client, which would leak the bearer token to a third party.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, presignedURL, nil)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}
	resp, err := (&http.Client{Timeout: downloadTimeout}).Do(req) //nolint:gosec // pre-signed download URL from QBO API
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitRetryable, "download failed", err)
	}
	if resp.StatusCode >= 400 {
		_ = resp.Body.Close()
		return nil, errfmt.New(errfmt.ExitError, fmt.Sprintf("download failed (%d)", resp.StatusCode))
	}
	return resp.Body, nil
}
