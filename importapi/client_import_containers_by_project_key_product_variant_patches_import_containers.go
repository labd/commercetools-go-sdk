package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyProductVariantPatchesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductVariantPatchesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}