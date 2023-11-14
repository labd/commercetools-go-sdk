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
*	When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyQuotes` [Permission](ctp:api:type:Permission).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
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
