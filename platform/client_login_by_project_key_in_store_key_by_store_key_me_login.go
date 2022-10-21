package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Retrieves the authenticated Customer (that matches the given email/password pair) if they are part of a specific [Store](ctp:api:type:Store).
*
*	- If the Customer does not have a Cart, the most recently modified anonymous cart becomes the Customer's Cart.
*	- If the Customer already has a Cart, the most recently modified anonymous cart is handled according to [AnonymousCartSignInMode](ctp:api:type:AnonymousCartSignInMode).
*
*	If a Cart is returned as part of [CustomerSignInResult](ctp:api:type:CustomerSignInResult), it has been [recalculated](/../api/projects/carts#recalculate) with up-to-date prices, taxes, discounts, and invalid line items removed.
*
*	If an account with the given credentials is not found, an [InvalidCredentials](ctp:api:type:InvalidCredentialsError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestBuilder) Post(body MyCustomerSignin) *ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeLoginRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/login", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
