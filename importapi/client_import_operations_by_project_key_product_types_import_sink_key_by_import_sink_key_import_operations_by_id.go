package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder struct {
	projectKey    string
	importSinkKey string
	id            string
	client        *Client
}

/**
*	Retrieves the import operation with the given id.
*
 */
func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder) Get() *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/importSinkKey=%s/import-operations/%s", rb.projectKey, rb.importSinkKey, rb.id),
		client: rb.client,
	}
}
