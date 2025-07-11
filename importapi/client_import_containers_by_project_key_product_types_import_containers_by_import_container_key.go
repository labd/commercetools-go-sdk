package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for ProductTypes.
 */
func (rb *ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestBuilder) Post(body ProductTypeImportRequest) *ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyProductTypesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-types/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
