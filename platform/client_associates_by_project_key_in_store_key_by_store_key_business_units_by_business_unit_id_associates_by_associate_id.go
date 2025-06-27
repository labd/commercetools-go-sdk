package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder struct {
	projectKey     string
	storeKey       string
	businessUnitId string
	associateId    string
	client         *Client
}

/**
*	Retrieves roles and permissions of an Associate in a Business Unit in a Store.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByBusinessUnitIdAssociatesByAssociateIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/%s/associates/%s", rb.projectKey, rb.storeKey, rb.businessUnitId, rb.associateId),
		client: rb.client,
	}
}
