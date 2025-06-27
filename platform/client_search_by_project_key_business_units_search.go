package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	If the initial indexing is in progress or the feature is inactive, A [SearchNotReady](ctp:api:type:SearchNotReadyError) error is returned. If inactive, you can [reactivate](/../api/projects/business-unit-search#reactivate) it.
*
 */
func (rb *ByProjectKeyBusinessUnitsSearchRequestBuilder) Post(body BusinessUnitSearchRequest) *ByProjectKeyBusinessUnitsSearchRequestMethodPost {
	return &ByProjectKeyBusinessUnitsSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks whether a search index of Business Units exists for a Project.
*	Returns a `200 OK` if an index exists; otherwise, returns a `409 Conflict`.
*
 */
func (rb *ByProjectKeyBusinessUnitsSearchRequestBuilder) Head() *ByProjectKeyBusinessUnitsSearchRequestMethodHead {
	return &ByProjectKeyBusinessUnitsSearchRequestMethodHead{
		url:    fmt.Sprintf("/%s/business-units/search", rb.projectKey),
		client: rb.client,
	}
}
