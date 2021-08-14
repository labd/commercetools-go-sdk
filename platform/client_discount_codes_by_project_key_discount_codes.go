// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyDiscountCodesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDiscountCodesRequestBuilder) WithId(id string) *ByProjectKeyDiscountCodesByIdRequestBuilder {
	return &ByProjectKeyDiscountCodesByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query discount-codes
 */
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Get() *ByProjectKeyDiscountCodesRequestMethodGet {
	return &ByProjectKeyDiscountCodesRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create DiscountCode
 */
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Post(body DiscountCodeDraft) *ByProjectKeyDiscountCodesRequestMethodPost {
	return &ByProjectKeyDiscountCodesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}
