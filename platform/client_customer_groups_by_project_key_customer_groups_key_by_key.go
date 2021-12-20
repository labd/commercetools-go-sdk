// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Get() *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Post(body CustomerGroupUpdate) *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
