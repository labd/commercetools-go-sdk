package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Retrieves a ShoppingList with the provided `key` for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store). Returns `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the provided `key` in a Store.
*	- If a ShoppingList matches the provided `key` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the provided `key` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope,
*	   or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `key` for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store). Returns a `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no ShoppingList exists that matches the provided `key` in a Store.
*	- If a ShoppingList matches the provided `key` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the provided `key` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/me-shoppingLists#update-actions). Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the provided `key` in a Store.
*	- If a ShoppingList matches the provided `key` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the provided `key` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in a [Store](ctp:api:type:Store). Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList matches the provided `key` in a Store.
*	- If a ShoppingList matches the provided `key` but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList matches the provided `key` in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
