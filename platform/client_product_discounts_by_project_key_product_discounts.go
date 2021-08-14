// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductDiscountsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductDiscountsRequestBuilder) Matching() *ByProjectKeyProductDiscountsMatchingRequestBuilder {
	return &ByProjectKeyProductDiscountsMatchingRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductDiscountsRequestBuilder) WithKey(key string) *ByProjectKeyProductDiscountsKeyByKeyRequestBuilder {
	return &ByProjectKeyProductDiscountsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductDiscountsRequestBuilder) WithId(id string) *ByProjectKeyProductDiscountsByIdRequestBuilder {
	return &ByProjectKeyProductDiscountsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query product-discounts
 */
func (rb *ByProjectKeyProductDiscountsRequestBuilder) Get() *ByProjectKeyProductDiscountsRequestMethodGet {
	return &ByProjectKeyProductDiscountsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-discounts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create ProductDiscount
 */
func (rb *ByProjectKeyProductDiscountsRequestBuilder) Post(body ProductDiscountDraft) *ByProjectKeyProductDiscountsRequestMethodPost {
	return &ByProjectKeyProductDiscountsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts", rb.projectKey),
		client: rb.client,
	}
}
