package importapi

// Generated file, please do not change!!!

type ByProjectKeyCustomersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersRequestBuilder) ImportContainers() *ByProjectKeyCustomersImportContainersRequestBuilder {
	return &ByProjectKeyCustomersImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
