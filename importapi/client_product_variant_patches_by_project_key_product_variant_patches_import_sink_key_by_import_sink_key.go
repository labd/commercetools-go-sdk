package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestBuilder struct {
	projectKey    string
	importSinkKey string
	client        *Client
}

func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestBuilder) ImportOperations() *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder {
	return &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyImportOperationsRequestBuilder{
		projectKey:    rb.projectKey,
		importSinkKey: rb.importSinkKey,
		client:        rb.client,
	}
}

/**
*	Creates a new import request for product variant patches
 */
func (rb *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestBuilder) Post(body ProductVariantPatchRequest) *ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestMethodPost {
	return &ByProjectKeyProductVariantPatchesImportSinkKeyByImportSinkKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variant-patches/importSinkKey=%s", rb.projectKey, rb.importSinkKey),
		client: rb.client,
	}
}
