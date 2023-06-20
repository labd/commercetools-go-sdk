package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAssociateRolesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyAssociateRolesRequestBuilder) WithKey(key string) *ByProjectKeyAssociateRolesKeyByKeyRequestBuilder {
	return &ByProjectKeyAssociateRolesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyAssociateRolesRequestBuilder) WithId(id string) *ByProjectKeyAssociateRolesByIDRequestBuilder {
	return &ByProjectKeyAssociateRolesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyAssociateRolesRequestBuilder) Get() *ByProjectKeyAssociateRolesRequestMethodGet {
	return &ByProjectKeyAssociateRolesRequestMethodGet{
		url:    fmt.Sprintf("/%s/associate-roles", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a Associate Role generates the [AssociateRoleCreated](ctp:api:type:AssociateRoleCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyAssociateRolesRequestBuilder) Post(body AssociateRoleDraft) *ByProjectKeyAssociateRolesRequestMethodPost {
	return &ByProjectKeyAssociateRolesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/associate-roles", rb.projectKey),
		client: rb.client,
	}
}
