package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*
*	If used with an optional [access token for an anonymous session](ctp:api:type:AnonymousSession), all Orders and Carts that belong to the `anonymousId` are assigned to the newly created Customer.
*
*	If omitted in the request body, the [Customer](ctp:api:type:Customer) `stores` field is set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	A Cart returned in the [CustomerSignInResult](ctp:api:type:CustomerSignInResult) has any invalid Line Items removed and is [updated](/api/carts-orders-overview#update-a-cart) with the latest prices, taxes, and discounts. During these updates, the following errors can be returned: [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError) and [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError).
*
*	Creating a Customer produces the [CustomerCreated](ctp:api:type:CustomerCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestBuilder) Post(body MyCustomerDraft) *ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeSignupRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/signup", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
