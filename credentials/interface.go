package credentials

import (
	"log"
	"os"
)

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type AuthProvider interface {
	GetAuthToken() (string, error)
}

// DUPLICATE FIXME
func getConfigValue(value string, envName string) string {
	log.Println(value)
	if value != "" {
		return value
	}
	return os.Getenv(envName)
}
