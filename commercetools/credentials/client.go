package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

type ClientCredentialsProvider struct {
	authURL      string
	clientID     string
	clientSecret string
	scope        string
	authResponse AuthResponse
	timestamp    time.Time
}

// TODO: Re-use client ?
func NewClientCredentials(options ...func(*ClientCredentialsProvider) error) AuthProvider {
	provider := &ClientCredentialsProvider{
		clientID:     os.Getenv("CTP_CLIENT_ID"),
		clientSecret: os.Getenv("CTP_CLIENT_SECRET"),
		scope:        os.Getenv("CTP_SCOPES"),
		authURL:      os.Getenv("CTP_AUTH_URL"),
	}

	for _, option := range options {
		option(provider)
	}
	return provider
}

func ClientID(value string) func(*ClientCredentialsProvider) error {
	return func(c *ClientCredentialsProvider) error {
		c.clientID = value
		return nil
	}
}

func ClientSecret(value string) func(*ClientCredentialsProvider) error {
	return func(c *ClientCredentialsProvider) error {
		c.clientSecret = value
		return nil
	}
}

func Scope(value string) func(*ClientCredentialsProvider) error {
	return func(c *ClientCredentialsProvider) error {
		c.scope = value
		return nil
	}
}

func AuthURL(value string) func(*ClientCredentialsProvider) error {
	return func(c *ClientCredentialsProvider) error {
		c.authURL = value
		return nil
	}
}

func (c *ClientCredentialsProvider) GetAuthToken() (string, error) {
	if c.IsExpired() {
		err := c.RefreshToken()
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%s %s", c.authResponse.TokenType, c.authResponse.AccessToken), nil
}

func (c *ClientCredentialsProvider) IsExpired() bool {
	if c.authResponse.ExpiresIn == 0 {
		return true
	}
	return false
}

func (c *ClientCredentialsProvider) RefreshToken() error {
	u, err := url.Parse("https://auth.sphere.io/oauth/token")
	if err != nil {
		log.Fatal("Error")
	}

	query := u.Query()
	query.Set("grant_type", "client_credentials")
	query.Set("scope", c.scope)
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("POST", u.String(), nil)
	req.SetBasicAuth(c.clientID, c.clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	client := cleanhttp.DefaultClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// TODO: Improve error handling
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return errors.New(string(body))
	}
	err = json.Unmarshal(body, &c.authResponse)
	if err != nil {
		return err
	}

	c.timestamp = time.Now()
	return nil
}
