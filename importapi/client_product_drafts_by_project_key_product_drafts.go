package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductDraftsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductDraftsRequestBuilder) ImportContainers() *ByProjectKeyProductDraftsImportContainersRequestBuilder {
	return &ByProjectKeyProductDraftsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
