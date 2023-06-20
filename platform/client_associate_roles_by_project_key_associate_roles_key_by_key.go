package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAssociateRolesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyAssociateRolesKeyByKeyRequestBuilder) Get() *ByProjectKeyAssociateRolesKeyByKeyRequestMethodGet {
	return &ByProjectKeyAssociateRolesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/associate-roles/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAssociateRolesKeyByKeyRequestBuilder) Post(body AssociateRoleUpdate) *ByProjectKeyAssociateRolesKeyByKeyRequestMethodPost {
	return &ByProjectKeyAssociateRolesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/associate-roles/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deleting an AssociateRole generates the [AssociateRoleDeleted](ctp:api:type:AssociateRoleDeletedMessage) Message. An AssociateRole can only be deleted if it is not assigned to any [Associates](ctp:api:type:Associate).
*
 */
func (rb *ByProjectKeyAssociateRolesKeyByKeyRequestBuilder) Delete() *ByProjectKeyAssociateRolesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyAssociateRolesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/associate-roles/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
