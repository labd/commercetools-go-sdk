package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) KeyWithKeyValueAssociatesWithAssociateIdValue(key string, associateId string) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder{
		key:         key,
		associateId: associateId,
		projectKey:  rb.projectKey,
		storeKey:    rb.storeKey,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) WithBusinessUnitIdValueAssociatesWithAssociateIdValue(businessUnitId string, associateId string) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder{
		businessUnitId: businessUnitId,
		associateId:    associateId,
		projectKey:     rb.projectKey,
		storeKey:       rb.storeKey,
		client:         rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more BusinessUnits exist for the provided query predicate. Returns a `200 OK` status if any BusinessUnits match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestBuilder) Post(body BusinessUnitDraft) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
