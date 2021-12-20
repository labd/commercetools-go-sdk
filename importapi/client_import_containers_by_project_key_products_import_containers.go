// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyProductsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
