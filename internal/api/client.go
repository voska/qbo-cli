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
	"golang.org/x/oauth2"
)

const (
	BaseURLProduction = "https://quickbooks.api.intuit.com"
	BaseURLSandbox    = "https://sandbox-quickbooks.api.intuit.com"
)

type Client struct {
	httpClient   *http.Client
	baseURL      string
	realmID      string
	minorVersion int
}

func NewClient(token *oauth2.Token, realmID string, sandbox bool, minorVersion int) *Client {
	base := BaseURLProduction
	if sandbox {
		base = BaseURLSandbox
	}
	return &Client{
		httpClient:   oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token)),
		baseURL:      base,
		realmID:      realmID,
		minorVersion: minorVersion,
	}
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s/v3/company/%s/%s", c.baseURL, c.realmID, path)
}

func (c *Client) addMinorVersion(u string) string {
	if c.minorVersion > 0 {
		sep := "?"
		if strings.Contains(u, "?") {
			sep = "&"
		}
		return fmt.Sprintf("%s%sminorversion=%d", u, sep, c.minorVersion)
	}
	return u
}

func (c *Client) Query(ctx context.Context, query string) (map[string]any, error) {
	endpoint := c.url("query")
	endpoint = c.addMinorVersion(endpoint)
	endpoint += "&query=" + url.QueryEscape(query)
	return c.get(ctx, endpoint)
}

func (c *Client) Read(ctx context.Context, entity string, id string) (map[string]any, error) {
	endpoint := c.url(entity + "/" + id)
	endpoint = c.addMinorVersion(endpoint)
	return c.get(ctx, endpoint)
}

func (c *Client) Create(ctx context.Context, entity string, body []byte) (map[string]any, error) {
	endpoint := c.url(entity)
	endpoint = c.addMinorVersion(endpoint)
	return c.post(ctx, endpoint, body)
}

func (c *Client) Update(ctx context.Context, entity string, body []byte, sparse bool) (map[string]any, error) {
	endpoint := c.url(entity)
	if sparse {
		endpoint += "?operation=update"
	}
	endpoint = c.addMinorVersion(endpoint)
	return c.post(ctx, endpoint, body)
}

func (c *Client) Delete(ctx context.Context, entity string, body []byte) (map[string]any, error) {
	endpoint := c.url(entity)
	endpoint += "?operation=delete"
	endpoint = c.addMinorVersion(endpoint)
	return c.post(ctx, endpoint, body)
}

func (c *Client) Report(ctx context.Context, reportType string, params map[string]string) (map[string]any, error) {
	endpoint := c.url("reports/" + reportType)
	endpoint = c.addMinorVersion(endpoint)
	if len(params) > 0 {
		vals := url.Values{}
		for k, v := range params {
			vals.Set(k, v)
		}
		endpoint += "&" + vals.Encode()
	}
	return c.get(ctx, endpoint)
}

func (c *Client) get(ctx context.Context, endpoint string) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}
	req.Header.Set("Accept", "application/json")
	return c.do(req)
}

func (c *Client) post(ctx context.Context, endpoint string, body []byte) (map[string]any, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(string(body)))
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot build request", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

func (c *Client) do(req *http.Request) (map[string]any, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitRetryable, "request failed", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot read response", err)
	}

	if resp.StatusCode >= 400 {
		return nil, mapHTTPError(resp.StatusCode, respBody)
	}

	var result map[string]any
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, errfmt.Wrap(errfmt.ExitError, "cannot parse response", err)
	}
	return result, nil
}

func mapHTTPError(status int, body []byte) *errfmt.Error {
	detail := string(body)
	switch status {
	case 401:
		return errfmt.Auth("authentication failed — run: qbo auth login")
	case 403:
		return &errfmt.Error{Code: errfmt.ExitForbidden, Message: "permission denied", Detail: detail}
	case 404:
		return &errfmt.Error{Code: errfmt.ExitNotFound, Message: "not found", Detail: detail}
	case 429:
		return errfmt.RateLimit()
	default:
		if status >= 500 {
			return &errfmt.Error{Code: errfmt.ExitRetryable, Message: fmt.Sprintf("server error (%d)", status), Detail: detail}
		}
		return &errfmt.Error{Code: errfmt.ExitError, Message: fmt.Sprintf("API error (%d)", status), Detail: detail}
	}
}
