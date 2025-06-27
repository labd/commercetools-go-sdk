package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder struct {
	projectKey  string
	key         string
	associateId string
	client      *Client
}

/**
*	Retrieves roles and permissions of an Associate in a Business Unit.
*
 */
func (rb *ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder) Get() *ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestMethodGet {
	return &ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units/key=%s/associates/%s", rb.projectKey, rb.key, rb.associateId),
		client: rb.client,
	}
}
