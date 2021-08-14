// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCategoriesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Category by key
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Get() *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Category by key
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesKeyByKeyRequestMethodPost {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Category by key
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Delete() *ByProjectKeyCategoriesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
