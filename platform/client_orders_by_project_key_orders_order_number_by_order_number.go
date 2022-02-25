package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder struct {
	projectKey  string
	orderNumber string
	client      *Client
}

/**
*	In case the orderNumber does not match the regular expression [a-zA-Z0-9_\-]+,
*	it should be provided in URL-encoded format.
*
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Delete() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}
