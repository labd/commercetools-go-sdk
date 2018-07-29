package credentials

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
)

type ClientCredentialsProvider struct {
	clientID     string
	clientSecret string
	scope        string
	authResponse AuthResponse
	timestamp    time.Time
}

// TODO: Re-use client ?
func NewClientCredentialsProvider(clientID string, clientSecret string, scope string) AuthProvider {
	return &ClientCredentialsProvider{
		clientID:     getConfigValue(clientID, "CT_CLIENT_ID"),
		clientSecret: getConfigValue(clientSecret, "CT_CLIENT_SECRET"),
		scope:        scope,
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
