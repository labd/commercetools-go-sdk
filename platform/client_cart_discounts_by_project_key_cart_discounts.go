package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartDiscountsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCartDiscountsRequestBuilder) WithKey(key string) *ByProjectKeyCartDiscountsKeyByKeyRequestBuilder {
	return &ByProjectKeyCartDiscountsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartDiscountsRequestBuilder) WithId(id string) *ByProjectKeyCartDiscountsByIDRequestBuilder {
	return &ByProjectKeyCartDiscountsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartDiscountsRequestBuilder) Get() *ByProjectKeyCartDiscountsRequestMethodGet {
	return &ByProjectKeyCartDiscountsRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a CartDiscount exists for a given Query Predicate. Returns a `200 OK` status if any CartDiscounts match the Query Predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartDiscountsRequestBuilder) Head() *ByProjectKeyCartDiscountsRequestMethodHead {
	return &ByProjectKeyCartDiscountsRequestMethodHead{
		url:    fmt.Sprintf("/%s/cart-discounts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a Cart Discount produces the [CartDiscountCreated](ctp:api:type:CartDiscountCreatedMessage) Message.
 */
func (rb *ByProjectKeyCartDiscountsRequestBuilder) Post(body CartDiscountDraft) *ByProjectKeyCartDiscountsRequestMethodPost {
	return &ByProjectKeyCartDiscountsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts", rb.projectKey),
		client: rb.client,
	}
}
