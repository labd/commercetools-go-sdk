// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyOrderPatchesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrderPatchesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
