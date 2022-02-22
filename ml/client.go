// Generated file, please do not change!!!
package ml

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"net/http"
	"net/url"
	"runtime"
)

// Version identifies the current library version. Should match the git tag
const Version = "0.1.0"

type Client struct {
	httpClient *http.Client
	url        *url.URL
	logLevel   int
	userAgent  string
}

type ClientConfig struct {
	URL         string
	Credentials *clientcredentials.Config
	LogLevel    int
	HTTPClient  *http.Client
	UserAgent   string
}

// NewClient creates a new client based on the provided ClientConfig
func NewClient(cfg *ClientConfig) (*Client, error) {

	// If a custom httpClient is passed use that
	var httpClient *http.Client
	if cfg.HTTPClient != nil {
		if cfg.Credentials != nil {
			httpClient = cfg.Credentials.Client(
				context.WithValue(context.TODO(), oauth2.HTTPClient, cfg.HTTPClient))
		} else {
			httpClient = cfg.HTTPClient
		}
	} else {
		httpClient = cfg.Credentials.Client(context.TODO())
	}

	var userAgent = cfg.UserAgent
	if userAgent == "" {
		userAgent = GetUserAgent()
	}

	url, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}
	client := &Client{
		url:        url,
		logLevel:   cfg.LogLevel,
		httpClient: httpClient,
		userAgent:  userAgent,
	}

	return client, nil
}

func (c *Client) createEndpoint(p string) (*url.URL, error) {
	url, err := url.Parse(p)
	if err != nil {
		return nil, err
	}
	return c.url.ResolveReference(url), nil
}

func (c *Client) head(ctx context.Context, path string, queryParams url.Values, headers http.Header) (*http.Response, error) {
	return c.execute(ctx, http.MethodHead, path, queryParams, headers, nil)
}

func (c *Client) get(ctx context.Context, path string, queryParams url.Values, headers http.Header) (*http.Response, error) {
	return c.execute(ctx, http.MethodGet, path, queryParams, headers, nil)
}

func (c *Client) post(ctx context.Context, path string, queryParams url.Values, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, http.MethodPost, path, queryParams, headers, body)
}

func (c *Client) put(ctx context.Context, path string, queryParams url.Values, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, http.MethodPut, path, queryParams, headers, body)
}

func (c *Client) delete(ctx context.Context, path string, queryParams url.Values, headers http.Header, body io.Reader) (*http.Response, error) {
	return c.execute(ctx, http.MethodDelete, path, queryParams, headers, body)
}

func (c *Client) execute(ctx context.Context, method string, path string, params url.Values, headers http.Header, body io.Reader) (*http.Response, error) {
	endpoint, err := c.createEndpoint(path)
	if err != nil {
		return nil, err
	}

	if params != nil {
		endpoint.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), body)
	if err != nil {
		return nil, fmt.Errorf("creating new request: %w", err)
	}

	if headers != nil {
		req.Header = headers
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.userAgent)

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

func GetUserAgent() string {
	return fmt.Sprintf("commercetools-go-sdk/%s Go/%s (%s; %s)",
		Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
