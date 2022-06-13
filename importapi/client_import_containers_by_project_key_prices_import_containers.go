package importapi

// Generated file, please do not change!!!

type ByProjectKeyPricesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPricesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
