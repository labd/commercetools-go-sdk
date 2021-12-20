// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new product drafts or updating existing ones.
*
 */
func (rb *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestBuilder) Post(body ProductDraftImportRequest) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-drafts/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
