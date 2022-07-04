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
