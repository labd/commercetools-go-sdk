package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyPaymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Post(body PaymentUpdate) *ByProjectKeyPaymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyPaymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyPaymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
