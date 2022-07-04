package importapi

// Generated file, please do not change!!!

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
