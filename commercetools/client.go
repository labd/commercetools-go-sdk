package commercetools

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/clientcredentials"
)

type Config struct {
	ProjectKey string
	URL        string
	HTTPClient *http.Client
}

// Client bundles the logic for sending requests to the CommerceTools platform.
type Client struct {
	httpClient *http.Client
	url        string
	projectKey string
	logLevel   int
}

// New creates a new client based on the provided Config.
func New(cfg *Config) *Client {
	client := &Client{
		projectKey: getConfigValue(cfg.ProjectKey, "CTP_PROJECT_KEY"),
		url:        getConfigValue(cfg.URL, "CTP_API_URL"),
		httpClient: cfg.HTTPClient,
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
	return errors.New("NOT IMPLEMENTED")
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
func (c *Client) Update(endpoint string, queryParams url.Values, version int, actions UpdateActions, output interface{}) error {
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
	url := c.url + "/" + c.projectKey + "/" + endpoint
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

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
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
		return parseErrorResponse(resp, body)
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
