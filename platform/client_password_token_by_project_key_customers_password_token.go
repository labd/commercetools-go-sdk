package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersPasswordTokenRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Use this method to create a password reset token for a global Customer during their [password reset process](/../api/customers-overview#customer-password-reset).
*
*	Creating a password reset token for the Customer produces the [CustomerPasswordTokenCreated](ctp:api:type:CustomerPasswordTokenCreatedMessage) Message.
*	The Message will include the token's value, if the token's validity is 60 minutes or less.
*
 */
func (rb *ByProjectKeyCustomersPasswordTokenRequestBuilder) Post(body CustomerCreatePasswordResetToken) *ByProjectKeyCustomersPasswordTokenRequestMethodPost {
	return &ByProjectKeyCustomersPasswordTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/password-token", rb.projectKey),
		client: rb.client,
	}
}
