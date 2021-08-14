// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Gets a shopping list by ID.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ShoppingList by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ShoppingList by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
