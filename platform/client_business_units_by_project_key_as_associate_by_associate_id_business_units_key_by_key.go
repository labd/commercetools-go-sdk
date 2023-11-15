package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder struct {
	projectKey  string
	associateId string
	key         string
	client      *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/key=%s", rb.projectKey, rb.associateId, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists for a given `key`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder) Head() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodHead {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/key=%s", rb.projectKey, rb.associateId, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/key=%s", rb.projectKey, rb.associateId, rb.key),
		client: rb.client,
	}
}
