package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Gets the current or staged representation of a product found by Key.
*	When used with an API client that has the view_published_products:{projectKey} scope,
*	this endpoint only returns published (current) product projections.
*
 */
func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
