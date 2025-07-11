package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder struct {
	projectKey  string
	storeKey    string
	key         string
	associateId string
	client      *Client
}

/**
*	Retrieves roles and permissions of an Associate in a Business Unit in a Store.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyAssociatesByAssociateIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/key=%s/associates/%s", rb.projectKey, rb.storeKey, rb.key, rb.associateId),
		client: rb.client,
	}
}
