package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductTailoringsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTailoringsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
