package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsByIDProductSelectionsRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductsByIDProductSelectionsRequestBuilder) Get() *ByProjectKeyProductsByIDProductSelectionsRequestMethodGet {
	return &ByProjectKeyProductsByIDProductSelectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/%s/product-selections", rb.projectKey, rb.id),
		client: rb.client,
	}
}
