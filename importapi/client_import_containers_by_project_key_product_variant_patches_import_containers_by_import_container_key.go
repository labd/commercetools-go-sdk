// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a new import request for product variant patches
 */
func (rb *ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder) Post(body ProductVariantPatchRequest) *ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variant-patches/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
