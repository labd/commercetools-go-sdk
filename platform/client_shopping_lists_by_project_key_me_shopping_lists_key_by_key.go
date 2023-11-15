package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists for a given `key`. Returns a `200 OK` status if the ShoppingList exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Head() *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeShoppingListsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/shopping-lists/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
