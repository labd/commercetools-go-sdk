package checkout

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringPaymentJobsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecurringPaymentJobsRequestBuilder) WithId(id string) *ByProjectKeyRecurringPaymentJobsByIdRequestBuilder {
	return &ByProjectKeyRecurringPaymentJobsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRecurringPaymentJobsRequestBuilder) WithKey(key string) *ByProjectKeyRecurringPaymentJobsKeyByKeyRequestBuilder {
	return &ByProjectKeyRecurringPaymentJobsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Recurring Payment Jobs in a [Project](ctp:api:type:Project).
*
*	The results are [paginated](/api/general-concepts#pagination).
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsRequestBuilder) Get() *ByProjectKeyRecurringPaymentJobsRequestMethodGet {
	return &ByProjectKeyRecurringPaymentJobsRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-payment-jobs", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new [RecurringPaymentJob](ctp:checkout:type:RecurringPaymentJob). Specific Error Codes:
*	- [InvalidInput](ctp:checkout:type:InvalidInputError)
*	- [ReferencedResourceNotFound](ctp:checkout:type:ReferencedResourceNotFoundError)
*	- [RequiredField](ctp:checkout:type:RequiredFieldError)
*
 */
func (rb *ByProjectKeyRecurringPaymentJobsRequestBuilder) Post(body RecurringPaymentJobDraft) *ByProjectKeyRecurringPaymentJobsRequestMethodPost {
	return &ByProjectKeyRecurringPaymentJobsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-payment-jobs", rb.projectKey),
		client: rb.client,
	}
}
