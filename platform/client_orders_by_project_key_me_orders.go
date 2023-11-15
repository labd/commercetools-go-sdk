package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyMeOrdersByIDRequestBuilder {
	return &ByProjectKeyMeOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeOrdersRequestBuilder) OrderQuote() *ByProjectKeyMeOrdersQuotesRequestBuilder {
	return &ByProjectKeyMeOrdersQuotesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeOrdersRequestBuilder) Get() *ByProjectKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given Query Predicate. Returns a `200 OK` status if any Orders match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Head() *ByProjectKeyMeOrdersRequestMethodHead {
	return &ByProjectKeyMeOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
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
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}
