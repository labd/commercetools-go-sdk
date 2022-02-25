package ml

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMissingDataPricesStatusByTaskIdRequestBuilder struct {
	projectKey string
	taskId     string
	client     *Client
}

func (rb *ByProjectKeyMissingDataPricesStatusByTaskIdRequestBuilder) Get() *ByProjectKeyMissingDataPricesStatusByTaskIdRequestMethodGet {
	return &ByProjectKeyMissingDataPricesStatusByTaskIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/missing-data/prices/status/%s", rb.projectKey, rb.taskId),
		client: rb.client,
	}
}
