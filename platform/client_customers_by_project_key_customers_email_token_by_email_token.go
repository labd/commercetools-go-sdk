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

/**
*	Use this method to retrieve a global Customer's details by using the email token during their [email verification process](/../api/customers-overview#customer-email-verification).
*
 */
func (rb *ByProjectKeyCustomersEmailTokenByEmailTokenRequestBuilder) Get() *ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet {
	return &ByProjectKeyCustomersEmailTokenByEmailTokenRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/email-token=%s", rb.projectKey, rb.emailToken),
		client: rb.client,
	}
}
