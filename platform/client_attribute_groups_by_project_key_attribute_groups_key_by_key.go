package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder) Get() *ByProjectKeyAttributeGroupsKeyByKeyRequestMethodGet {
	return &ByProjectKeyAttributeGroupsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/attribute-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if an AttributeGroup exists with the provided `key`. Returns `200 OK` status if the AttributeGroup exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder) Head() *ByProjectKeyAttributeGroupsKeyByKeyRequestMethodHead {
	return &ByProjectKeyAttributeGroupsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/attribute-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder) Post(body AttributeGroupUpdate) *ByProjectKeyAttributeGroupsKeyByKeyRequestMethodPost {
	return &ByProjectKeyAttributeGroupsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/attribute-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAttributeGroupsKeyByKeyRequestBuilder) Delete() *ByProjectKeyAttributeGroupsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyAttributeGroupsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/attribute-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
