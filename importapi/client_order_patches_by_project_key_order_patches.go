// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyOrderPatchesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrderPatchesRequestBuilder) ImportContainers() *ByProjectKeyOrderPatchesImportContainersRequestBuilder {
	return &ByProjectKeyOrderPatchesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrderPatchesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
