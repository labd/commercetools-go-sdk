package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShoppingListsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a ShoppingList with the provided `id`.
 */
func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Get() *ByProjectKeyShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists with the provided `id`. Returns a `200 OK` status if the ShoppingList exists, or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Head() *ByProjectKeyShoppingListsByIDRequestMethodHead {
	return &ByProjectKeyShoppingListsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a ShoppingList in the Project using one or more [update actions](/../api/projects/shoppingLists#update-actions).
 */
func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a ShoppingList in the Project.
 */
func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
