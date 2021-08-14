// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyProductsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductsRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
