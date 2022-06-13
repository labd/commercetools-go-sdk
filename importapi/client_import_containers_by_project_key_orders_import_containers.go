package importapi

// Generated file, please do not change!!!

type ByProjectKeyOrdersImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyOrdersImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyOrdersImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
