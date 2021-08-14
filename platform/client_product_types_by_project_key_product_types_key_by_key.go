// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductTypesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get ProductType by key
 */
func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Get() *ByProjectKeyProductTypesKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update ProductType by key
 */
func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Post(body ProductTypeUpdate) *ByProjectKeyProductTypesKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete ProductType by key
 */
func (rb *ByProjectKeyProductTypesKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductTypesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductTypesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-types/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
