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
*	Retrieves Orders in a [Store](ctp:api:type:Store) for the authenticated Customer or anonymous user.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Orders exist that match the provided query predicate.
*	- If an Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Orders exist for the provided query predicate in a [Store](ctp:api:type:Store) for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
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
*	Creates an Order from a Cart in a [Store](ctp:api:type:Store) for the Customer or anonymous user. The `customerId` or `anonymousId` field on the Order is automatically set based on the [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	The Cart must have a shipping address and an active Shipping Method set.
*
*	When creating [B2B Orders](/associates-overview#b2b-resources), the Customer must have the `CreateMyOrdersFromMyCarts` [Permission](ctp:api:type:Permission).
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
