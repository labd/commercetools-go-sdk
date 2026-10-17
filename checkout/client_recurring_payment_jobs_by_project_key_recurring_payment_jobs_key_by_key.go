package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentJobsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Recurring Payment Job with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsKeyByKeyRequestBuilder) Get() *ByProjectKeyRecurringPaymentJobsKeyByKeyRequestMethodGet {
	return &ByProjectKeyRecurringPaymentJobsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payment-jobs/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Recurring Payment Job with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsKeyByKeyRequestBuilder) Delete() *ByProjectKeyRecurringPaymentJobsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyRecurringPaymentJobsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurring-payment-jobs/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
