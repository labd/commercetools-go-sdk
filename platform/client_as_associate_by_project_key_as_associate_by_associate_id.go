package platform

// Generated file, please do not change!!!

type ByProjectKeyAsAssociateByAssociateIdRequestBuilder struct {
	projectKey  string
	associateId string
	client      *Client
}

/**
*	A Business Unit can represent a Company or a Division.
 */
func (rb *ByProjectKeyAsAssociateByAssociateIdRequestBuilder) BusinessUnits() *ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdBusinessUnitsRequestBuilder{
		projectKey:  rb.projectKey,
		associateId: rb.associateId,
		client:      rb.client,
	}
}
func (rb *ByProjectKeyAsAssociateByAssociateIdRequestBuilder) InBusinessUnitKeyWithBusinessUnitKeyValue(businessUnitKey string) *ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdInBusinessUnitKeyByBusinessUnitKeyRequestBuilder{
		businessUnitKey: businessUnitKey,
		projectKey:      rb.projectKey,
		associateId:     rb.associateId,
		client:          rb.client,
	}
}
