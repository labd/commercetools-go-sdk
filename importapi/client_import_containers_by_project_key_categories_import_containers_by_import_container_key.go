package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new Categories or updating existing ones.
 */
func (rb *ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestBuilder) Post(body CategoryImportRequest) *ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
