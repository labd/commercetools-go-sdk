package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for updating Product Variants.
*	Returns an [InvalidField](ctp:import:type:InvalidFieldError) error if the [ProductVariantPatchRequest](ctp:import:type:ProductVariantPatchRequest) contains patches with and without the `product` field set.
*
 */
func (rb *ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestBuilder) Post(body ProductVariantPatchRequest) *ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductVariantPatchesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variant-patches/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
