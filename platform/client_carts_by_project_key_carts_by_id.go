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
