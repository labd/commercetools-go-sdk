package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	If the initial indexing is in progress or the feature is inactive, a [SearchNotReady](ctp:api:type:SearchNotReadyError) error is returned.
*	If inactive, you can [reactivate](/../api/projects/customer-search#reactivate) it.
*
 */
func (rb *ByProjectKeyCustomersSearchRequestBuilder) Post(body CustomerSearchRequest) *ByProjectKeyCustomersSearchRequestMethodPost {
	return &ByProjectKeyCustomersSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks whether a search index of Customers exists for a Project.
*	Returns a `200 OK` if an index exists; otherwise, returns a `409 Conflict`.
*
 */
func (rb *ByProjectKeyCustomersSearchRequestBuilder) Head() *ByProjectKeyCustomersSearchRequestMethodHead {
	return &ByProjectKeyCustomersSearchRequestMethodHead{
		url:    fmt.Sprintf("/%s/customers/search", rb.projectKey),
		client: rb.client,
	}
}
