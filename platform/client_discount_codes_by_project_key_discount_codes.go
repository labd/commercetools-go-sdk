package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountCodesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDiscountCodesRequestBuilder) WithId(id string) *ByProjectKeyDiscountCodesByIDRequestBuilder {
	return &ByProjectKeyDiscountCodesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDiscountCodesRequestBuilder) WithKey(key string) *ByProjectKeyDiscountCodesKeyByKeyRequestBuilder {
	return &ByProjectKeyDiscountCodesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Deprecated OAuth 2.0 Scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Get() *ByProjectKeyDiscountCodesRequestMethodGet {
	return &ByProjectKeyDiscountCodesRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountCode exists for a given Query Predicate. Returns a `200 OK` status if any DiscountCodes match the Query Predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Head() *ByProjectKeyDiscountCodesRequestMethodHead {
	return &ByProjectKeyDiscountCodesRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a Discount Code produces the [DiscountCodeCreated](ctp:api:type:DiscountCodeCreatedMessage) Message.
*
*	Deprecated OAuth 2.0 Scope: `manage_orders:{projectKey}`
*
 */
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Post(body DiscountCodeDraft) *ByProjectKeyDiscountCodesRequestMethodPost {
	return &ByProjectKeyDiscountCodesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}
