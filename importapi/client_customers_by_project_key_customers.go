package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyCustomersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersRequestBuilder) ImportContainers() *ByProjectKeyCustomersImportContainersRequestBuilder {
	return &ByProjectKeyCustomersImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCustomersRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
