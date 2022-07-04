package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTypesRequestBuilder) ImportContainers() *ByProjectKeyProductTypesImportContainersRequestBuilder {
	return &ByProjectKeyProductTypesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
