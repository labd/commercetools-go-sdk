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
*	Gets the current or staged representation of a [Product](ctp:api:type:Product) found by Key.
*	When used with an API Client that has the `view_published_products:{projectKey}` scope,
*	this endpoint only returns published (current) Product Projections.
*
 */
func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductProjectionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if the current or staged representation of a Product exists for a given `key`. Returns a `200 OK` status if the ProductProjection exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductProjectionsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductProjectionsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductProjectionsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-projections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
