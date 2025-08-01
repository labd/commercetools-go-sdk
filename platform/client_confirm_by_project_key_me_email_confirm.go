package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeEmailConfirmRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	This is the last step in the [email verification process of a Customer](/../api/projects/customers#email-verification-of-customer).
*
*	After the email is verified, all email tokens issued previously through the [email verification flow](/../api/projects/customers#email-verification-of-customer) are invalidated. This invalidation of tokens is [eventually consistent](/../api/general-concepts#eventual-consistency).
*
 */
func (rb *ByProjectKeyMeEmailConfirmRequestBuilder) Post(body MyCustomerEmailVerify) *ByProjectKeyMeEmailConfirmRequestMethodPost {
	return &ByProjectKeyMeEmailConfirmRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/email/confirm", rb.projectKey),
		client: rb.client,
	}
}
