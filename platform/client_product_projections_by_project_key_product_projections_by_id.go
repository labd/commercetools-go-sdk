package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) by its ID. When used with an API Client that has the `view_published_products:{projectKey}` scope, this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyProductProjectionsByIDRequestBuilder) Get() *ByProjectKeyProductProjectionsByIDRequestMethodGet {
	return &ByProjectKeyProductProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
