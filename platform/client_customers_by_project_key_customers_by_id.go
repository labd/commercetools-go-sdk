// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCustomersByIDRequestBuilder) Get() *ByProjectKeyCustomersByIDRequestMethodGet {
	return &ByProjectKeyCustomersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomersByIDRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyCustomersByIDRequestMethodPost {
	return &ByProjectKeyCustomersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomersByIDRequestBuilder) Delete() *ByProjectKeyCustomersByIDRequestMethodDelete {
	return &ByProjectKeyCustomersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
