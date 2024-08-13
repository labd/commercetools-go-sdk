package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeSignupRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly created Customer.
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyMeSignupRequestMethodPost {
	return &ByProjectKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/signup", rb.projectKey),
		client: rb.client,
	}
}
