package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for ProductVariants.
 */
func (rb *ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductVariantImportRequest) *ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variants/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
