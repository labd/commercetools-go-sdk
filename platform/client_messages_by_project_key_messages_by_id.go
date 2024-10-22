package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMessagesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Deprecated OAuth 2.0 Scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyMessagesByIDRequestBuilder) Get() *ByProjectKeyMessagesByIDRequestMethodGet {
	return &ByProjectKeyMessagesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/messages/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Message exists for a given `id`. Returns a `200 OK` status if the Message exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMessagesByIDRequestBuilder) Head() *ByProjectKeyMessagesByIDRequestMethodHead {
	return &ByProjectKeyMessagesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/messages/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
