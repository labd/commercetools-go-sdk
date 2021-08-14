// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCategoriesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
