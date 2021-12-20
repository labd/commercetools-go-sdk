// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyImportSinksByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

/**
*	Updates the import sink given by the key.
 */
func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestBuilder) Put(body ImportSinkUpdateDraft) *ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut {
	return &ByProjectKeyImportSinksByImportSinkKeyRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/import-sinks/%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}

/**
*	Retrieves the import sink given by the key.
 */
func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestBuilder) Get() *ByProjectKeyImportSinksByImportSinkKeyRequestMethodGet {
	return &ByProjectKeyImportSinksByImportSinkKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-sinks/%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}

/**
*	Deletes the import sink given by the key.
 */
func (rb *ByProjectKeyImportSinksByImportSinkKeyRequestBuilder) Delete() *ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete {
	return &ByProjectKeyImportSinksByImportSinkKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/import-sinks/%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
