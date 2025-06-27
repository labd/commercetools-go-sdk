package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Products.
 */
func (rb *ByProjectKeyProductsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductImportRequest) *ByProjectKeyProductsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
