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

	"github.com/hashicorp/go-cleanhttp"
	"github.com/labd/commercetools-go-sdk/commercetools/credentials"
)

type Client struct {
	HTTPClient *http.Client
	region     string
	projectKey string
	auth       credentials.AuthProvider
	logLevel   int
}

type Config struct {
	ProjectKey   string
	AuthProvider credentials.AuthProvider
	Region       string
	LogLevel     int
}

func NewClient(cfg *Config) (*Client, error) {

	client := &Client{
		auth:       cfg.AuthProvider,
		projectKey: getConfigValue(cfg.ProjectKey, "CT_PROJECT_KEY"),
		region:     getConfigValue(cfg.Region, "CT_REGION"),
		HTTPClient: cleanhttp.DefaultClient(),
		logLevel:   cfg.LogLevel,
	}

	if client.projectKey == "" {
		return nil, errors.New("Missing ProjectKey")
	}
	if client.region == "" {
		return nil, errors.New("Missing Region")
	}

	if os.Getenv("CT_DEBUG") != "" {
		client.logLevel = 1

	}
	return client, nil
}

func getConfigValue(value string, envName string) string {
	if value != "" {
		return value
	}
	return os.Getenv(envName)
}

func (c *Client) CreateURL(endpoint string) string {
	return c.region + "/" + c.projectKey + "/" + endpoint
}

func (c *Client) Get(endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest("GET", endpoint, queryParams, nil, output)
	return err
}

func (c *Client) Query(endpoint string, queryParams url.Values, output interface{}) error {
	return errors.New("NOT IMPLEMENTED")
}

func (c *Client) Create(endpoint string, queryParams url.Values, input interface{}, output interface{}) error {
	data, err := serializeInput(input)
	if err != nil {
		return err
	}
	err = c.doRequest("POST", endpoint, queryParams, data, output)
	return err
}

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

func (c *Client) Delete(endpoint string, queryParams url.Values, output interface{}) error {
	err := c.doRequest("DELETE", endpoint, queryParams, nil, output)
	return err
}

func (c *Client) doRequest(method string, endpoint string, params url.Values, data io.Reader, output interface{}) error {
	authToken, err := c.auth.GetAuthToken()
	if err != nil {
		return err
	}

	url := c.CreateURL(endpoint)
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
