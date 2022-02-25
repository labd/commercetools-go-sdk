package importapi

// Generated file, please do not change!!!

import ()

type ByProjectKeyProductVariantPatchesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductVariantPatchesRequestBuilder) ImportContainers() *ByProjectKeyProductVariantPatchesImportContainersRequestBuilder {
	return &ByProjectKeyProductVariantPatchesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductVariantPatchesRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
