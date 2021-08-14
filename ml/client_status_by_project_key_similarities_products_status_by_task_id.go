// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeySimilaritiesProductsStatusByTaskIdRequestBuilder struct {
	projectKey string
	taskId     string
	client     *Client
}

func (rb *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestBuilder) Get() *ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet {
	return &ByProjectKeySimilaritiesProductsStatusByTaskIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/similarities/products/status/%s", rb.projectKey, rb.taskId),
		client: rb.client,
	}
}
