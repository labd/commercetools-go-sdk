// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all product import operations of an import sink key.
 */
func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
