package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyMeShoppingListsByIDRequestBuilder {
	return &ByProjectKeyMeShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Returns all ShoppingLists that match the given Query Predicate. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given Query Predicate.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Get() *ByProjectKeyMeShoppingListsRequestMethodGet {
	return &ByProjectKeyMeShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList matches the given Query Predicate. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given Query Predicate.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Head() *ByProjectKeyMeShoppingListsRequestMethodHead {
	return &ByProjectKeyMeShoppingListsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a ShoppingList for the Customer or anonymous user. The `customerId` or `anonymousId` on the ShoppingList is automatically set based on the given [customer:{id}](/scopes#composable-commerce-oauth) or [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Post(body MyShoppingListDraft) *ByProjectKeyMeShoppingListsRequestMethodPost {
	return &ByProjectKeyMeShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}
