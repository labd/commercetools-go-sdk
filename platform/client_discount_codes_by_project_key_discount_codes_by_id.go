// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyDiscountCodesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Get() *ByProjectKeyDiscountCodesByIDRequestMethodGet {
	return &ByProjectKeyDiscountCodesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Post(body DiscountCodeUpdate) *ByProjectKeyDiscountCodesByIDRequestMethodPost {
	return &ByProjectKeyDiscountCodesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyDiscountCodesByIDRequestBuilder) Delete() *ByProjectKeyDiscountCodesByIDRequestMethodDelete {
	return &ByProjectKeyDiscountCodesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-codes/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
