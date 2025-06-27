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
*
*	Retrieves the authenticated Customer (that matches the given email/password pair) if they are part of a specific [Store](ctp:api:type:Store).
*
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly logged-in Customer.
*
*	- If the Customer does not have a Cart, the most recently modified anonymous cart becomes the Customer's Cart.
*	- If the Customer already has a Cart, the most recently modified anonymous cart is handled according to [AnonymousCartSignInMode](ctp:api:type:AnonymousCartSignInMode).
*
*	A Cart returned in the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) has any invalid Line Items removed and is [updated](/api/carts-orders-overview#update-a-cart) with the latest prices, taxes, and discounts. During these updates, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
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
