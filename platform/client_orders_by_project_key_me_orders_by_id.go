package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Get() *ByProjectKeyMeOrdersByIDRequestMethodGet {
	return &ByProjectKeyMeOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `id`. Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeOrdersByIDRequestBuilder) Head() *ByProjectKeyMeOrdersByIDRequestMethodHead {
	return &ByProjectKeyMeOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
