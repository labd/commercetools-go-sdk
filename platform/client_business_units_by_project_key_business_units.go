package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyBusinessUnitsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) KeyWithKeyValueAssociatesWithAssociateIdValue(key string, associateId string) *ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder {
	return &ByProjectKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder{
		key:         key,
		associateId: associateId,
		projectKey:  rb.projectKey,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) WithBusinessUnitIdValueAssociatesWithAssociateIdValue(businessUnitId string, associateId string) *ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder {
	return &ByProjectKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder{
		businessUnitId: businessUnitId,
		associateId:    associateId,
		projectKey:     rb.projectKey,
		client:         rb.client,
	}
}

/**
*	This endpoint provides high-performance search queries over Business Units.
*
 */
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Search() *ByProjectKeyBusinessUnitsSearchRequestBuilder {
	return &ByProjectKeyBusinessUnitsSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) SearchIndexingStatus() *ByProjectKeyBusinessUnitsSearchIndexingStatusRequestBuilder {
	return &ByProjectKeyBusinessUnitsSearchIndexingStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Get() *ByProjectKeyBusinessUnitsRequestMethodGet {
	return &ByProjectKeyBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more BusinessUnits exist for the provided query predicate. Returns a `200 OK` status if any BusinessUnits match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Head() *ByProjectKeyBusinessUnitsRequestMethodHead {
	return &ByProjectKeyBusinessUnitsRequestMethodHead{
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Post(body BusinessUnitDraft) *ByProjectKeyBusinessUnitsRequestMethodPost {
	return &ByProjectKeyBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}
