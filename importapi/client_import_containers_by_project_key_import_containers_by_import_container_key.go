package importapi

// Generated file, please do not change!!!

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
*	Updates an [ImportContainer](ctp:import:type:ImportContainer) in the Project.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Put(body ImportContainerUpdateDraft) *ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}

/**
*	Retrieves an [ImportContainer](ctp:import:type:ImportContainer) with the provided `importContainerKey`.
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Get() *ByProjectKeyImportContainersByImportContainerKeyRequestMethodGet {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}

/**
*	Deletes an Import Container in the Project.
*
*	Generates the [ImportContainerDeleted](/projects/events#import-container-deleted-event) Event.
*
 */
func (rb *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder) Delete() *ByProjectKeyImportContainersByImportContainerKeyRequestMethodDelete {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
