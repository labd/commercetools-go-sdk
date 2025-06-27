package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves ShoppingLists for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store). Returns `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingLists exist in a Store.
*	- If a ShoppingList exists but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList exists in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more ShoppingLists exist for the provided query predicate for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store). Returns `200 OK` status if successful.
*
*	A [Not Found](/../api/errors#404-not-found) error is returned in the following scenarios:
*
*	- If no ShoppingLists exist for the provided query predicate in a Store.
*	- If a ShoppingList matches the query predicate but does not have a `store` specified, or the `store` field references a different Store.
*	- If a ShoppingList exists in a Store but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*
*	Creates a ShoppingList for the authenticated Customer or anonymous user in a [Store](ctp:api:type:Store). The `customer` or `anonymousId` field on the ShoppingList is automatically set based on the given [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
*	When using this endpoint, the `store` field of a ShoppingList is always set to the Store specified in the path parameter.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Post(body MyShoppingListDraft) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
