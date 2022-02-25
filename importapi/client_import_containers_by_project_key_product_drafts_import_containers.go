package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyProductDraftsImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductDraftsImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}
