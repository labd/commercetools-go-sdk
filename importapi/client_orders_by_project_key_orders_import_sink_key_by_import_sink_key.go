// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new orders or updating existing ones.
 */
func (rb *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestBuilder) Post(body OrderImportRequest) *ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyOrdersImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
