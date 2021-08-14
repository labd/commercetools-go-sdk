// Generated file, please do not change!!!
package ml

import ()

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
