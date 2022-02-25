package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestBuilder) Post(body ShoppingListDraft) *ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
