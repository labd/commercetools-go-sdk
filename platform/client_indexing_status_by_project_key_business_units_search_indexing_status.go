package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsSearchIndexingStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Returns the indexing status of the Business Unit Search for a Project.
*
 */
func (rb *ByProjectKeyBusinessUnitsSearchIndexingStatusRequestBuilder) Get() *ByProjectKeyBusinessUnitsSearchIndexingStatusRequestMethodGet {
	return &ByProjectKeyBusinessUnitsSearchIndexingStatusRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units/search/indexing-status", rb.projectKey),
		client: rb.client,
	}
}
