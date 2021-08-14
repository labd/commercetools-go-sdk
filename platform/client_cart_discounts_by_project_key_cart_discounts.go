// Generated file, please do not change!!!
package platform

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
func (rb *ByProjectKeyCartDiscountsRequestBuilder) WithId(id string) *ByProjectKeyCartDiscountsByIdRequestBuilder {
	return &ByProjectKeyCartDiscountsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query cart-discounts
 */
func (rb *ByProjectKeyCartDiscountsRequestBuilder) Get() *ByProjectKeyCartDiscountsRequestMethodGet {
	return &ByProjectKeyCartDiscountsRequestMethodGet{
		url:    fmt.Sprintf("/%s/cart-discounts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create CartDiscount
 */
func (rb *ByProjectKeyCartDiscountsRequestBuilder) Post(body CartDiscountDraft) *ByProjectKeyCartDiscountsRequestMethodPost {
	return &ByProjectKeyCartDiscountsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/cart-discounts", rb.projectKey),
		client: rb.client,
	}
}
