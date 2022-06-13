package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductsRequestBuilder) ImportContainers() *ByProjectKeyProductsImportContainersRequestBuilder {
	return &ByProjectKeyProductsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductsRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
