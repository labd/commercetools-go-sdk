package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductDiscountsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductDiscountsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductDiscountsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductDiscountsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ProductDiscount exists with the provided `key`. Returns a `200` status if the ProductDiscount exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyProductDiscountsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductDiscountsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductDiscountsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductDiscountsKeyByKeyRequestBuilder) Post(body ProductDiscountUpdate) *ByProjectKeyProductDiscountsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductDiscountsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductDiscountsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductDiscountsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductDiscountsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
