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
*	Retrieves an Order with the provided `id` in a [Store](ctp:api:type:Store) for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Orders exists in the Store with the provided `id`.
*	- If an Order exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If an Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `id` in a [Store](ctp:api:type:Store) for the authenticated Customer or anonymous user. Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no Order exists in the Store with the provided `id`.
*	- If an Order exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If an Order exists but does not have a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/orders/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
