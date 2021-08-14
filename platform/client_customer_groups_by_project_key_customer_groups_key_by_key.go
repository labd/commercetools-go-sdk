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

/**
*	Gets a customer group by Key.
 */
func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Get() *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a customer group by Key.
 */
func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Post(body CustomerGroupUpdate) *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete CustomerGroup by key
 */
func (rb *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCustomerGroupsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customer-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
