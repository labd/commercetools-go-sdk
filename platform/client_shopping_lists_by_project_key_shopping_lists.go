package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShoppingListsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyShoppingListsRequestBuilder) WithKey(key string) *ByProjectKeyShoppingListsKeyByKeyRequestBuilder {
	return &ByProjectKeyShoppingListsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyShoppingListsRequestBuilder) WithId(id string) *ByProjectKeyShoppingListsByIDRequestBuilder {
	return &ByProjectKeyShoppingListsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves ShoppingLists in the Project.
 */
func (rb *ByProjectKeyShoppingListsRequestBuilder) Get() *ByProjectKeyShoppingListsRequestMethodGet {
	return &ByProjectKeyShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more ShoppingLists exist for the provided query predicate. Returns a `200 OK` status if any ShoppingLists match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShoppingListsRequestBuilder) Head() *ByProjectKeyShoppingListsRequestMethodHead {
	return &ByProjectKeyShoppingListsRequestMethodHead{
		url:    fmt.Sprintf("/%s/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create a ShoppingList in the Project.
 */
func (rb *ByProjectKeyShoppingListsRequestBuilder) Post(body ShoppingListDraft) *ByProjectKeyShoppingListsRequestMethodPost {
	return &ByProjectKeyShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}
