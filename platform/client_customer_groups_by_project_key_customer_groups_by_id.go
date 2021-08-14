// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomerGroupsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get CustomerGroup by ID
 */
func (rb *ByProjectKeyCustomerGroupsByIdRequestBuilder) Get() *ByProjectKeyCustomerGroupsByIdRequestMethodGet {
	return &ByProjectKeyCustomerGroupsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update CustomerGroup by ID
 */
func (rb *ByProjectKeyCustomerGroupsByIdRequestBuilder) Post(body CustomerGroupUpdate) *ByProjectKeyCustomerGroupsByIdRequestMethodPost {
	return &ByProjectKeyCustomerGroupsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete CustomerGroup by ID
 */
func (rb *ByProjectKeyCustomerGroupsByIdRequestBuilder) Delete() *ByProjectKeyCustomerGroupsByIdRequestMethodDelete {
	return &ByProjectKeyCustomerGroupsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
