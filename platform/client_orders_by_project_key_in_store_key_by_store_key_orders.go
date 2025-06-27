package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) OrderQuote() *ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) WithOrderNumber(orderNumber string) *ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersOrderNumberByOrderNumberRequestBuilder{
		orderNumber: orderNumber,
		projectKey:  rb.projectKey,
		storeKey:    rb.storeKey,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Orders in a [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Orders exist for the provided query predicate in a [Store](ctp:api:type:Store). Returns a `200 OK` status if any Orders match the query predicate, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates an Order from a Cart in a [Store](ctp:api:type:Store).
*	The Cart must have a shipping address and an active Shipping Method set.
*
*	The shipping address is used for tax calculation for a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	If a server-side problem occurs, indicated by a 500 Internal Server Error HTTP response, the Order creation may still successfully complete after the error is returned.
*	If you receive this error, you should verify the status of the Order by querying a unique identifier supplied during the creation request, such as the Order number.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [ShippingMethodDoesNotMatchCart](ctp:api:type:ShippingMethodDoesNotMatchCartError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestBuilder) Post(body OrderFromCartDraft) *ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
