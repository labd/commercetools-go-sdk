// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	The cart may not contain up-to-date prices, discounts etc.
*	If you want to ensure they're up-to-date, send an Update request with the Recalculate update action instead.
*
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
