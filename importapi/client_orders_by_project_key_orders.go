package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersRequestBuilder) ImportContainers() *ByProjectKeyOrdersImportContainersRequestBuilder {
	return &ByProjectKeyOrdersImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
