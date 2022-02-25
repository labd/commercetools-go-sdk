package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShoppingListsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Gets a shopping list by ID.
 */
func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Get() *ByProjectKeyShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Post(body ShoppingListUpdate) *ByProjectKeyShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
