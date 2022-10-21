package platform

// Generated file, please do not change!!!

type ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeRequestBuilder struct {
	projectKey      string
	businessUnitKey string
	client          *Client
}

func (rb *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeRequestBuilder) Customers() *ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestBuilder {
	return &ByProjectKeyInBusinessUnitKeyByBusinessUnitKeyMeCustomersRequestBuilder{
		projectKey:      rb.projectKey,
		businessUnitKey: rb.businessUnitKey,
		client:          rb.client,
	}
}
