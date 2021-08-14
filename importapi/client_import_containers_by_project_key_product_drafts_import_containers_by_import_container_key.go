// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new ProductDrafts or updating existing ones.
*
 */
func (rb *ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductDraftImportRequest) *ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductDraftsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-drafts/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
