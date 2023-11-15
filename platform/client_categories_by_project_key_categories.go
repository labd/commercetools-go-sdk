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

/**
*	Either the [scope](/../api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesRequestBuilder) Get() *ByProjectKeyCategoriesRequestMethodGet {
	return &ByProjectKeyCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Category exists for on a given Query Predicate. Returns a `200 OK` status if any Categories match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCategoriesRequestBuilder) Head() *ByProjectKeyCategoriesRequestMethodHead {
	return &ByProjectKeyCategoriesRequestMethodHead{
		url:    fmt.Sprintf("/%s/categories", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Either the [scope](/../api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
*	Creating a Category produces the [CategoryCreated](ctp:api:type:CategoryCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyCategoriesRequestBuilder) Post(body CategoryDraft) *ByProjectKeyCategoriesRequestMethodPost {
	return &ByProjectKeyCategoriesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories", rb.projectKey),
		client: rb.client,
	}
}
