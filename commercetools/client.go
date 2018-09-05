package commercetools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/labd/commercetools-go-sdk/commercetools/credentials"
)

// Client bundles the logic for sending requests to the CommerceTools platform.
type Client struct {
	HTTPClient *http.Client
	apiURL     string
	projectKey string
	auth       credentials.AuthProvider
	logLevel   int
}

// NewClient creates a new client based on the provided Config.
func NewClient(options ...func(*Client) error) (*Client, error) {
	client := &Client{
		auth:       credentials.NewClientCredentials(),
		projectKey: os.Getenv("CTP_PROJECT_KEY"),
		apiURL:     os.Getenv("CTP_API_URL"),
		HTTPClient: cleanhttp.DefaultClient(),
	}

	if os.Getenv("CTP_DEBUG") != "" {
		client.logLevel = 1
	}

	// Apply the functional options
	for _, option := range options {
		option(client)
	}

	return client, nil
}

func Debug(value bool) func(*Client) error {
	return func(c *Client) error {
		if value {
			c.logLevel = 1
		} else {
			c.logLevel = 0
		}
		return nil
	}
}

func ProjectKey(value string) func(*Client) error {
	return func(c *Client) error {
		c.projectKey = value
		return nil
	}
}

func ApiURL(value string) func(*Client) error {
	return func(c *Client) error {
		c.apiURL = value
		return nil
	}
}

func AuthProvider(value credentials.AuthProvider) func(*Client) error {
	return func(c *Client) error {
		c.auth = value
		return nil
	}
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
	authToken, err := c.auth.GetAuthToken()
	if err != nil {
		return err
	}

	url := c.apiURL + "/" + c.projectKey + "/" + endpoint
	req, err := http.NewRequest(method, url, data)

	req.Header.Add("Authorization", authToken)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	if c.logLevel > 0 {
		logRequest(req)
	}

	resp, err := c.HTTPClient.Do(req)
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
		return processErrorResponse(resp.StatusCode, body)
	}
}

func processErrorResponse(statusCode int, body []byte) error {
	data := make(map[string]interface{})
	err := json.Unmarshal(body, &data)
	st := http.StatusText(statusCode)
	if err == nil {
		if val, ok := data["message"]; ok {
			return fmt.Errorf("HTTP %s: %s", st, val.(string))
		}
		return fmt.Errorf("HTTP %s: %s", st, string(body))
	}
	return fmt.Errorf("HTTP %s: %v", st, err)
}

func serializeInput(input interface{}) (io.Reader, error) {
	m, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return nil, err
	}
	data := bytes.NewReader(m)
	return data, nil
}
