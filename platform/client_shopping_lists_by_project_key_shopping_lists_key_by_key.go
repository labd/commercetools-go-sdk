package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Gets a shopping list by Key.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update a shopping list found by its Key.
 */
func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
