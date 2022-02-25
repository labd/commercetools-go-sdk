package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new categories or updating existing ones.
 */
func (rb *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body CategoryImportRequest) *ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyCategoriesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
