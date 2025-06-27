package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersEmailTokenRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Use this method to create an email token for a global Customer during their [email verification process](/../api/customers-overview#customer-email-verification).
*
*	Creating an email token for the Customer produces the [CustomerEmailTokenCreated](ctp:api:type:CustomerEmailTokenCreatedMessage) Message.
*	The Message will include the token's value, if the token's validity is 60 minutes or less.
*
 */
func (rb *ByProjectKeyCustomersEmailTokenRequestBuilder) Post(body CustomerCreateEmailToken) *ByProjectKeyCustomersEmailTokenRequestMethodPost {
	return &ByProjectKeyCustomersEmailTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/email-token", rb.projectKey),
		client: rb.client,
	}
}
