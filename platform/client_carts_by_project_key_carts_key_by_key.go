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
*	Retrieves a Cart with the provided `key`.
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
*	Checks if a Cart exists with the provided `key`. Returns a `200 OK` status if the Cart exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Head() *ByProjectKeyCartsKeyByKeyRequestMethodHead {
	return &ByProjectKeyCartsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Cart in the Project using one or more [update actions](/../api/projects/carts#update-actions).
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCartsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Cart in the Project.
 */
func (rb *ByProjectKeyCartsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCartsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCartsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
