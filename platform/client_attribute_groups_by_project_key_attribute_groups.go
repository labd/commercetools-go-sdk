package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAttributeGroupsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyAttributeGroupsRequestBuilder) WithKey(key string) *ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder {
	return &ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyAttributeGroupsRequestBuilder) WithId(id string) *ByProjectKeyAttributeGroupsByIDRequestBuilder {
	return &ByProjectKeyAttributeGroupsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyAttributeGroupsRequestBuilder) Get() *ByProjectKeyAttributeGroupsRequestMethodGet {
	return &ByProjectKeyAttributeGroupsRequestMethodGet{
		url:    fmt.Sprintf("/%s/attribute-groups", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more AttributeGroups exist for the provided query predicate. Returns `200 OK` status if any AttributeGroups match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAttributeGroupsRequestBuilder) Head() *ByProjectKeyAttributeGroupsRequestMethodHead {
	return &ByProjectKeyAttributeGroupsRequestMethodHead{
		url:    fmt.Sprintf("/%s/attribute-groups", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAttributeGroupsRequestBuilder) Post(body AttributeGroupDraft) *ByProjectKeyAttributeGroupsRequestMethodPost {
	return &ByProjectKeyAttributeGroupsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/attribute-groups", rb.projectKey),
		client: rb.client,
	}
}
