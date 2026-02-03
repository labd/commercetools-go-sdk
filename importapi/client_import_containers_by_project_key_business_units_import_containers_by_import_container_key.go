package importapi

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestBuilder struct {
	projectKey         string
	importContainerKey string
	client             *Client
}

/**
*	Creates an Import Request for Business Units.
 */
func (rb *ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestBuilder) Post(body BusinessUnitImportRequest) *ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestMethodPost {
	return &ByProjectKeyBusinessUnitsImportContainersByImportContainerKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units/import-containers/%s", rb.projectKey, rb.importContainerKey),
		client: rb.client,
	}
}
