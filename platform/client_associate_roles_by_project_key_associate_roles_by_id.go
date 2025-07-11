package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAssociateRolesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyAssociateRolesByIDRequestBuilder) Get() *ByProjectKeyAssociateRolesByIDRequestMethodGet {
	return &ByProjectKeyAssociateRolesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/associate-roles/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an AssociateRole exists with the provided `id`. Returns a `200 OK` status if the AssociateRole exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAssociateRolesByIDRequestBuilder) Head() *ByProjectKeyAssociateRolesByIDRequestMethodHead {
	return &ByProjectKeyAssociateRolesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/associate-roles/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAssociateRolesByIDRequestBuilder) Post(body AssociateRoleUpdate) *ByProjectKeyAssociateRolesByIDRequestMethodPost {
	return &ByProjectKeyAssociateRolesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/associate-roles/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAssociateRolesByIDRequestBuilder) Delete() *ByProjectKeyAssociateRolesByIDRequestMethodDelete {
	return &ByProjectKeyAssociateRolesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/associate-roles/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
