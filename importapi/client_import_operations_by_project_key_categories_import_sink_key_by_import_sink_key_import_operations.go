package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all category import operations of an import sink key.
*
 */
func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
