package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersSearchRequestBuilder) Post(body OrderSearchRequest) *ByProjectKeyOrdersSearchRequestMethodPost {
	return &ByProjectKeyOrdersSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/search", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks whether a search index for the Project's Orders exists.
*	Returns a `200 OK` status if the index exists or a `404 Not Found` error otherwise.
*
 */
func (rb *ByProjectKeyOrdersSearchRequestBuilder) Head() *ByProjectKeyOrdersSearchRequestMethodHead {
	return &ByProjectKeyOrdersSearchRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/search", rb.projectKey),
		client: rb.client,
	}
}
