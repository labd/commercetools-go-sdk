// Generated file, please do not change!!!
package platform

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
func (rb *ByProjectKeyShoppingListsRequestBuilder) Get() *ByProjectKeyShoppingListsRequestMethodGet {
	return &ByProjectKeyShoppingListsRequestMethodGet{
		url:    fmt.Sprintf("/%s/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyShoppingListsRequestBuilder) Post(body ShoppingListDraft) *ByProjectKeyShoppingListsRequestMethodPost {
	return &ByProjectKeyShoppingListsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/shopping-lists", rb.projectKey),
		client: rb.client,
	}
}
