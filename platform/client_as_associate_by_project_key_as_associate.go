package platform

// Generated file, please do not change!!!

type ByProjectKeyAsAssociateRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyAsAssociateRequestBuilder) WithAssociateIdValue(associateId string) *ByProjectKeyAsAssociateByAssociateIdRequestBuilder {
	return &ByProjectKeyAsAssociateByAssociateIdRequestBuilder{
		associateId: associateId,
		projectKey:  rb.projectKey,
		client:      rb.client,
	}
}
