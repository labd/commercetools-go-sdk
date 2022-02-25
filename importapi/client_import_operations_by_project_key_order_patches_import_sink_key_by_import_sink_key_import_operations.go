package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all order-patches import operations of an import sink key.
 */
func (rb *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyOrderPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/order-patches/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
