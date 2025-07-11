package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Prices.
 */
func (rb *ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder) Post(body PriceImportRequest) *ByProjectKeyPricesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyPricesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/prices/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
