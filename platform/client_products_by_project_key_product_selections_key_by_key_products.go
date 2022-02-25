package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsKeyByKeyProductsRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyProductsRequestBuilder) Get() *ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet {
	return &ByProjectKeyProductSelectionsKeyByKeyProductsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-selections/key=%s/products", rb.projectKey, rb.key),
		client: rb.client,
	}
}
