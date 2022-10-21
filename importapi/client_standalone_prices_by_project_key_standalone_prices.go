package importapi

// Generated file, please do not change!!!

type ByProjectKeyStandalonePricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesRequestBuilder) ImportContainers() *ByProjectKeyStandalonePricesImportContainersRequestBuilder {
	return &ByProjectKeyStandalonePricesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
