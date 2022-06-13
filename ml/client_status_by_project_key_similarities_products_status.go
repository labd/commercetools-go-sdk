package ml

// Generated file, please do not change!!!

type ByProjectKeySimilaritiesProductsStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeySimilaritiesProductsStatusRequestBuilder) WithTaskId(taskId string) *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestBuilder {
	return &ByProjectKeySimilaritiesProductsStatusByTaskIdRequestBuilder{
		taskId:     taskId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
