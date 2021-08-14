// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeCartsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Cart by ID
 */
func (rb *ByProjectKeyMeCartsByIdRequestBuilder) Get() *ByProjectKeyMeCartsByIdRequestMethodGet {
	return &ByProjectKeyMeCartsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Cart by ID
 */
func (rb *ByProjectKeyMeCartsByIdRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyMeCartsByIdRequestMethodPost {
	return &ByProjectKeyMeCartsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Cart by ID
 */
func (rb *ByProjectKeyMeCartsByIdRequestBuilder) Delete() *ByProjectKeyMeCartsByIdRequestMethodDelete {
	return &ByProjectKeyMeCartsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
