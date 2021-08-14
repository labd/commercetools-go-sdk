// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyOrderPatchesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrderPatchesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
