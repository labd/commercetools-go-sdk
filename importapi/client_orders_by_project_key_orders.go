package importapi

// Generated file, please do not change!!!

type ByProjectKeyOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersRequestBuilder) ImportContainers() *ByProjectKeyOrdersImportContainersRequestBuilder {
	return &ByProjectKeyOrdersImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
