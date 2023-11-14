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

/**
*	Checks if a Payment exists for a given `id`. Returns a `200 OK` status if the Payment exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyPaymentsByIDRequestBuilder) Head() *ByProjectKeyPaymentsByIDRequestMethodHead {
	return &ByProjectKeyPaymentsByIDRequestMethodHead{
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
