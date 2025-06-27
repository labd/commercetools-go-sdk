package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Product Selections.
 */
func (rb *ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestBuilder) Post(body ProductSelectionImportRequest) *ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductSelectionsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-selections/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
