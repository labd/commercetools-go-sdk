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
*	A [SearchNotReady](ctp:api:type:SearchNotReadyError) error is returned if the indexing is in progress or the feature is deactivated. If deactivated, you can [reactivate](/../api/projects/customer-search#reactivate) it.
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
*	If an index exists, a `200 OK` is returned; otherwise, a `409 Conflict`.
*
 */
func (rb *ByProjectKeyCustomersSearchRequestBuilder) Head() *ByProjectKeyCustomersSearchRequestMethodHead {
	return &ByProjectKeyCustomersSearchRequestMethodHead{
		url:    fmt.Sprintf("/%s/customers/search", rb.projectKey),
		client: rb.client,
	}
}
