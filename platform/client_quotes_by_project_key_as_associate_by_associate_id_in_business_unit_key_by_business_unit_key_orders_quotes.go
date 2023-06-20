package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

/**
*	Creates an Order from a [Quote](ctp:api:type:Cart) in a [BusinessUnit](ctp:api:type:BusinessUnit).
*	Creating an Order fails with an [InvalidOperation](ctp:api:type:InvalidOperationError) if the Quote does not reference the same BusinessUnit as the `businessUnitKey` path parameter.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/quotes", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
