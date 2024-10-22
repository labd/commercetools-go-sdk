package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns a ShoppingList for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given `id`.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Get() *ByProjectKeyMeShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given `id`.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Head() *ByProjectKeyMeShoppingListsByIDRequestMethodHead {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given `id`.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyMeShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes the ShoppingList for a given `id`. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no ShoppingList exists for the given `id`.
*	- If a ShoppingList exists but does not contain either an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope, or a `customer` with `id` value that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
