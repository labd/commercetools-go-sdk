package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductProjectionsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Product Projection Search
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Post(body string) *ByProjectKeyProductProjectionsSearchRequestMethodPost {
	return &ByProjectKeyProductProjectionsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Product Projection Search
 */
func (rb *ByProjectKeyProductProjectionsSearchRequestBuilder) Get() *ByProjectKeyProductProjectionsSearchRequestMethodGet {
	return &ByProjectKeyProductProjectionsSearchRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/search", rb.projectKey),
		client: rb.client,
	}
}
