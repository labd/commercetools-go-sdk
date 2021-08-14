// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductProjectionsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Gets the current or staged representation of a product in a catalog by ID.
*	When used with an API client that has the view_published_products:{projectKey} scope,
*	this endpoint only returns published (current) product projections.
*
 */
func (rb *ByProjectKeyProductProjectionsByIdRequestBuilder) Get() *ByProjectKeyProductProjectionsByIdRequestMethodGet {
	return &ByProjectKeyProductProjectionsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
