package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAttributeGroupsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyAttributeGroupsByIDRequestBuilder) Get() *ByProjectKeyAttributeGroupsByIDRequestMethodGet {
	return &ByProjectKeyAttributeGroupsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/attribute-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAttributeGroupsByIDRequestBuilder) Post(body AttributeGroupUpdate) *ByProjectKeyAttributeGroupsByIDRequestMethodPost {
	return &ByProjectKeyAttributeGroupsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/attribute-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAttributeGroupsByIDRequestBuilder) Delete() *ByProjectKeyAttributeGroupsByIDRequestMethodDelete {
	return &ByProjectKeyAttributeGroupsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/attribute-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
