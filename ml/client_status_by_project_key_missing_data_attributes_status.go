package ml

// Generated file, please do not change!!!

type ByProjectKeyMissingDataAttributesStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataAttributesStatusRequestBuilder) WithTaskId(taskId string) *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestBuilder {
	return &ByProjectKeyMissingDataAttributesStatusByTaskIdRequestBuilder{
		taskId:     taskId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
