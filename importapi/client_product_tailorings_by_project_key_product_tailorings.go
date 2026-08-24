package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductTailoringsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTailoringsRequestBuilder) ImportContainers() *ByProjectKeyProductTailoringsImportContainersRequestBuilder {
	return &ByProjectKeyProductTailoringsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
