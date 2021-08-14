// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Gets the full representation of a product by Key.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Product by key
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Product by key
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
