package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store, the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
