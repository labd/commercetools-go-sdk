package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductSelectionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsRequestBuilder) ImportContainers() *ByProjectKeyProductSelectionsImportContainersRequestBuilder {
	return &ByProjectKeyProductSelectionsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
