package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [InvalidOperation](ctp:api:type:InvalidOperationError)
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*
 */
func (rb *ByProjectKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/quotes", rb.projectKey),
		client: rb.client,
	}
}
