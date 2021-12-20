// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyCategoriesImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCategoriesImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
