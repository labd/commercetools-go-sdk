// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Gets a shopping list by ID.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ShoppingList by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ShoppingList by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
