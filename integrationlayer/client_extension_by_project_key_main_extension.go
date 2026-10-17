package integrationlayer

// Generated file, please do not change!!!

type ByProjectKeyMainExtensionRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainExtensionRequestBuilder) Config() *ByProjectKeyMainExtensionConfigRequestBuilder {
	return &ByProjectKeyMainExtensionConfigRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMainExtensionRequestBuilder) Allowlist() *ByProjectKeyMainExtensionAllowlistRequestBuilder {
	return &ByProjectKeyMainExtensionAllowlistRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
