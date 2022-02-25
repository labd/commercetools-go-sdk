package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all import operations of an import sink key.
 */
func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-types/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
