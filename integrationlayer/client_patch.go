package integrationlayer

// Not generated: copied in by the generate-integrationlayer task. The
// generator writes PATCH methods but no patch helper on the client, since no
// other commercetools API has a PATCH route.

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) patch(ctx context.Context, path string, queryParams url.Values, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, http.MethodPatch, path, queryParams, headers, body)
}
