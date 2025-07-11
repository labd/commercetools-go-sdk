package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a ShoppingList with the provided `key`.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `key`. Returns a `200 OK` status if the ShoppingList exists, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Head() *ByProjectKeyShoppingListsKeyByKeyRequestMethodHead {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList in the Project using one or more [update actions](/../api/projects/shoppingLists#update-actions).
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in the Project.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
