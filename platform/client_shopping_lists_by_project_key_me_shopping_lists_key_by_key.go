// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get ShoppingList by key
 */
func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update ShoppingList by key
 */
func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete ShoppingList by key
 */
func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
