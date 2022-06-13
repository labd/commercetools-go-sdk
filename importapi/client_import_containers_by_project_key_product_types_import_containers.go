package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductTypesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTypesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
