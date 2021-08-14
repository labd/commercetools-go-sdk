// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new products or updating existing ones.
 */
func (rb *ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestBuilder) Post(body ProductImportRequest) *ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyProductsImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
