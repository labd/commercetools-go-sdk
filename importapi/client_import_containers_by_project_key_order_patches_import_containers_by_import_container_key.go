package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for updating Orders.
 */
func (rb *ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder) Post(body OrderPatchImportRequest) *ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/order-patches/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
