package importapi

// Generated file, please do not change!!!

type ByProjectKeyOrderPatchesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrderPatchesRequestBuilder) ImportContainers() *ByProjectKeyOrderPatchesImportContainersRequestBuilder {
	return &ByProjectKeyOrderPatchesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
