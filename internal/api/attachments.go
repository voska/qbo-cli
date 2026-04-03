package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/voska/qbo-cli/internal/errfmt"
)

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

	return strings.TrimSpace(string(respBody)), nil
}

func (c *Client) Download(ctx context.Context, attachableID string) (io.ReadCloser, error) {
	presignedQBOAttachableURL, err := c.FetchDownloadURL(ctx, attachableID)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(presignedQBOAttachableURL) //nolint:gosec // pre-signed download URL from QBO API
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitRetryable, "download failed", err)
	}
	if resp.StatusCode >= 400 {
		_ = resp.Body.Close()
		return nil, errfmt.New(errfmt.ExitError, fmt.Sprintf("download failed (%d)", resp.StatusCode))
	}
	return resp.Body, nil
}
