// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyProductTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTypesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
