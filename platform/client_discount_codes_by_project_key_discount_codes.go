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
func (rb *ByProjectKeyDiscountCodesRequestBuilder) Get() *ByProjectKeyDiscountCodesRequestMethodGet {
	return &ByProjectKeyDiscountCodesRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDiscountCodesRequestBuilder) Post(body DiscountCodeDraft) *ByProjectKeyDiscountCodesRequestMethodPost {
	return &ByProjectKeyDiscountCodesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes", rb.projectKey),
		client: rb.client,
	}
}
