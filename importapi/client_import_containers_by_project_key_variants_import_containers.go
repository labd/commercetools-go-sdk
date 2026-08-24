package importapi

// Generated file, please do not change!!!

type ByProjectKeyVariantsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyVariantsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyVariantsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyVariantsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
