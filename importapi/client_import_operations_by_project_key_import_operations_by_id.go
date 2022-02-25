package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImportOperationsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves the [ImportOperation](ctp:import:type:ImportOperation) of a given ID.
*
 */
func (rb *ByProjectKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-operations/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
