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
*
*	Creates an Order from a [Quote](ctp:api:type:Quote) in a [Store](ctp:api:type:Store). The referenced Quote must have the `Pending` [state](ctp:api:type:QuoteState) and must be valid (not past the `validTo` date); otherwise, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
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
