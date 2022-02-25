package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Get() *ByProjectKeyPaymentsByIDRequestMethodGet {
	return &ByProjectKeyPaymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Post(body PaymentUpdate) *ByProjectKeyPaymentsByIDRequestMethodPost {
	return &ByProjectKeyPaymentsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Delete() *ByProjectKeyPaymentsByIDRequestMethodDelete {
	return &ByProjectKeyPaymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
