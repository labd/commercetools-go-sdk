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
*	Checks if a CartDiscount exists for a given `id`. Returns a `200 OK` status if the CartDiscount exists or a `404 Not Found` otherwise.
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

func (rb *ByProjectKeyCartDiscountsByIDRequestBuilder) Delete() *ByProjectKeyCartDiscountsByIDRequestMethodDelete {
	return &ByProjectKeyCartDiscountsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/cart-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
