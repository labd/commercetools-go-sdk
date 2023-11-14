package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeCartsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for a given `key`. Returns a `200 OK` status if the Cart exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Head() *ByProjectKeyMeCartsKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyMeCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
