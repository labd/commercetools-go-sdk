package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithCustomerId(customerId string) *ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsCustomerIdByCustomerIdRequestBuilder{
		customerId: customerId,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Replicate() *ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Carts in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Carts exist for the provided query predicate in a [Store](ctp:api:type:Store). Returns a `200 OK` status if any Carts match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a Cart in a [Store](ctp:api:type:Store).
*
*	If the referenced [ShippingMethod](ctp:api:type:ShippingMethod) in the [CartDraft](ctp:api:type:CartDraft) has a predicate that does not match, or if the Shipping Method is not active, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Specific Error Codes:
*
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/carts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
