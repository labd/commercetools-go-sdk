// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all Customer ImportOperations of a given ImportSink key.
*
 */
func (rb *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
