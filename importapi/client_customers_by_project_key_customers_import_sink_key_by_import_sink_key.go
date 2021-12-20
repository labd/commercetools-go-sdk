// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyCustomersImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new customers or updating existing ones.
 */
func (rb *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestBuilder) Post(body CustomerImportRequest) *ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyCustomersImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
