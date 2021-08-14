// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeOrdersByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Order by ID
 */
func (rb *ByProjectKeyMeOrdersByIdRequestBuilder) Get() *ByProjectKeyMeOrdersByIdRequestMethodGet {
	return &ByProjectKeyMeOrdersByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
