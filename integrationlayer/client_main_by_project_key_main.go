package integrationlayer

// Generated file, please do not change!!!

type ByProjectKeyMainRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainRequestBuilder) Settings() *ByProjectKeyMainSettingsRequestBuilder {
	return &ByProjectKeyMainSettingsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMainRequestBuilder) Checkout() *ByProjectKeyMainCheckoutRequestBuilder {
	return &ByProjectKeyMainCheckoutRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMainRequestBuilder) Schema() *ByProjectKeyMainSchemaRequestBuilder {
	return &ByProjectKeyMainSchemaRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMainRequestBuilder) Idp() *ByProjectKeyMainIdpRequestBuilder {
	return &ByProjectKeyMainIdpRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMainRequestBuilder) Extension() *ByProjectKeyMainExtensionRequestBuilder {
	return &ByProjectKeyMainExtensionRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
