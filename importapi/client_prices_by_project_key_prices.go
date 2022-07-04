package importapi

// Generated file, please do not change!!!

type ByProjectKeyPricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPricesRequestBuilder) ImportContainers() *ByProjectKeyPricesImportContainersRequestBuilder {
	return &ByProjectKeyPricesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
