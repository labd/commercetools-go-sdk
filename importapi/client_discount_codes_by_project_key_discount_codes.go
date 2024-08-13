package importapi

// Generated file, please do not change!!!

type ByProjectKeyDiscountCodesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDiscountCodesRequestBuilder) ImportContainers() *ByProjectKeyDiscountCodesImportContainersRequestBuilder {
	return &ByProjectKeyDiscountCodesImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
