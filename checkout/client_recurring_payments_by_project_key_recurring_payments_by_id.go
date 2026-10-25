package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Recurring Payment with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsByIdRequestBuilder) Get() *ByProjectKeyRecurringPaymentsByIdRequestMethodGet {
	return &ByProjectKeyRecurringPaymentsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Recurring Payment with a given `id`. Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsByIdRequestBuilder) Post(body RecurringPaymentUpdateActions) *ByProjectKeyRecurringPaymentsByIdRequestMethodPost {
	return &ByProjectKeyRecurringPaymentsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Recurring Payment with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsByIdRequestBuilder) Delete() *ByProjectKeyRecurringPaymentsByIdRequestMethodDelete {
	return &ByProjectKeyRecurringPaymentsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurring-payments/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
