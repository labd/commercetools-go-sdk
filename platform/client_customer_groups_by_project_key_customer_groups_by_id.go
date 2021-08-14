// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomerGroupsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestBuilder) Get() *ByProjectKeyCustomerGroupsByIDRequestMethodGet {
	return &ByProjectKeyCustomerGroupsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestBuilder) Post(body CustomerGroupUpdate) *ByProjectKeyCustomerGroupsByIDRequestMethodPost {
	return &ByProjectKeyCustomerGroupsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestBuilder) Delete() *ByProjectKeyCustomerGroupsByIDRequestMethodDelete {
	return &ByProjectKeyCustomerGroupsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customer-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
