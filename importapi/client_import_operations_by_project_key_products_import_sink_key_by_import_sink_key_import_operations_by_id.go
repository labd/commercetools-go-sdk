package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder struct {
	projectKey    string
	importSinkKey string
	id            string
	client        *Client
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/importSinkKey=%s/import-operations/%s", rb.projectKey, rb.importSinkKey, rb.id),
		client: rb.client,
	}
}
