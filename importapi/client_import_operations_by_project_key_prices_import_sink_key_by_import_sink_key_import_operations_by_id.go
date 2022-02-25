package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder struct {
	projectKey    string
	importSinkKey string
	id            string
	client        *Client
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/prices/importSinkKey=%s/import-operations/%s", rb.projectKey, rb.importSinkKey, rb.id),
		client: rb.client,
	}
}
