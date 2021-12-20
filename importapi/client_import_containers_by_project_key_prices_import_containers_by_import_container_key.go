// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new Prices or updating existing ones.
 */
func (rb *ByProjectKeyPricesImportContainersByImportContainerKeyRequestBuilder) Post(body PriceImportRequest) *ByProjectKeyPricesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyPricesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/prices/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
