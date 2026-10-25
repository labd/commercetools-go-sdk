package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentIntegrationsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves Payment Integrations in a [Project](ctp:api:type:Project) with a given `id`.
*
 */
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestBuilder) Get() *ByProjectKeyPaymentIntegrationsByIdRequestMethodGet {
	return &ByProjectKeyPaymentIntegrationsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-integrations/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a PaymentIntegration with a given `id` exists. Returns `200 OK` if the PaymentIntegration exists, or `404 Not Found` if it doesn't. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyPaymentIntegrationsByIdRequestBuilder) Head() *ByProjectKeyPaymentIntegrationsByIdRequestMethodHead {
	return &ByProjectKeyPaymentIntegrationsByIdRequestMethodHead{
		url:    fmt.Sprintf("/%s/payment-integrations/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestBuilder) Post(body PaymentIntegrationUpdateActions) *ByProjectKeyPaymentIntegrationsByIdRequestMethodPost {
	return &ByProjectKeyPaymentIntegrationsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-integrations/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyPaymentIntegrationsByIdRequestBuilder) Delete() *ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete {
	return &ByProjectKeyPaymentIntegrationsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/payment-integrations/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
