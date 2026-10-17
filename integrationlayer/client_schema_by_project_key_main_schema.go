package integrationlayer

// Generated file, please do not change!!!

type ByProjectKeyMainSchemaRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainSchemaRequestBuilder) Visibility() *ByProjectKeyMainSchemaVisibilityRequestBuilder {
	return &ByProjectKeyMainSchemaVisibilityRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
