package importapi

// Generated file, please do not change!!!

type ByProjectKeyTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTypesRequestBuilder) ImportContainers() *ByProjectKeyTypesImportContainersRequestBuilder {
	return &ByProjectKeyTypesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
