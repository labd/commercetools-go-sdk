// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductProjectionsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Search Product Projection
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Post(body string) *ByProjectKeyProductProjectionsSearchRequestMethodPost {
	return &ByProjectKeyProductProjectionsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Search Product Projection
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Get() *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	return &ByProjectKeyProductProjectionsSearchRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}
