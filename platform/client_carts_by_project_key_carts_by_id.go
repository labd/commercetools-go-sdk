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
*	Checks if a Cart exists for a given `id`. Returns a `200 OK` status if the Cart exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartsByIDRequestBuilder) Head() *ByProjectKeyCartsByIDRequestMethodHead {
	return &ByProjectKeyCartsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsByIDRequestBuilder) Post(body CartUpdate) *ByProjectKeyCartsByIDRequestMethodPost {
	return &ByProjectKeyCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartsByIDRequestBuilder) Delete() *ByProjectKeyCartsByIDRequestMethodDelete {
	return &ByProjectKeyCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/carts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
