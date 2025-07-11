package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTypesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductTypesByIDRequestBuilder) Get() *ByProjectKeyProductTypesByIDRequestMethodGet {
	return &ByProjectKeyProductTypesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ProductType exists with the provided `id`. Returns a `200 OK` status if the ProductType exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductTypesByIDRequestBuilder) Head() *ByProjectKeyProductTypesByIDRequestMethodHead {
	return &ByProjectKeyProductTypesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTypesByIDRequestBuilder) Post(body ProductTypeUpdate) *ByProjectKeyProductTypesByIDRequestMethodPost {
	return &ByProjectKeyProductTypesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTypesByIDRequestBuilder) Delete() *ByProjectKeyProductTypesByIDRequestMethodDelete {
	return &ByProjectKeyProductTypesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
