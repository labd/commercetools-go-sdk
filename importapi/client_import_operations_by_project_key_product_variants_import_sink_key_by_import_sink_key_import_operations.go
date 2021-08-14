// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all product-variant import operations of an import sink key.
 */
func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-variants/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
