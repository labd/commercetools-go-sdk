package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder struct {
	projectKey string
	emailToken string
	client     *Client
}

func (rb *ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/email-token=%s", rb.projectKey, rb.emailToken),
		client: rb.client,
	}
}
