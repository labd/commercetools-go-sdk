package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImportContainersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyImportContainersRequestBuilder) WithImportContainerKeyValue(importContainerKey string) *ByProjectKeyImportContainersByImportContainerKeyRequestBuilder {
	return &ByProjectKeyImportContainersByImportContainerKeyRequestBuilder{
		importContainerKey: importContainerKey,
		projectKey:         rb.projectKey,
		client:             rb.client,
	}
}

/**
*	Creates an Import Container in the Project.
*
*	Generates the [ImportContainerCreated](/projects/events#import-container-created-event) Event.
*
 */
func (rb *ByProjectKeyImportContainersRequestBuilder) Post(body ImportContainerDraft) *ByProjectKeyImportContainersRequestMethodPost {
	return &ByProjectKeyImportContainersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/import-containers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Retrieves all ImportContainers in the Project.
 */
func (rb *ByProjectKeyImportContainersRequestBuilder) Get() *ByProjectKeyImportContainersRequestMethodGet {
	return &ByProjectKeyImportContainersRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers", rb.projectKey),
		client: rb.client,
	}
}
