package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductsByIDRequestBuilder) Images() *ByProjectKeyProductsByIDImagesRequestBuilder {
	return &ByProjectKeyProductsByIDImagesRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductsByIDRequestBuilder) ProductSelections() *ByProjectKeyProductsByIDProductSelectionsRequestBuilder {
	return &ByProjectKeyProductsByIDProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	Gets the full representation of a product by ID.
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Get() *ByProjectKeyProductsByIDRequestMethodGet {
	return &ByProjectKeyProductsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if product with given ID exists.
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Head() *ByProjectKeyProductsByIDRequestMethodHead {
	return &ByProjectKeyProductsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductsByIDRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsByIDRequestMethodPost {
	return &ByProjectKeyProductsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductsByIDRequestBuilder) Delete() *ByProjectKeyProductsByIDRequestMethodDelete {
	return &ByProjectKeyProductsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
