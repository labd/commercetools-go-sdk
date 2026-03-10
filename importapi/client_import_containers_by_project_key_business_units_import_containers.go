package importapi

// Generated file, please do not change!!!

type ByProjectKeyBusinessUnitsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
