// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	The cart may not contain up-to-date prices, discounts etc.
*	If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Get() *ByProjectKeyCartsByIDRequestMethodGet {
	return &ByProjectKeyCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsByIDRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsByIDRequestMethodPost {
	return &ByProjectKeyCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsByIDRequestBuilder) Delete() *ByProjectKeyCartsByIDRequestMethodDelete {
	return &ByProjectKeyCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
