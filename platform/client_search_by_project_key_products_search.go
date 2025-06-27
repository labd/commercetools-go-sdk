package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	If indexing is in progress or if Product Search is inactive, an [ObjectNotFound](ctp:api:type:ObjectNotFoundError) error is returned.
*	If inactive, you can [reactivate](/../api/projects/product-search#activate-the-product-search-api) it.
*
 */
func (rb *ByProjectKeyProductsSearchRequestBuilder) Post(body ProductSearchRequest) *ByProjectKeyProductsSearchRequestMethodPost {
	return &ByProjectKeyProductsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/search", rb.projectKey),
		client: rb.client,
	}
}
