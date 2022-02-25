package ml

// Generated file, please do not change!!!

import ()

type ByProjectKeyMissingDataPricesStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataPricesStatusRequestBuilder) WithTaskId(taskId string) *ByProjectKeyMissingDataPricesStatusByTaskIdRequestBuilder {
	return &ByProjectKeyMissingDataPricesStatusByTaskIdRequestBuilder{
		taskId:     taskId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
