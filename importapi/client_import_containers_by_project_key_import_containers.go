// Generated file, please do not change!!!
package importapi

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
*	Creates a new import container.
 */
func (rb *ByProjectKeyImportContainersRequestBuilder) Post(body ImportContainerDraft) *ByProjectKeyImportContainersRequestMethodPost {
	return &ByProjectKeyImportContainersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/import-containers", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Retrieves all import containers of a given project key.
 */
func (rb *ByProjectKeyImportContainersRequestBuilder) Get() *ByProjectKeyImportContainersRequestMethodGet {
	return &ByProjectKeyImportContainersRequestMethodGet{
		url:    fmt.Sprintf("/%s/import-containers", rb.projectKey),
		client: rb.client,
	}
}
