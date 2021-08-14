// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyDiscountCodesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get DiscountCode by ID
 */
func (rb *ByProjectKeyDiscountCodesByIdRequestBuilder) Get() *ByProjectKeyDiscountCodesByIdRequestMethodGet {
	return &ByProjectKeyDiscountCodesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update DiscountCode by ID
 */
func (rb *ByProjectKeyDiscountCodesByIdRequestBuilder) Post(body DiscountCodeUpdate) *ByProjectKeyDiscountCodesByIdRequestMethodPost {
	return &ByProjectKeyDiscountCodesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete DiscountCode by ID
 */
func (rb *ByProjectKeyDiscountCodesByIdRequestBuilder) Delete() *ByProjectKeyDiscountCodesByIdRequestMethodDelete {
	return &ByProjectKeyDiscountCodesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
