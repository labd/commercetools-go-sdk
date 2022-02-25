package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) WithIdValue(id string) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder {
	return &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsByIdRequestBuilder{
		id:            id,
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Retrieves all import operations of an import sink key.
 */
func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder) Get() *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet {
	return &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-drafts/importSinkKey=%s/import-operations", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
