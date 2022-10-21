package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeLoginRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the authenticated customer (that matches the given email/password pair).
*
*	If used with [an access token for an anonymous session](/../api/authorization#tokens-for-anonymous-sessions), all Orders and Carts that belong to the `anonymousId` are assigned to the newly logged-in Customer.
*
*	- If the Customer does not have a Cart yet, the most recently modified anonymous cart becomes the Customer's Cart.
*	- If the Customer already has a Cart, the most recently modified anonymous cart is handled in accordance with [AnonymousCartSignInMode](ctp:api:type:AnonymousCartSignInMode).
*
*	A Cart returned as part of the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) is [recalculated](ctp:api:type:Recalculate) with up-to-date prices, taxes, discounts, and invalid line items removed.
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
*
 */
func (rb *ByProjectKeyMeLoginRequestBuilder) Post(body MyCustomerSignin) *ByProjectKeyMeLoginRequestMethodPost {
	return &ByProjectKeyMeLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/login", rb.projectKey),
		client: rb.client,
	}
}
