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
*
*	Creates an Order from a [Quote](ctp:api:type:Quote) in a [BusinessUnit](ctp:api:type:BusinessUnit).
*
*	The Quote must reference the same Business Unit as the `businessUnitKey` path parameter, must have the `Pending` [state](ctp:api:type:QuoteState), and must be valid (not past the `validTo` date). If these criteria are not met, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
*
*	Specific Error Codes:
*
*	- [InvalidItemShippingDetails](ctp:api:type:InvalidItemShippingDetailsError)
*	- [OutOfStock](ctp:api:type:OutOfStockError)
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestBuilder) Post(body OrderFromQuoteDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyOrdersQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/orders/quotes", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
