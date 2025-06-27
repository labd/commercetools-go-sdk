package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductSelectionsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
