// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Retrieves an [ImportSummary](ctp:import:type:ImportSummary) for the given import container. An [ImportSummary](ctp:import:type:ImportSummary) is calculated on demand.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s/import-summaries", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
