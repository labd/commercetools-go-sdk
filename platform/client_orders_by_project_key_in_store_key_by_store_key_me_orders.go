package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given Query Predicate. Returns a `200 OK` status if any Orders match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	The Cart must have a [shipping address set](ctp:api:type:CartSetShippingAddressAction) for taxes to be calculated. When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyCarts` [Permission](ctp:api:type:Permission).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
