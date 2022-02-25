package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCategoriesRequestBuilder) WithKey(key string) *ByProjectKeyCategoriesKeyByKeyRequestBuilder {
	return &ByProjectKeyCategoriesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCategoriesRequestBuilder) WithId(id string) *ByProjectKeyCategoriesByIDRequestBuilder {
	return &ByProjectKeyCategoriesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCategoriesRequestBuilder) Get() *ByProjectKeyCategoriesRequestMethodGet {
	return &ByProjectKeyCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a category produces the CategoryCreated message.
 */
func (rb *ByProjectKeyCategoriesRequestBuilder) Post(body CategoryDraft) *ByProjectKeyCategoriesRequestMethodPost {
	return &ByProjectKeyCategoriesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories", rb.projectKey),
		client: rb.client,
	}
}
