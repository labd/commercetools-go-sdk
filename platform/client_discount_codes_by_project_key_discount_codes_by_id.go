package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountCodesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Deprecated OAuth 2.0 Scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Get() *ByProjectKeyDiscountCodesByIDRequestMethodGet {
	return &ByProjectKeyDiscountCodesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountCode exists for a given `id`. Returns a `200 OK` status if the DiscountCode exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Head() *ByProjectKeyDiscountCodesByIDRequestMethodHead {
	return &ByProjectKeyDiscountCodesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deprecated OAuth 2.0 Scope: `manage_orders:{projectKey}`
 */
func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Post(body DiscountCodeUpdate) *ByProjectKeyDiscountCodesByIDRequestMethodPost {
	return &ByProjectKeyDiscountCodesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deleting a Discount Code produces the [DiscountCodeDeleted](ctp:api:type:DiscountCodeDeletedMessage) Message.
*
*	Deprecated OAuth 2.0 Scope: `manage_orders:{projectKey}`
*
 */
func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Delete() *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	return &ByProjectKeyDiscountCodesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
