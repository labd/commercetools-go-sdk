// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyCustomersImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyCustomersImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyCustomersImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
