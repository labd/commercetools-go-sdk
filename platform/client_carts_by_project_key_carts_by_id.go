// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	The cart may not contain up-to-date prices, discounts etc.
*	If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyCartsByIdRequestBuilder) Get() *ByProjectKeyCartsByIdRequestMethodGet {
	return &ByProjectKeyCartsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Cart by ID
 */
func (rb *ByProjectKeyCartsByIdRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsByIdRequestMethodPost {
	return &ByProjectKeyCartsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Cart by ID
 */
func (rb *ByProjectKeyCartsByIdRequestBuilder) Delete() *ByProjectKeyCartsByIdRequestMethodDelete {
	return &ByProjectKeyCartsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
