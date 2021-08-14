// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new product types or updating existing ones.
 */
func (rb *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body ProductTypeImportRequest) *ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyProductTypesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
