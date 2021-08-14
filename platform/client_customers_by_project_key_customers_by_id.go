// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Customer by ID
 */
func (rb *ByProjectKeyCustomersByIdRequestBuilder) Get() *ByProjectKeyCustomersByIdRequestMethodGet {
	return &ByProjectKeyCustomersByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Customer by ID
 */
func (rb *ByProjectKeyCustomersByIdRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyCustomersByIdRequestMethodPost {
	return &ByProjectKeyCustomersByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Customer by ID
 */
func (rb *ByProjectKeyCustomersByIdRequestBuilder) Delete() *ByProjectKeyCustomersByIdRequestMethodDelete {
	return &ByProjectKeyCustomersByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
