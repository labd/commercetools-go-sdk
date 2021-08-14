// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates import request for creating new product variants or updating existing ones.
 */
func (rb *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestBuilder) Post(body ProductVariantImportRequest) *ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyProductVariantsImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variants/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
