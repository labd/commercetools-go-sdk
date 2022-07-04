package importapi

// Generated file, please do not change!!!

type ByProjectKeyInventoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyInventoriesRequestBuilder) ImportContainers() *ByProjectKeyInventoriesImportContainersRequestBuilder {
	return &ByProjectKeyInventoriesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
