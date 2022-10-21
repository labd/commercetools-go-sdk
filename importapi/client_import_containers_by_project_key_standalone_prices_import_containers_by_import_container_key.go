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
*	Creates a request for creating new Standalone Prices or updating existing ones.
 */
func (rb *ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestBuilder) Post(body StandalonePriceImportRequest) *ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyStandalonePricesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/standalone-prices/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
