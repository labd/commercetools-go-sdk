// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomerGroupsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomerGroupsRequestBuilder) WithKey(key string) *ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder {
	return &ByProjectKeyCustomerGroupsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) WithId(id string) *ByProjectKeyCustomerGroupsByIdRequestBuilder {
	return &ByProjectKeyCustomerGroupsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query customer-groups
 */
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) Get() *ByProjectKeyCustomerGroupsRequestMethodGet {
	return &ByProjectKeyCustomerGroupsRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create CustomerGroup
 */
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) Post(body CustomerGroupDraft) *ByProjectKeyCustomerGroupsRequestMethodPost {
	return &ByProjectKeyCustomerGroupsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups", rb.projectKey),
		client: rb.client,
	}
}
