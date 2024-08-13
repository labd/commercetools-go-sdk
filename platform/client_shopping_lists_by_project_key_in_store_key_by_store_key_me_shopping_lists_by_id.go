package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Returns a ShoppingList for a given `id` in a Store. Returns `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the given `id` in a Store.
*	- If a ShoppingList matches the given `id` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the given `id` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the given `id` in a Store.
*	- If a ShoppingList matches the given `id` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the given `id` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the given `id` in a Store.
*	- If a ShoppingList matches the given `id` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the given `id` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes the ShoppingList for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the given `id` in a Store.
*	- If a ShoppingList matches the given `id` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the given `id` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#anonymous_idid) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#customer_idid) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
