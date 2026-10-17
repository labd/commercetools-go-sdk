package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyPaymentIntegrationsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyPaymentIntegrationsRequestBuilder) WithId(id string) *ByProjectKeyPaymentIntegrationsByIdRequestBuilder {
	return &ByProjectKeyPaymentIntegrationsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyPaymentIntegrationsRequestBuilder) WithKey(key string) *ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder {
	return &ByProjectKeyPaymentIntegrationsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Payment Integrations in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyPaymentIntegrationsRequestBuilder) Get() *ByProjectKeyPaymentIntegrationsRequestMethodGet {
	return &ByProjectKeyPaymentIntegrationsRequestMethodGet{
		url:    fmt.Sprintf("/%s/payment-integrations", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new PaymentIntegration.
*
 */
func (rb *ByProjectKeyPaymentIntegrationsRequestBuilder) Post(body PaymentIntegrationDraft) *ByProjectKeyPaymentIntegrationsRequestMethodPost {
	return &ByProjectKeyPaymentIntegrationsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/payment-integrations", rb.projectKey),
		client: rb.client,
	}
}
