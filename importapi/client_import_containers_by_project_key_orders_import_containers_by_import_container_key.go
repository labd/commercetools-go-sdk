package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new Orders.
 */
func (rb *ByProjectKeyOrdersImportContainersByImportContainerKeyRequestBuilder) Post(body OrderImportRequest) *ByProjectKeyOrdersImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyOrdersImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
