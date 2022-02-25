package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsByIDProductsRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsByIDProductsRequestBuilder) Get() *ByProjectKeyProductSelectionsByIDProductsRequestMethodGet {
	return &ByProjectKeyProductSelectionsByIDProductsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-selections/%s/products", rb.projectKey, rb.id),
		client: rb.client,
	}
}
