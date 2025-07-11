package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Retrieves an [ImportSummary](ctp:import:type:ImportSummary) for the [ImportContainer](ctp:import:type:ImportContainer) with the provided `importContainerKey`.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s/import-summaries", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
