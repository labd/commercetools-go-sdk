package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartDiscountsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Get() *ByProjectKeyCartDiscountsKeyByKeyRequestMethodGet {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a CartDiscount exists with the provided `key`. Returns a `200 OK` status if the CartDiscount exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Head() *ByProjectKeyCartDiscountsKeyByKeyRequestMethodHead {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyCartDiscountsKeyByKeyRequestMethodPost {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deleting a Cart Discount produces the [CartDiscountDeleted](ctp:api:type:CartDiscountDeletedMessage) Message.
 */
func (rb *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder) Delete() *ByProjectKeyCartDiscountsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
