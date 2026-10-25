package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves Payment Integrations in a [Project](ctp:api:type:Project) with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder) Get() *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodGet {
	return &ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-integrations/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder) Post(body PaymentIntegrationUpdateActions) *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodPost {
	return &ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-integrations/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a PaymentIntegration with a given `key` exists. Returns `200 OK` if the PaymentIntegration exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder) Head() *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodHead {
	return &ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/payment-integrations/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder) Delete() *ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyPaymentIntegrationsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payment-integrations/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
