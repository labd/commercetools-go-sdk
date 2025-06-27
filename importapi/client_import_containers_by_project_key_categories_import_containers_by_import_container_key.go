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
*	Creates an Import Request for Categories.
 */
func (rb *ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestBuilder) Post(body CategoryImportRequest) *ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyCategoriesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
