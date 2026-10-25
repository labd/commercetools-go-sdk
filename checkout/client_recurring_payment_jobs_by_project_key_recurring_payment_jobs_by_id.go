package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentJobsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Recurring Payment Job with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestBuilder) Get() *ByProjectKeyRecurringPaymentJobsByIdRequestMethodGet {
	return &ByProjectKeyRecurringPaymentJobsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payment-jobs/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Recurring Payment Job with a given `id`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsByIdRequestBuilder) Delete() *ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete {
	return &ByProjectKeyRecurringPaymentJobsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurring-payment-jobs/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
