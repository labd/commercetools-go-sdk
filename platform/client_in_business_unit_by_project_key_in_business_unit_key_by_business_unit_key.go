package platform

// Generated file, please do not change!!!

type ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyRequestBuilder struct {
	projectKey      string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyRequestBuilder) Me() *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeRequestBuilder {
	return &ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeRequestBuilder{
		projectKey:      rb.projectKey,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
