package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder struct {
	projectKey    string
	passwordToken string
	client        *Client
}

func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder) Get() *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	return &ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/password-token=%s", rb.projectKey, rb.passwordToken),
		client: rb.client,
	}
}
