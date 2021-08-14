// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Order by ID
 */
func (rb *ByProjectKeyOrdersByIdRequestBuilder) Get() *ByProjectKeyOrdersByIdRequestMethodGet {
	return &ByProjectKeyOrdersByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Order by ID
 */
func (rb *ByProjectKeyOrdersByIdRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersByIdRequestMethodPost {
	return &ByProjectKeyOrdersByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Order by ID
 */
func (rb *ByProjectKeyOrdersByIdRequestBuilder) Delete() *ByProjectKeyOrdersByIdRequestMethodDelete {
	return &ByProjectKeyOrdersByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
