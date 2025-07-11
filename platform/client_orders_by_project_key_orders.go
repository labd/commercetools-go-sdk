package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersRequestBuilder) ImportOrder() *ByProjectKeyOrdersImportRequestBuilder {
	return &ByProjectKeyOrdersImportRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) OrderQuote() *ByProjectKeyOrdersQuotesRequestBuilder {
	return &ByProjectKeyOrdersQuotesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) WithOrderNumber(orderNumber string) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder{
		orderNumber: orderNumber,
		projectKey:  rb.projectKey,
		client:      rb.client,
	}
}

/**
*	OrderEdit are containers for financial changes after an Order has been placed.
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Edits() *ByProjectKeyOrdersEditsRequestBuilder {
	return &ByProjectKeyOrdersEditsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) WithId(id string) *ByProjectKeyOrdersByIDRequestBuilder {
	return &ByProjectKeyOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	This endpoint provides high performance search queries over Orders. The order search allows searching through all orders (currently supporting a limit of the 10.000.000 newest orders) in your project.
*
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Search() *ByProjectKeyOrdersSearchRequestBuilder {
	return &ByProjectKeyOrdersSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Orders in the Project.
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Get() *ByProjectKeyOrdersRequestMethodGet {
	return &ByProjectKeyOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Orders exist for the provided query predicate. Returns a `200 OK` status if any Orders match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Head() *ByProjectKeyOrdersRequestMethodHead {
	return &ByProjectKeyOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates an Order from a Cart.
*
*	The Cart must have a shipping address and an active Shipping Method set.
*
*	The shipping address is used for tax calculation for a Cart with `Platform` [TaxMode](ctp:api:type:TaxMode).
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message. If the Order is created from a Recurring Order schedule, the [OrderCreatedFromRecurringOrder](ctp:api:type:OrderCreatedFromRecurringOrderMessage) Message is generated.
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
*
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Post(body OrderFromCartDraft) *ByProjectKeyOrdersRequestMethodPost {
	return &ByProjectKeyOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders", rb.projectKey),
		client: rb.client,
	}
}
