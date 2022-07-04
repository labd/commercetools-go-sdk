package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new Inventories or updating existing ones.
 */
func (rb *ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestBuilder) Post(body InventoryImportRequest) *ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventories/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
