package importapi

// Generated file, please do not change!!!

type ByProjectKeyVariantsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyVariantsRequestBuilder) ImportContainers() *ByProjectKeyVariantsImportContainersRequestBuilder {
	return &ByProjectKeyVariantsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
