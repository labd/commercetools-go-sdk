package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a ShoppingList with the provided `id` in a [Store](ctp:api:type:Store).
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `id` in a [Store](ctp:api:type:Store). Returns a `200 OK` status if the ShoppingList exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList in a [Store](ctp:api:type:Store) using one or more [update actions](/../api/projects/shoppingLists#update-actions).
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in a [Store](ctp:api:type:Store).
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
