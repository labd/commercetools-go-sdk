package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Variants.
*
*	This endpoint is only available for Projects with `productCatalogModel` set to `Modular`.
*
 */
func (rb *ByProjectKeyVariantsImportContainersByImportContainerKeyRequestBuilder) Post(body VariantImportRequest) *ByProjectKeyVariantsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyVariantsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
