package importapi

// Generated file, please do not change!!!

type ByProjectKeyStandalonePricesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStandalonePricesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
