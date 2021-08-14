// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartDiscountsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get CartDiscount by ID
 */
func (rb *ByProjectKeyCartDiscountsByIdRequestBuilder) Get() *ByProjectKeyCartDiscountsByIdRequestMethodGet {
	return &ByProjectKeyCartDiscountsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update CartDiscount by ID
 */
func (rb *ByProjectKeyCartDiscountsByIdRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyCartDiscountsByIdRequestMethodPost {
	return &ByProjectKeyCartDiscountsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete CartDiscount by ID
 */
func (rb *ByProjectKeyCartDiscountsByIdRequestBuilder) Delete() *ByProjectKeyCartDiscountsByIdRequestMethodDelete {
	return &ByProjectKeyCartDiscountsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
