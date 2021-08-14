// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCategoriesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Category by ID
 */
func (rb *ByProjectKeyCategoriesByIdRequestBuilder) Get() *ByProjectKeyCategoriesByIdRequestMethodGet {
	return &ByProjectKeyCategoriesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Category by ID
 */
func (rb *ByProjectKeyCategoriesByIdRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesByIdRequestMethodPost {
	return &ByProjectKeyCategoriesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Category by ID
 */
func (rb *ByProjectKeyCategoriesByIdRequestBuilder) Delete() *ByProjectKeyCategoriesByIdRequestMethodDelete {
	return &ByProjectKeyCategoriesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
