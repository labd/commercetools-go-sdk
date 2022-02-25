package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all order import operations of an import sink key.
 */
func (rb *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
