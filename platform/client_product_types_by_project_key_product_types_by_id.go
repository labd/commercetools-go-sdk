// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductTypesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get ProductType by ID
 */
func (rb *ByProjectKeyProductTypesByIdRequestBuilder) Get() *ByProjectKeyProductTypesByIdRequestMethodGet {
	return &ByProjectKeyProductTypesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ProductType by ID
 */
func (rb *ByProjectKeyProductTypesByIdRequestBuilder) Post(body ProductTypeUpdate) *ByProjectKeyProductTypesByIdRequestMethodPost {
	return &ByProjectKeyProductTypesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ProductType by ID
 */
func (rb *ByProjectKeyProductTypesByIdRequestBuilder) Delete() *ByProjectKeyProductTypesByIdRequestMethodDelete {
	return &ByProjectKeyProductTypesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
