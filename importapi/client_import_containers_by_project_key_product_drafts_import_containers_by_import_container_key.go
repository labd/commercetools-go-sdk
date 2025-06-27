package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Products.
*
 */
func (rb *ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductDraftImportRequest) *ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-drafts/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
