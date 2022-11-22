package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If a ShoppingList exists in a Project but does _not_ have the `store` field, or the `store` field references a different Store,
*	the [ResourceNotFound](/errors#404-not-found-1) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/shopping-lists/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
