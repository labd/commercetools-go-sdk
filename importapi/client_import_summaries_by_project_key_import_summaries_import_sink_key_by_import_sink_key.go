package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

/**
*	Retrieves an import summary for the given import sink.
*
*	The import summary is calculated on demand.
*
 */
func (rb *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestBuilder) Get() *ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet {
	return &ByProjectKeyImportSummariesImportSinkKeyByImportSinkKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-summaries/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
