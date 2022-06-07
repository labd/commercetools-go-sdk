package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Either the [scope](/../api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Get() *ByProjectKeyCategoriesByIDRequestMethodGet {
	return &ByProjectKeyCategoriesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesByIDRequestMethodPost {
	return &ByProjectKeyCategoriesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Delete() *ByProjectKeyCategoriesByIDRequestMethodDelete {
	return &ByProjectKeyCategoriesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
