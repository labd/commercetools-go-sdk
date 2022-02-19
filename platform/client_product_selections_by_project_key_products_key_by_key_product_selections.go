// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder) Get() *ByProjectKeyProductsKeyByKeyProductSelectionsRequestMethodGet {
	return &ByProjectKeyProductsKeyByKeyProductSelectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/key=%s/product-selections", rb.projectKey, rb.key),
		client: rb.client,
	}
}
