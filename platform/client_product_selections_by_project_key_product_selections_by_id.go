package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsByIDRequestBuilder) Products() *ByProjectKeyProductSelectionsByIDProductsRequestBuilder {
	return &ByProjectKeyProductSelectionsByIDProductsRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductSelectionsByIDRequestBuilder) Get() *ByProjectKeyProductSelectionsByIDRequestMethodGet {
	return &ByProjectKeyProductSelectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-selections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a ProductSelection exists for a given `id`. Returns a `200 OK` status if the ProductSelection exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductSelectionsByIDRequestBuilder) Head() *ByProjectKeyProductSelectionsByIDRequestMethodHead {
	return &ByProjectKeyProductSelectionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-selections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductSelectionsByIDRequestBuilder) Post(body ProductSelectionUpdate) *ByProjectKeyProductSelectionsByIDRequestMethodPost {
	return &ByProjectKeyProductSelectionsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-selections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletion will only succeed if the Product Selection is not assigned to any [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyProductSelectionsByIDRequestBuilder) Delete() *ByProjectKeyProductSelectionsByIDRequestMethodDelete {
	return &ByProjectKeyProductSelectionsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-selections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
