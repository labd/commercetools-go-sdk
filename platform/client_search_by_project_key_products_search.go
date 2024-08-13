package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductsSearchRequestBuilder) Post(body ProductSearchRequest) *ByProjectKeyProductsSearchRequestMethodPost {
	return &ByProjectKeyProductsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/search", rb.projectKey),
		client: rb.client,
	}
}
