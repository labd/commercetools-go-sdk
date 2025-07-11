package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartDiscountsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Get() *ByProjectKeyCartDiscountsByIDRequestMethodGet {
	return &ByProjectKeyCartDiscountsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a CartDiscount exists with the provided `id`. Returns a `200 OK` status if the CartDiscount exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Head() *ByProjectKeyCartDiscountsByIDRequestMethodHead {
	return &ByProjectKeyCartDiscountsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyCartDiscountsByIDRequestMethodPost {
	return &ByProjectKeyCartDiscountsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deleting a Cart Discount produces the [CartDiscountDeleted](ctp:api:type:CartDiscountDeletedMessage) Message.
 */
func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Delete() *ByProjectKeyCartDiscountsByIDRequestMethodDelete {
	return &ByProjectKeyCartDiscountsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
