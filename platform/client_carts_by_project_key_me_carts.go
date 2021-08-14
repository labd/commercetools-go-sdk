// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeCartsRequestBuilder) WithKey(key string) *ByProjectKeyMeCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeCartsRequestBuilder) WithId(id string) *ByProjectKeyMeCartsByIdRequestBuilder {
	return &ByProjectKeyMeCartsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query carts
 */
func (rb *ByProjectKeyMeCartsRequestBuilder) Get() *ByProjectKeyMeCartsRequestMethodGet {
	return &ByProjectKeyMeCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create Cart
 */
func (rb *ByProjectKeyMeCartsRequestBuilder) Post(body MyCartDraft) *ByProjectKeyMeCartsRequestMethodPost {
	return &ByProjectKeyMeCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts", rb.projectKey),
		client: rb.client,
	}
}
