package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePaymentsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Get() *ByProjectKeyMePaymentsByIDRequestMethodGet {
	return &ByProjectKeyMePaymentsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Post(body MyPaymentUpdate) *ByProjectKeyMePaymentsByIDRequestMethodPost {
	return &ByProjectKeyMePaymentsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMePaymentsByIDRequestBuilder) Delete() *ByProjectKeyMePaymentsByIDRequestMethodDelete {
	return &ByProjectKeyMePaymentsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
