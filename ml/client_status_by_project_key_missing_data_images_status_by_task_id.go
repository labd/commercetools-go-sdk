// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeyMissingDataImagesStatusByTaskIdRequestBuilder struct {
	projectKey string
	taskId     string
	client     *Client
}

func (rb *ByProjectKeyMissingDataImagesStatusByTaskIdRequestBuilder) Get() *ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet {
	return &ByProjectKeyMissingDataImagesStatusByTaskIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/missing-data/images/status/%s", rb.projectKey, rb.taskId),
		client: rb.client,
	}
}
