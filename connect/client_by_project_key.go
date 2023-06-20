package connect

// Generated file, please do not change!!!

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) Deployments() *ByProjectKeyDeploymentsRequestBuilder {
	return &ByProjectKeyDeploymentsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
