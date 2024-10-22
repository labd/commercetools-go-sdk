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

/**
*	Returns all Orders that match a given Query Predicate.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Orders exist for a given Query Predicate.
*	- If the Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Get() *ByProjectKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given Query Predicate. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists that matches the Query Predicate.
*	- If one or more Orders exist but don't have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Head() *ByProjectKeyMeOrdersRequestMethodHead {
	return &ByProjectKeyMeOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*
*	Creates an Order from a Cart for the Customer or anonymous user. The `customerId` or `anonymousId` field on the Order is automatically set based on the [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	The Cart must have a [shipping address set](ctp:api:type:CartSetShippingAddressAction) for taxes to be calculated. When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyCarts` [Permission](ctp:api:type:Permission).
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	If the Cart's `customerId` does not match the [customer:{id}](/scopes#composable-commerce-oauth) scope, or the `anonymousId` does not match the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
*	If a server-side problem occurs, indicated by a 500 Internal Server Error HTTP response, the Order creation may still successfully complete after the error is returned.
*	If you receive this error, you should verify the status of the Order by querying a unique identifier supplied during the creation request, such as the Order number.
*
*	Specific Error Codes:
*
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [ShippingMethodDoesNotMatchCart](ctp:api:type:ShippingMethodDoesNotMatchCartError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}
