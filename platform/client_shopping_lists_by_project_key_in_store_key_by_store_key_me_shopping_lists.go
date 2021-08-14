// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Post(body MyShoppingListDraft) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
