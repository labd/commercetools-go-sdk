// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get ShoppingList by ID
 */
func (rb *ByProjectKeyMeShoppingListsByIdRequestBuilder) Get() *ByProjectKeyMeShoppingListsByIdRequestMethodGet {
	return &ByProjectKeyMeShoppingListsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ShoppingList by ID
 */
func (rb *ByProjectKeyMeShoppingListsByIdRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyMeShoppingListsByIdRequestMethodPost {
	return &ByProjectKeyMeShoppingListsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ShoppingList by ID
 */
func (rb *ByProjectKeyMeShoppingListsByIdRequestBuilder) Delete() *ByProjectKeyMeShoppingListsByIdRequestMethodDelete {
	return &ByProjectKeyMeShoppingListsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
