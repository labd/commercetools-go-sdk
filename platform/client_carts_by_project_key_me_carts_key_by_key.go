// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeCartsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Cart by key
 */
func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Cart by key
 */
func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyMeCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Cart by key
 */
func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
