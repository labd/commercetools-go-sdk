// Generated file, please do not change!!!
package importapi

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	url        string
	logLevel   int
}

type ClientConfig struct {
	URL         string
	Credentials *clientcredentials.Config
	LogLevel    int
	HTTPClient  *http.Client
}

// NewClient creates a new client based on the provided ClientConfig
func NewClient(cfg *ClientConfig) (*Client, error) {

	// If a custom httpClient is passed use that
	var httpClient *http.Client
	if cfg.HTTPClient != nil {
		httpClient = cfg.Credentials.Client(
			context.WithValue(oauth2.NoContext, oauth2.HTTPClient, cfg.HTTPClient))
	} else {
		httpClient = cfg.Credentials.Client(context.TODO())
	}

	client := &Client{
		url:        cfg.URL,
		logLevel:   cfg.LogLevel,
		httpClient: httpClient,
	}

	return client, nil
}

func (c *Client) get(ctx context.Context, url string, queryParams url.Values) (*http.Response, error) {
	return c.execute(ctx, "GET", c.url+url, queryParams, nil, nil)
}

func (c *Client) post(ctx context.Context, url string, queryParams url.Values, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, "POST", c.url+url, queryParams, body, nil)
}

func (c *Client) delete(ctx context.Context, url string, queryParams url.Values, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, "DELETE", c.url+url, queryParams, body, nil)
}

func (c *Client) execute(ctx context.Context, method string, url string, params url.Values, data io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Creating new request: %w", err)
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for headerName, headerValue := range headers {
		req.Header.Set(headerName, headerValue)
	}

	if c.logLevel > 0 {
		logRequest(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c.logLevel > 0 {
		logResponse(resp)
	}

	return resp, nil
}
