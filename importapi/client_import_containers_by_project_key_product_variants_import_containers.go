// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyProductVariantsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductVariantsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
