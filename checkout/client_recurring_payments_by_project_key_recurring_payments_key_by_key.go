package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Recurring Payment with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [GeneralError](ctp:checkout:type:GeneralError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder) Get() *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodGet {
	return &ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Recurring Payment with a given `key`. Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder) Post(body RecurringPaymentUpdateActions) *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost {
	return &ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Recurring Payment with a given `key`. Specific Error Codes:
*	- [ResourceNotFound](ctp:checkout:type:ResourceNotFoundError)
*	- [ConcurrentModification](ctp:checkout:type:ConcurrentModificationError)
*
 */
func (rb *ByProjectKeyRecurringPaymentsKeyByKeyRequestBuilder) Delete() *ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyRecurringPaymentsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurring-payments/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
