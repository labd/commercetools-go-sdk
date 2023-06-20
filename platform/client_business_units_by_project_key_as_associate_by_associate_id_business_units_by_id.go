package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestBuilder struct {
	projectKey  string
	associateId string
	id          string
	client      *Client
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestBuilder) Get() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/%s", rb.projectKey, rb.associateId, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodPost {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/as-associate/%s/business-units/%s", rb.projectKey, rb.associateId, rb.id),
		client: rb.client,
	}
}
