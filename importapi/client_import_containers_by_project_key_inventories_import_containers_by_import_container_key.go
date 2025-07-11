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
*	Creates an Import Request for InventoryEntries.
 */
func (rb *ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestBuilder) Post(body InventoryImportRequest) *ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyInventoriesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventories/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
