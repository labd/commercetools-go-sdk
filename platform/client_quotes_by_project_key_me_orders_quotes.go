package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*
*	Creates an Order from a [Quote](ctp:api:type:Quote) for the authenticated Customer. To create [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyQuotes` [Permission](ctp:api:type:Permission).
*
*	The referenced Quote must have the `Pending` [state](ctp:api:type:QuoteState) and must be valid (not past the `validTo` date); otherwise, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*
 */
func (rb *ByProjectKeyMeOrdersQuotesRequestBuilder) Post(body MyOrderFromQuoteDraft) *ByProjectKeyMeOrdersQuotesRequestMethodPost {
	return &ByProjectKeyMeOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders/quotes", rb.projectKey),
		client: rb.client,
	}
}
