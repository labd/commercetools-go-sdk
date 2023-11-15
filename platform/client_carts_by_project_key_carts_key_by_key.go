package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Get() *ByProjectKeyCartsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCartsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for a given `key`. Returns a `200 OK` status if the Cart exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Head() *ByProjectKeyCartsKeyByKeyRequestMethodHead {
	return &ByProjectKeyCartsKeyByKeyRequestMethodHead{
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
