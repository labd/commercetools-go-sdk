package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/orders/quotes", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
