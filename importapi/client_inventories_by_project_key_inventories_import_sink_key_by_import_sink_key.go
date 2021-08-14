// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates a request for creating new Inventories or updating existing ones.
 */
func (rb *ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body InventoryImportRequest) *ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyInventoriesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/inventories/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
