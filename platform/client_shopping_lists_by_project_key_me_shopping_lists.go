package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyMeShoppingListsByIDRequestBuilder {
	return &ByProjectKeyMeShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Get() *ByProjectKeyMeShoppingListsRequestMethodGet {
	return &ByProjectKeyMeShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a ShoppingList exists for a given Query Predicate. Returns a `200 OK` status if any ShoppingLists match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Head() *ByProjectKeyMeShoppingListsRequestMethodHead {
	return &ByProjectKeyMeShoppingListsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeShoppingListsRequestBuilder) Post(body MyShoppingListDraft) *ByProjectKeyMeShoppingListsRequestMethodPost {
	return &ByProjectKeyMeShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}
