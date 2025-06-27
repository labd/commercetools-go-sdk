package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Standalone Prices.
 */
func (rb *ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestBuilder) Post(body StandalonePriceImportRequest) *ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
