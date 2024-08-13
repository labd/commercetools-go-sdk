package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersSearchRequestBuilder) Post(body CustomerSearchRequest) *ByProjectKeyCustomersSearchRequestMethodPost {
	return &ByProjectKeyCustomersSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks whether a search index for the Project's Customers exists.
*	Returns a `200 OK` status if the index exists or `404 Not Found` otherwise.
*
 */
func (rb *ByProjectKeyCustomersSearchRequestBuilder) Head() *ByProjectKeyCustomersSearchRequestMethodHead {
	return &ByProjectKeyCustomersSearchRequestMethodHead{
		url:    fmt.Sprintf("/%s/customers/search", rb.projectKey),
		client: rb.client,
	}
}
