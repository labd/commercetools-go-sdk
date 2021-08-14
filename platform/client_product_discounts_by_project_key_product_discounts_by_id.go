// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductDiscountsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get ProductDiscount by ID
 */
func (rb *ByProjectKeyProductDiscountsByIdRequestBuilder) Get() *ByProjectKeyProductDiscountsByIdRequestMethodGet {
	return &ByProjectKeyProductDiscountsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update ProductDiscount by ID
 */
func (rb *ByProjectKeyProductDiscountsByIdRequestBuilder) Post(body ProductDiscountUpdate) *ByProjectKeyProductDiscountsByIdRequestMethodPost {
	return &ByProjectKeyProductDiscountsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ProductDiscount by ID
 */
func (rb *ByProjectKeyProductDiscountsByIdRequestBuilder) Delete() *ByProjectKeyProductDiscountsByIdRequestMethodDelete {
	return &ByProjectKeyProductDiscountsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
