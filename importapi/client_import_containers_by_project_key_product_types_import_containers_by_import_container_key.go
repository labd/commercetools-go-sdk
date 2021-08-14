// Generated file, please do not change!!!
package importapi

import (
	"fmt"
)

type ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates a request for creating new ProductTypes or updating existing ones.
 */
func (rb *ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder) Post(body ProductTypeImportRequest) *ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
