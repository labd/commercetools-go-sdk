package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Use this method to verify a global Customer's email during their [email verification process](/../api/customers-overview#customer-email-verification).
*
*	Verifying the email of the Customer produces the [CustomerEmailVerified](ctp:api:type:CustomerEmailVerifiedMessage) Message.
*
*	After the email is verified, all email tokens issued previously through the [email verification flow](/../api/projects/customers#email-verification-of-customer) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
 */
func (rb *ByProjectKeyCustomersEmailConfirmRequestBuilder) Post(body CustomerEmailVerify) *ByProjectKeyCustomersEmailConfirmRequestMethodPost {
	return &ByProjectKeyCustomersEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
