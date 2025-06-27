package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder struct {
	projectKey     string
	businessUnitId string
	associateId    string
	client         *Client
}

/**
*	Retrieves roles and permissions of an Associate in a Business Unit.
*
 */
func (rb *ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder) Get() *ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestMethodGet {
	return &ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units/%s/associates/%s", rb.projectKey, rb.businessUnitId, rb.associateId),
		client: rb.client,
	}
}
