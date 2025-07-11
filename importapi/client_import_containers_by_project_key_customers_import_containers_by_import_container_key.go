package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Customers.
 */
func (rb *ByProjectKeyCustomersImportContainersByImportContainerKeyRequestBuilder) Post(body CustomerImportRequest) *ByProjectKeyCustomersImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyCustomersImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
