// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a new import request for order patches
 */
func (rb *ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestBuilder) Post(body OrderPatchImportRequest) *ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyOrderPatchesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/order-patches/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
