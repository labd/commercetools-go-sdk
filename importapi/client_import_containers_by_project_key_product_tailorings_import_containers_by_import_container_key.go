package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Product Tailoring.
 */
func (rb *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductTailoringImportRequest) *ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductTailoringsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-tailorings/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
