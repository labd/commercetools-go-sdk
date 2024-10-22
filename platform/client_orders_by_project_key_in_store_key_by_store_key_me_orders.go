package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Returns all Orders in a Store that match a given Query Predicate and contain either a `customerId` that matches the [customer_id:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given Query Predicate in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Orders exist in the Store that match the Query Predicate.
*	- If an Order matches the Query Predicate, but no `store` is specified, or the `store` field references a different Store.
*	- If an Order matches the Query Predicate, but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*
*	Creates an Order in a Store from a Cart for the Customer or anonymous user. The `customerId` or `anonymousId` field on the Order is automatically set based on the [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	The Cart must have a [shipping address set](ctp:api:type:CartSetShippingAddressAction) for taxes to be calculated. When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyCarts` [Permission](ctp:api:type:Permission).
*
*	If the Cart's `customerId` does not match the [customer:{id}](/scopes#composable-commerce-oauth) scope, or the `anonymousId` does not match the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
*	Creating an Order produces the [OrderCreated](ctp:api:type:OrderCreatedMessage) Message.
*
*	Specific Error Codes:
*
*	- [AssociateMissingPermission](ctp:api:type:AssociateMissingPermissionError)
*	- [CountryNotConfiguredInStore](ctp:api:type:CountryNotConfiguredInStoreError)
*	- [DiscountCodeNonApplicable](ctp:api:type:DiscountCodeNonApplicableError)
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [MatchingPriceNotFound](ctp:api:type:MatchingPriceNotFoundError)
*	- [MissingTaxRateForCountry](ctp:api:type:MissingTaxRateForCountryError)
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*	- [PriceChanged](ctp:api:type:PriceChangedError)
*	- [ShippingMethodDoesNotMatchCart](ctp:api:type:ShippingMethodDoesNotMatchCartError)
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
