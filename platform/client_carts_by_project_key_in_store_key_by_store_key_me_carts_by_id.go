// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Get Cart by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Cart by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Cart by ID
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
