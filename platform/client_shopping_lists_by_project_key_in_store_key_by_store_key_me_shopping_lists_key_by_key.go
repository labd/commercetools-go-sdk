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
*	Gets a shopping list by Key.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update a shopping list found by its Key.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
