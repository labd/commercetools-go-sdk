package frontend

// Generated file, please do not change!!!

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Version identifies the current library version. Should match the git tag
const Version = "1.3.0"

type Client struct {
	httpClient *http.Client
	url        *url.URL
}

type ClientConfig struct {
	URL         string
	Credentials *clientcredentials.Config
	LogLevel    int
	HTTPClient  *http.Client
	UserAgent   string
}

type SetUserAgentTransport struct {
	T         http.RoundTripper
	userAgent string
}

func (sat *SetUserAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", sat.userAgent)
	if sat.T != nil {
		return sat.T.RoundTrip(req)
	}
	return http.DefaultTransport.RoundTrip(req)
}

// NewClient creates a new client based on the provided ClientConfig
func NewClient(cfg *ClientConfig) (*Client, error) {
	userAgent := cfg.UserAgent
	if userAgent == "" {
		userAgent = GetUserAgent()
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient.Transport = &SetUserAgentTransport{
		T: httpClient.Transport, userAgent: userAgent}

	if cfg.Credentials != nil {
		httpClient = cfg.Credentials.Client(
			context.WithValue(context.TODO(), oauth2.HTTPClient, httpClient))
	}

	url, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}
	client := &Client{
		url:        url,
		httpClient: httpClient,
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetUserAgent() string {
	return fmt.Sprintf("commercetools-go-sdk/%s Go/%s (%s; %s)",
		Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
