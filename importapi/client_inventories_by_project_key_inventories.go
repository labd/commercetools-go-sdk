// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyInventoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoriesRequestBuilder) ImportContainers() *ByProjectKeyInventoriesImportContainersRequestBuilder {
	return &ByProjectKeyInventoriesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInventoriesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
