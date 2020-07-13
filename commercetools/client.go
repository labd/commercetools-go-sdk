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
	"strconv"
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

// QueryInput provides the data required to query types.
type QueryInput struct {
	// The queryable APIs support ad-hoc filtering of resources through flexible
	// predicates. They do so via the where query parameter that accepts a
	// predicate expression to determine whether a specific resource
	// representation should be included in the result. The structure of
	// predicates and the names of the fields follow the structure and naming of
	// the fields in the documented response representation of the query
	// results.
	// https://docs.commercetools.com/http-api-query-predicates.html
	Where string

	// A query endpoint that supports sorting does so through the sort query
	// parameter. The provided value must be a valid sort expression. The
	// default sort direction is ASC. The allowed sort paths are typically
	// listed on the specific query endpoints. If multiple sort expressions are
	// specified via multiple sort parameters, they are combined into a composed
	// sort where the results are first sorted by the first expression, followed
	// by equal values being sorted according to the second expression, and so
	// on.
	// https://docs.commercetools.com/http-api.html#sorting
	Sort []string

	// Reference expansion is a feature of the resources listed in the table
	// below that enables clients to request server-side expansion of Reference
	// resources, thereby reducing the number of required client-server
	// roundtrips to obtain the data that a client needs for a specific
	// use-case. Reference expansion can be used when creating, updating,
	// querying, and deleting these resources.
	// https://docs.commercetools.com/http-api.html#reference-expansion
	Expand string

	Limit  int
	Offset int
}

func (qi QueryInput) toParams() (values url.Values) {
	values = url.Values{}

	if qi.Where != "" {
		values.Set("where", qi.Where)
	}

	for i := range qi.Sort {
		values.Add("sort", qi.Sort[i])
	}

	if qi.Expand != "" {
		values.Set("expand", qi.Expand)
	}

	if qi.Limit != 0 {
		values.Set("limit", strconv.Itoa(qi.Limit))
	}

	if qi.Offset != 0 {
		values.Set("offset", strconv.Itoa(qi.Offset))
	}

	return
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

	client := &Client{
		projectKey: getConfigValue(cfg.ProjectKey, "CTP_PROJECT_KEY"),
		endpoints:  *cfg.Endpoints,
		httpClient: auth.Client(context.TODO()),
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

	client := &Client{
		projectKey: getConfigValue(cfg.ProjectKey, "CTP_PROJECT_KEY"),
		url:        getConfigValue(cfg.URL, "CTP_API_URL"),
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

// Get accomodates get requests tot the CommerceTools platform.
func (c *Client) Get(endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest("GET", endpoint, queryParams, nil, output)
	return err
}

// Query accomodates query requests tot the CommerceTools platform.
func (c *Client) Query(endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest("GET", endpoint, queryParams, nil, output)
	return err
}

// Create accomodates post intended for creation requests tot the CommerceTools
// platform.
func (c *Client) Create(endpoint string, queryParams url.Values, input interface{}, output interface{}) error {
	data, err := serializeInput(input)
	if err != nil {
		return err
	}
	err = c.doRequest("POST", endpoint, queryParams, data, output)
	return err
}

// Update accomodates post requests intended for updates tot the CommerceTools
// platform.
func (c *Client) Update(endpoint string, queryParams url.Values, version int, actions interface{}, output interface{}) error {
	data, err := serializeInput(&map[string]interface{}{
		"version": version,
		"actions": actions,
	})
	if err != nil {
		return err
	}
	err = c.doRequest("POST", endpoint, queryParams, data, output)
	return err
}

// Delete accomodates delete requests tot the CommerceTools platform.
func (c *Client) Delete(endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest("DELETE", endpoint, queryParams, nil, output)
	return err
}

func (c *Client) doRequest(method string, endpoint string, params url.Values, data io.Reader, output interface{}) error {
	url := c.endpoints.API + "/" + c.projectKey + "/" + endpoint
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return errors.Wrap(err, "Creating new request")
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	if c.logLevel > 0 {
		logRequest(req)
	}
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return handleAuthError(err)
	}
	defer resp.Body.Close()

	if c.logLevel > 0 {
		logResponse(resp)
	}

	body, err := ioutil.ReadAll(resp.Body)

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
