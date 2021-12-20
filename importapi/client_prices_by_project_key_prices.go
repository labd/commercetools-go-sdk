// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyPricesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPricesRequestBuilder) ImportContainers() *ByProjectKeyPricesImportContainersRequestBuilder {
	return &ByProjectKeyPricesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyPricesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
