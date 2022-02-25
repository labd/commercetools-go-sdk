package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductDiscountsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductDiscountsByIDRequestBuilder) Get() *ByProjectKeyProductDiscountsByIDRequestMethodGet {
	return &ByProjectKeyProductDiscountsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductDiscountsByIDRequestBuilder) Post(body ProductDiscountUpdate) *ByProjectKeyProductDiscountsByIDRequestMethodPost {
	return &ByProjectKeyProductDiscountsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductDiscountsByIDRequestBuilder) Delete() *ByProjectKeyProductDiscountsByIDRequestMethodDelete {
	return &ByProjectKeyProductDiscountsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-discounts/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
