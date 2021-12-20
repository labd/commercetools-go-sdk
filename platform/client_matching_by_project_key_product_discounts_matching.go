// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductDiscountsMatchingRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductDiscountsMatchingRequestBuilder) Post(body ProductDiscountMatchQuery) *ByProjectKeyProductDiscountsMatchingRequestMethodPost {
	return &ByProjectKeyProductDiscountsMatchingRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-discounts/matching", rb.projectKey),
		client: rb.client,
	}
}
