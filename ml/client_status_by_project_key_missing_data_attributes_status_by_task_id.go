// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeyMissingDataAttributesStatusByTaskIdRequestBuilder struct {
	projectKey string
	taskId     string
	client     *Client
}

func (rb *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestBuilder) Get() *ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet {
	return &ByProjectKeyMissingDataAttributesStatusByTaskIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/missing-data/attributes/status/%s", rb.projectKey, rb.taskId),
		client: rb.client,
	}
}
