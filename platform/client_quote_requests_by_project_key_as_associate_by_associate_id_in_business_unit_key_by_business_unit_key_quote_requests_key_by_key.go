package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	key             string
	client          *Client
}

/**
*	Retrieves a QuoteRequest with the provided `key` in a BusinessUnit.
*	If the QuoteRequest exists in the [Project](ctp:api:type:Project) but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists with the provided `key` in a BusinessUnit. Returns a `200 OK` status if the QuoteRequest exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a QuoteRequest in a BusinessUnit using one or more [update actions](/../api/projects/quote-requests#update-actions).
*	If the QuoteRequest exists in the [Project](ctp:api:type:Project) but does not reference the requested [BusinessUnit](ctp:api:type:BusinessUnit), this method returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder) Post(body QuoteRequestUpdate) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests/key=%s", rb.projectKey, rb.associateId, rb.businessUnitKey, rb.key),
		client: rb.client,
	}
}
