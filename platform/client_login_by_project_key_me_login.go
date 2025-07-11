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
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly logged-in Customer.
*
*	- If the Customer does not have a Cart yet, the most recently modified anonymous cart becomes the Customer's Cart.
*	- If the Customer already has a Cart, the most recently modified anonymous cart is handled in accordance with [AnonymousCartSignInMode](ctp:api:type:AnonymousCartSignInMode).
*
*	A Cart returned in the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) has any invalid Line Items removed and is [updated](/api/carts-orders-overview#update-a-cart) with the latest prices, taxes, and discounts. During these updates, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
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
