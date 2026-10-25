package platform

// Generated file, please do not change!!!

type ByProjectKeyAgentsIntakeV1RequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyAgentsIntakeV1RequestBuilder) Responses() *ByProjectKeyAgentsIntakeV1ResponsesRequestBuilder {
	return &ByProjectKeyAgentsIntakeV1ResponsesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
