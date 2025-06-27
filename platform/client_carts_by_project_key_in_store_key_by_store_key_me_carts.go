package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Carts for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for a Store that matches the given Query Predicate, and contains a matching `customerId` or `anonymousId`. Returns a `200 OK` status if any Carts match these conditions, or [Not Found](/../api/errors#404-not-found) otherwise.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*
*	Creates a Cart in a Store for the Customer or anonymous user. The `customerId` or `anonymousId` field on the Cart is automatically set based on the [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	The `store` field in the created [Cart](ctp:api:type:Cart) is set to the Store specified by the `storeKey` path parameter.
*
*	Specific Error Codes:
*
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestBuilder) Post(body MyCartDraft) *ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
