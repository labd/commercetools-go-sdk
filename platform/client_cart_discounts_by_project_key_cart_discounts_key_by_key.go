// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCartDiscountsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get CartDiscount by key
 */
func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Get() *ByProjectKeyCartDiscountsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update CartDiscount by key
 */
func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyCartDiscountsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete CartDiscount by key
 */
func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCartDiscountsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
