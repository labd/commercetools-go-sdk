// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Retrieves all [ImportOperations](ctp:import:type:ImportOperation) of a given ImportContainer key.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s/import-operations", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
