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

/**
*	Use this method to retrieve the details of a global Customer by using the password token during their [password reset process](/../api/customers-overview#customer-password-reset).
*
 */
func (rb *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestBuilder) Get() *ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet {
	return &ByProjectKeyCustomersPasswordTokenByPasswordTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/password-token=%s", rb.projectKey, rb.passwordToken),
		client: rb.client,
	}
}
