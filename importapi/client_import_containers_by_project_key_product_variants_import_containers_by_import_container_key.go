// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new ProductVariants or updating existing ones.
 */
func (rb *ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductVariantImportRequest) *ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductVariantsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-variants/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
