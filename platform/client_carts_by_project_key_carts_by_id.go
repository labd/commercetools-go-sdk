package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Cart with the provided `id`.
*	To ensure the Cart is up-to-date with current values (such as Prices and Discounts), use the [Recalculate](ctp:api:type:CartRecalculateAction) update action.
*
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Get() *ByProjectKeyCartsByIDRequestMethodGet {
	return &ByProjectKeyCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for the provided `id`. Returns a `200 OK` status if the Cart exists or [Not Found](/../api/errors#404-not-found) otherwise.
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Head() *ByProjectKeyCartsByIDRequestMethodHead {
	return &ByProjectKeyCartsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Cart in the Project using one or more [update actions](/../api/projects/carts#update-actions).
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsByIDRequestMethodPost {
	return &ByProjectKeyCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Cart in the Project.
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Delete() *ByProjectKeyCartsByIDRequestMethodDelete {
	return &ByProjectKeyCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
