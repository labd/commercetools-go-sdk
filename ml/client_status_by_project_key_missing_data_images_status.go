package ml

// Generated file, please do not change!!!

type ByProjectKeyMissingDataImagesStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataImagesStatusRequestBuilder) WithTaskId(taskId string) *ByProjectKeyMissingDataImagesStatusByTaskIdRequestBuilder {
	return &ByProjectKeyMissingDataImagesStatusByTaskIdRequestBuilder{
		taskId:     taskId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
