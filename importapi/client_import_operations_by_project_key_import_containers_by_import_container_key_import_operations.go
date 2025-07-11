package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Retrieves all ImportOperations within an [ImportContainer](ctp:import:type:ImportContainer).
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s/import-operations", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
