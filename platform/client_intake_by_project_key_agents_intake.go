package platform

// Generated file, please do not change!!!

type ByProjectKeyAgentsIntakeRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyAgentsIntakeRequestBuilder) V1() *ByProjectKeyAgentsIntakeV1RequestBuilder {
	return &ByProjectKeyAgentsIntakeV1RequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
