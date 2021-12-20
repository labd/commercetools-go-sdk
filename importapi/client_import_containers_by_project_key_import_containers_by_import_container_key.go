// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) ImportSummaries() *ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder {
	return &ByProjectKeyImportContainersByImportContainerKeyImportSummariesRequestBuilder{
		projectKey:         rb.projectKey,
		importContainerKey: rb.importContainerKey,
		client:             rb.client,
	}
}
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) ImportOperations() *ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder {
	return &ByProjectKeyImportContainersByImportContainerKeyImportOperationsRequestBuilder{
		projectKey:         rb.projectKey,
		importContainerKey: rb.importContainerKey,
		client:             rb.client,
	}
}

/**
*	Updates the import container given by the key.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Put(body ImportContainerUpdateDraft) *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}

/**
*	Retrieves the import container given by the key.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}

/**
*	Deletes the import container given by the key.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Delete() *ByProjectKeyImportContainersByImportContainerKeyRequestMethodDelete {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
