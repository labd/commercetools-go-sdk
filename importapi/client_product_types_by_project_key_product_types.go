package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyProductTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTypesRequestBuilder) ImportContainers() *ByProjectKeyProductTypesImportContainersRequestBuilder {
	return &ByProjectKeyProductTypesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductTypesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
