package importapi

// Generated file, please do not change!!!

type ByProjectKeyTypesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTypesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyTypesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyTypesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
