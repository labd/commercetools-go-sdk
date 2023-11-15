package platform

// Generated file, please do not change!!!

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

/**
*	Checks if a ShoppingList exists for a given Query Predicate. Returns a `200 OK` status if any ShoppingLists match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	When using this endpoint, the `store` field of a ShoppingList is always set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestBuilder) Post(body MyShoppingListDraft) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
