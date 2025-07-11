package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	id              string
	client          *Client
}

/**
*	Retrieves an Order with the provided `id` in a [BusinessUnit](ctp:api:type:BusinessUnit).
*	If the Order exists in the [Project](ctp:api:type:Project) but does not reference the requested Business Unit, this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `id` in a [BusinessUnit](ctp:api:type:BusinessUnit). Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates an Order in a [BusinessUnit](ctp:api:type:BusinessUnit) using one or more [update actions](/../api/projects/orders#update-actions).
*	If the Order exists in the [Project](ctp:api:type:Project) but does not reference the requested Business Unit, this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestBuilder) Post(body OrderUpdate) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.id),
		client: rb.client,
	}
}
