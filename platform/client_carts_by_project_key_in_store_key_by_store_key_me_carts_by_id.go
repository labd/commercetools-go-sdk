package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns a Cart for a given `id` in a Store. Returns a `200 OK` status if the Cart exists.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists in the Store for the given `id`.
*	- If the Cart exists but does not belong to a Store, or the Cart's `store` field references a different Store.
*	- If the Cart exists but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates the Cart for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists in the Store for the given `id`.
*	- If the Cart exists but does not belong to a Store, or the Cart's `store` field references a different Store.
*	- If the Cart exists but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes the Cart for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists in the Store for the given `id`.
*	- If the Cart exists in the Project but does not belong to a Store, or the Cart's `store` field references a different Store.
*	- If the Cart exists in the Project but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
