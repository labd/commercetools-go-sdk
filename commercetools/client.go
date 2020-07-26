package commercetools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Version identifies the current library version. Should match the git tag
const Version = "0.1.0"

type ClientEndpoints struct {
	Auth              string
	API               string
	MerchantCenterAPI string
}

type ClientCredentials struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
}

type ClientConfig struct {
	ProjectKey     string
	Endpoints      *ClientEndpoints
	Credentials    *ClientCredentials
	HTTPClient     *http.Client
	LibraryName    string
	LibraryVersion string
	ContactURL     string
	ContactEmail   string
}

// Config is used to pass settings for creating a new Client object
type Config struct {
	ProjectKey     string
	URL            string
	HTTPClient     *http.Client
	LibraryName    string
	LibraryVersion string
	ContactURL     string
	ContactEmail   string
}

// Client bundles the logic for sending requests to the CommerceTools platform.
type Client struct {
	httpClient *http.Client
	url        string
	endpoints  ClientEndpoints
	projectKey string
	logLevel   int
	userAgent  string
}

// NewClient creates a new client based on the provided ClientConfig
func NewClient(cfg *ClientConfig) (*Client, error) {

	if cfg.Endpoints == nil {
		cfg.Endpoints = &ClientEndpoints{
			Auth:              os.Getenv("CTP_AUTH_URL"),
			API:               os.Getenv("CTP_API_URL"),
			MerchantCenterAPI: os.Getenv("CTP_MC_API_URL"),
		}
	}

	if cfg.Credentials == nil {
		cfg.Credentials = &ClientCredentials{
			ClientID:     os.Getenv("CTP_CLIENT_ID"),
			ClientSecret: os.Getenv("CTP_CLIENT_SECRET"),
			Scopes:       strings.Split(os.Getenv("CTP_SCOPES"), ","),
		}
	}

	auth := &clientcredentials.Config{
		ClientID:     cfg.Credentials.ClientID,
		ClientSecret: cfg.Credentials.ClientSecret,
		Scopes:       cfg.Credentials.Scopes,
		TokenURL:     cfg.Endpoints.Auth,
	}

	// If a custom httpClient is passed use that
	var httpClient *http.Client
	if cfg.HTTPClient != nil {
		httpClient = auth.Client(
			context.WithValue(oauth2.NoContext, oauth2.HTTPClient, cfg.HTTPClient))
	} else {
		httpClient = auth.Client(context.TODO())
	}

	client := &Client{
		projectKey: getConfigValue(cfg.ProjectKey, "CTP_PROJECT_KEY"),
		endpoints:  *cfg.Endpoints,
		httpClient: httpClient,
		userAgent:  GetUserAgent(cfg),
	}

	if os.Getenv("CTP_DEBUG") != "" {
		client.logLevel = 1
	}
	return client, nil
}

func NewClientEndpoints(region string, provider string) *ClientEndpoints {
	return &ClientEndpoints{
		Auth:              fmt.Sprintf("https://auth.%s.%s.commercetools.com/oauth/token", region, provider),
		API:               fmt.Sprintf("https://api.%s.%s.commercetools.com", region, provider),
		MerchantCenterAPI: fmt.Sprintf("https://mc-api.%s.%s.commercetools.com", region, provider),
	}
}

// New creates a new client based on the provided Config (Deprecated)
func New(cfg *Config) *Client {

	userAgent := GetUserAgent(&ClientConfig{
		LibraryName:    cfg.LibraryName,
		LibraryVersion: cfg.LibraryVersion,
		ContactURL:     cfg.ContactURL,
		ContactEmail:   cfg.ContactEmail,
	})

	apiURL, err := cleanURL(getConfigValue(cfg.URL, "CTP_API_URL"))
	if err != nil {
		return nil
	}

	client := &Client{
		projectKey: getConfigValue(cfg.ProjectKey, "CTP_PROJECT_KEY"),
		url:        apiURL,
		httpClient: cfg.HTTPClient,
		userAgent:  userAgent,
	}

	if client.httpClient == nil {
		auth := &clientcredentials.Config{
			ClientID:     os.Getenv("CTP_CLIENT_ID"),
			ClientSecret: os.Getenv("CTP_CLIENT_SECRET"),
			Scopes:       strings.Split(os.Getenv("CTP_SCOPES"), ","),
			TokenURL:     os.Getenv("CTP_AUTH_URL"),
		}
		client.httpClient = auth.Client(context.TODO())
	}

	client.endpoints = ClientEndpoints{
		API: getConfigValue(cfg.URL, "CTP_API_URL"),
	}

	if os.Getenv("CTP_DEBUG") != "" {
		client.logLevel = 1
	}
	return client
}

func getConfigValue(value string, envName string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envName)
}

func (c *Client) Endpoints() ClientEndpoints {
	return c.endpoints
}

func (c *Client) ProjectKey() string {
	return c.projectKey
}

// Get accomodates get requests tot the CommerceTools platform.
func (c *Client) get(ctx context.Context, endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest(ctx, "GET", endpoint, queryParams, nil, output)
	return err
}

// Query accomodates query requests tot the CommerceTools platform.
func (c *Client) query(ctx context.Context, endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest(ctx, "GET", endpoint, queryParams, nil, output)
	return err
}

// Create accomodates post intended for creation requests tot the CommerceTools
// platform.
func (c *Client) create(ctx context.Context, endpoint string, queryParams url.Values, input interface{}, output interface{}) error {
	data, err := serializeInput(input)
	if err != nil {
		return err
	}
	err = c.doRequest(ctx, "POST", endpoint, queryParams, data, output)
	return err
}

// Update accomodates post requests intended for updates tot the CommerceTools
// platform.
func (c *Client) update(ctx context.Context, endpoint string, queryParams url.Values, version int, actions interface{}, output interface{}) error {
	data, err := serializeInput(&map[string]interface{}{
		"version": version,
		"actions": actions,
	})
	if err != nil {
		return err
	}
	err = c.doRequest(ctx, "POST", endpoint, queryParams, data, output)
	return err
}

// Delete accomodates delete requests tot the CommerceTools platform.
func (c *Client) delete(ctx context.Context, endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest(ctx, "DELETE", endpoint, queryParams, nil, output)
	return err
}

func (c *Client) doRequest(ctx context.Context, method string, endpoint string, params url.Values, data io.Reader, output interface{}) error {
	url := c.endpoints.API + "/" + c.projectKey + "/" + endpoint
	resp, err := c.getResponse(ctx, method, url, params, data, nil)

	if err != nil {
		return err
	}
	return processResponse(resp, output)
}

func (c *Client) getResponse(ctx context.Context, method string, url string, params url.Values, data io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, data)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new request")
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("User-Agent", c.userAgent)
	for headerName, headerValue := range headers {
		req.Header.Set(headerName, headerValue)
	}

	if c.logLevel > 0 {
		logRequest(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, handleAuthError(err)
	}

	if c.logLevel > 0 {
		logResponse(resp)
	}

	return resp, nil
}

func processResponse(resp *http.Response, output interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200, 201:
		return json.Unmarshal(body, output)
	default:
		if resp.StatusCode == 404 && len(body) == 0 {
			return ErrorResponse{
				StatusCode: resp.StatusCode,
				Message:    "Not Found (404): ResourceNotFound",
			}
		}
		customErr := ErrorResponse{}
		err = json.Unmarshal(body, &customErr)
		if err != nil {
			return err
		}
		return customErr
	}
}

func cleanURL(baseURL string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL, err
	}
	u.Path = ""
	return u.String(), nil
}

func serializeInput(input interface{}) (io.Reader, error) {
	m, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return nil, errors.Wrap(err, "Unable to serialize content")
	}
	data := bytes.NewReader(m)
	return data, nil
}

// Unnwrap the error object and return an ErrorResponse
func handleAuthError(err error) error {
	if uErr, ok := err.(*url.Error); ok {
		if rErr, ok := uErr.Err.(*oauth2.RetrieveError); ok {
			customErr := ErrorResponse{}
			if len(rErr.Body) > 0 {
				jsonErr := json.Unmarshal(rErr.Body, &customErr)
				if jsonErr != nil {
					return jsonErr
				}
			} else {
				customErr.Message = rErr.Error()
			}

			return customErr
		}
	}
	return err
}
