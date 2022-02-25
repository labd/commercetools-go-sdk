package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeShoppingListsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Get() *ByProjectKeyMeShoppingListsByIDRequestMethodGet {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Post(body MyShoppingListUpdate) *ByProjectKeyMeShoppingListsByIDRequestMethodPost {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeShoppingListsByIDRequestBuilder) Delete() *ByProjectKeyMeShoppingListsByIDRequestMethodDelete {
	return &ByProjectKeyMeShoppingListsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/shopping-lists/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
