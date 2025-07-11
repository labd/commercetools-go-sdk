package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTypesImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Types.
 */
func (rb *ByProjectKeyTypesImportContainersByImportContainerKeyRequestBuilder) Post(body TypeImportRequest) *ByProjectKeyTypesImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyTypesImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
