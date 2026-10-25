package integrationlayer

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMainCheckoutRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMainCheckoutRequestBuilder) Put(body CheckoutConfig) *ByProjectKeyMainCheckoutRequestMethodPut {
	return &ByProjectKeyMainCheckoutRequestMethodPut{
		body:   body,
		url:    fmt.Sprintf("/%s/main/checkout", rb.projectKey),
		client: rb.client,
	}
}
