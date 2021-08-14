// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyProductVariantsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductVariantsRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
