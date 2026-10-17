package integrationlayer

// Generated file, please do not change!!!

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) Main() *ByProjectKeyMainRequestBuilder {
	return &ByProjectKeyMainRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
