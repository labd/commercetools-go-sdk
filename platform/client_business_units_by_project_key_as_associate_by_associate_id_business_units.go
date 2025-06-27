package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder struct {
	projectKey  string
	associateId string
	client      *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder{
		key:         key,
		projectKey:  rb.projectKey,
		associateId: rb.associateId,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestBuilder{
		id:          id,
		projectKey:  rb.projectKey,
		associateId: rb.associateId,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units", rb.projectKey, rb.associateId),
		client: rb.client,
	}
}

/**
*	Checks if one or more BusinessUnits exist for the provided query predicate. Returns a `200 OK` status if any BusinessUnits match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units", rb.projectKey, rb.associateId),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder) Post(body BusinessUnitDraft) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units", rb.projectKey, rb.associateId),
		client: rb.client,
	}
}
