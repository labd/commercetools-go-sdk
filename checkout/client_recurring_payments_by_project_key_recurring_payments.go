package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecurringPaymentsRequestBuilder) WithId(id string) *ByProjectKeyRecurringPaymentsByIdRequestBuilder {
	return &ByProjectKeyRecurringPaymentsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRecurringPaymentsRequestBuilder) WithKey(key string) *ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder {
	return &ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Recurring Payments in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyRecurringPaymentsRequestBuilder) Get() *ByProjectKeyRecurringPaymentsRequestMethodGet {
	return &ByProjectKeyRecurringPaymentsRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payments", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new [RecurringPayment](ctp:checkout:type:RecurringPayment). Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsRequestBuilder) Post(body RecurringPaymentDraft) *ByProjectKeyRecurringPaymentsRequestMethodPost {
	return &ByProjectKeyRecurringPaymentsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-payments", rb.projectKey),
		client: rb.client,
	}
}
