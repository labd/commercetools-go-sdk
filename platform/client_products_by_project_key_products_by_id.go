// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductsByIdRequestBuilder) Images() *ByProjectKeyProductsByIdImagesRequestBuilder {
	return &ByProjectKeyProductsByIdImagesRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	Gets the full representation of a product by ID.
 */
func (rb *ByProjectKeyProductsByIdRequestBuilder) Get() *ByProjectKeyProductsByIdRequestMethodGet {
	return &ByProjectKeyProductsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Product by ID
 */
func (rb *ByProjectKeyProductsByIdRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsByIdRequestMethodPost {
	return &ByProjectKeyProductsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Product by ID
 */
func (rb *ByProjectKeyProductsByIdRequestBuilder) Delete() *ByProjectKeyProductsByIdRequestMethodDelete {
	return &ByProjectKeyProductsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
