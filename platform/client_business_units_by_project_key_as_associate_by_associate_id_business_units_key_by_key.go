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

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/key=%s", rb.projectKey, rb.associateId, rb.key),
		client: rb.client,
	}
}
