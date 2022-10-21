package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordTokenRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password-token", rb.projectKey),
		client: rb.client,
	}
}
