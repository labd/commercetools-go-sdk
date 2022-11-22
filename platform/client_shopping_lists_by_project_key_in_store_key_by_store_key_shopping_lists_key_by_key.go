package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
