// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyShoppingListsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Gets a shopping list by ID.
 */
func (rb *ByProjectKeyShoppingListsByIdRequestBuilder) Get() *ByProjectKeyShoppingListsByIdRequestMethodGet {
	return &ByProjectKeyShoppingListsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ShoppingList by ID
 */
func (rb *ByProjectKeyShoppingListsByIdRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyShoppingListsByIdRequestMethodPost {
	return &ByProjectKeyShoppingListsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ShoppingList by ID
 */
func (rb *ByProjectKeyShoppingListsByIdRequestBuilder) Delete() *ByProjectKeyShoppingListsByIdRequestMethodDelete {
	return &ByProjectKeyShoppingListsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
