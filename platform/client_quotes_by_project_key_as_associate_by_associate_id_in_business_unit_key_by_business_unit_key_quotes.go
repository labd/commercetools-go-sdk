package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestBuilder struct {
	projectKey      string
	associateId     string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestBuilder) WithKey(key string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesKeyByKeyRequestBuilder{
		key:             key,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestBuilder) WithId(id string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesByIDRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesByIDRequestBuilder{
		id:              id,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quotes", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Quotes exist for the provided query predicate. Returns a `200 OK` status if any Quotes match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyQuotesRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/in-business-unit/key=%s/quotes", rb.projectKey, rb.associateId, rb.businessUnitKey),
		client: rb.client,
	}
}
