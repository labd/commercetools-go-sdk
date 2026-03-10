package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder) WithKey(key string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsKeyByKeyRequestBuilder{
		key:             key,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder) WithId(id string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsByIDRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsByIDRequestBuilder{
		id:              id,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}

/**
*	Retrieves all QuoteRequests in a BusinessUnit.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more QuoteRequests exist for the provided query predicate in a BusinessUnit. Returns a `200 OK` status if any QuoteRequests match the query predicate, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Creates a QuoteRequest in a BusinessUnit.
*	Creating QuoteRequest fails with an [InvalidOperation](ctp:api:type:InvalidOperationError) if the Cart does not reference the same BusinessUnit as the `businessUnitKey` path parameter.
*
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestBuilder) Post(body QuoteRequestDraft) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuoteRequestsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quote-requests", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
