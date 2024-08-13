package importapi

// Generated file, please do not change!!!

type ByProjectKeyDiscountCodesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDiscountCodesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyDiscountCodesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
