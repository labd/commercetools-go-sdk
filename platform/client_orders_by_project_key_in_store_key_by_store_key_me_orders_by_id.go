package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns an Order for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists in the Store for the given `id`.
*	- If the Order exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If the Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Order exists in the Store for the given `id`.
*	- If the Order exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If the Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
