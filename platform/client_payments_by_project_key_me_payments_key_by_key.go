// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMePaymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get MyPayment by key
 */
func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyMePaymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update MyPayment by key
 */
func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Post(body MyPaymentUpdate) *ByProjectKeyMePaymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete MyPayment by key
 */
func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMePaymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
