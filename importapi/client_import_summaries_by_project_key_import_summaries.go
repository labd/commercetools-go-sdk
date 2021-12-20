// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyImportSummariesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyImportSummariesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
