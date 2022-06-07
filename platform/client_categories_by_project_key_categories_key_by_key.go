package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Either the [scope](/../api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Get() *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesKeyByKeyRequestMethodPost {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Delete() *ByProjectKeyCategoriesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
