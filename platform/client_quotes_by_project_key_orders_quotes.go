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
*
*	Creates an Order from a Quote. The referenced Quote must have the `Pending` [state](ctp:api:type:QuoteState) and must be valid (not past the `validTo` date); otherwise, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [InvalidOperation](ctp:api:type:InvalidOperationError) is returned in several cases, including the following:
*	    - The referenced Shipping Method is not active.
*	    - The referenced Shipping Method is scoped to a Store that differs from the Store referenced by the [Quote](ctp:api:type:Quote).
*	    - The referenced Shipping Method is scoped to a Store, but the [Quote](ctp:api:type:Quote) does not belong to a Store.
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*
 */
func (rb *ByProjectKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/quotes", rb.projectKey),
		client: rb.client,
	}
}
