package importapi

// Generated file, please do not change!!!

type ByProjectKeyInventoriesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoriesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
