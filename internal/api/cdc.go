package api

import (
	"context"
	"net/url"
	"strings"
)

func (c *Client) CDC(ctx context.Context, entityNames []string, changedSince string) (map[string]any, error) {
	endpoint := c.url("cdc")
	params := url.Values{}
	params.Set("entities", strings.Join(entityNames, ","))
	params.Set("changedSince", changedSince)
	endpoint = c.addMinorVersion(endpoint) + "&" + params.Encode()
	return c.get(ctx, endpoint)
}
