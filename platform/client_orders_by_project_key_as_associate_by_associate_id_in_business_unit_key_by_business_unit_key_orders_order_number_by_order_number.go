package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	orderNumber     string
	client          *Client
}

/**
*	Retrieves an Order with the provided `orderNumber` in a [BusinessUnit](ctp:api:type:BusinessUnit).
*	If the Order exists in the [Project](ctp:api:type:Project) but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/order-number=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `orderNumber` in a [BusinessUnit](ctp:api:type:BusinessUnit). Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/order-number=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Updates an Order in a [BusinessUnit](ctp:api:type:BusinessUnit) using one or more [update actions](/../api/projects/orders#update-actions).
*	If the Order exists in the [Project](ctp:api:type:Project) but does not reference the requested Business Unit, this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/order-number=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.orderNumber),
		client: rb.client,
	}
}
