package importapi

// Generated file, please do not change!!!

type ByProjectKeyBusinessUnitsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsRequestBuilder) ImportContainers() *ByProjectKeyBusinessUnitsImportContainersRequestBuilder {
	return &ByProjectKeyBusinessUnitsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
