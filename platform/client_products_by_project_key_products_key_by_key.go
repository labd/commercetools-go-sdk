package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) ProductSelections() *ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder {
	return &ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
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
*	Checks if product with given key exists.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
