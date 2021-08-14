// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all Price ImportOperations of a given ImportSink key.
*
 */
func (rb *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyPricesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/prices/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
