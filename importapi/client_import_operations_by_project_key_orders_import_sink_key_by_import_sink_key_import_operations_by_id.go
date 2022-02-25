package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder struct {
	projectKey    string
	importSinkKey string
	id            string
	client        *Client
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/importSinkKey=%s/import-operations/%s", rb.projectKey, rb.importSinkKey, rb.id),
		client: rb.client,
	}
}
