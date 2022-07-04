package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductVariantsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductVariantsRequestBuilder) ImportContainers() *ByProjectKeyProductVariantsImportContainersRequestBuilder {
	return &ByProjectKeyProductVariantsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
