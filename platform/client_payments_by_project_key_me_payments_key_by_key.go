package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMePaymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyMePaymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	This endpoint can only update a Payment when it has no [Transactions](ctp:api:type:Transaction).
*
 */
func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Post(body MyPaymentUpdate) *ByProjectKeyMePaymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	This endpoint can only delete a Payment when it has no [Transactions](ctp:api:type:Transaction).
*
 */
func (rb *ByProjectKeyMePaymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMePaymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMePaymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
