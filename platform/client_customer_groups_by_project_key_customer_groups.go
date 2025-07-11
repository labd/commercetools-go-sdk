package platform

// Generated file, please do not change!!!

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
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) WithId(id string) *ByProjectKeyCustomerGroupsByIDRequestBuilder {
	return &ByProjectKeyCustomerGroupsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) Get() *ByProjectKeyCustomerGroupsRequestMethodGet {
	return &ByProjectKeyCustomerGroupsRequestMethodGet{
		url:    fmt.Sprintf("/%s/customer-groups", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more CustomerGroups exist for the provided query predicate. Returns a `200 OK` status if any CustomerGroup match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomerGroupsRequestBuilder) Head() *ByProjectKeyCustomerGroupsRequestMethodHead {
	return &ByProjectKeyCustomerGroupsRequestMethodHead{
		url:    fmt.Sprintf("/%s/customer-groups", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomerGroupsRequestBuilder) Post(body CustomerGroupDraft) *ByProjectKeyCustomerGroupsRequestMethodPost {
	return &ByProjectKeyCustomerGroupsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customer-groups", rb.projectKey),
		client: rb.client,
	}
}
