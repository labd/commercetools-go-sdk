package importapi

// Generated file, please do not change!!!

type ByProjectKeyCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCategoriesRequestBuilder) ImportContainers() *ByProjectKeyCategoriesImportContainersRequestBuilder {
	return &ByProjectKeyCategoriesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
