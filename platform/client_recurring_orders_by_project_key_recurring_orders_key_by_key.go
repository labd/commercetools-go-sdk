package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Recurring Order with the provided `key`.
*
 */
func (rb *ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder) Get() *ByProjectKeyRecurringOrdersKeyByKeyRequestMethodGet {
	return &ByProjectKeyRecurringOrdersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-orders/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Recurring Order exists with the provided `key`. Returns a `200 OK` status if the Recurring Order exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder) Head() *ByProjectKeyRecurringOrdersKeyByKeyRequestMethodHead {
	return &ByProjectKeyRecurringOrdersKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurring-orders/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Recurring Order using one or more [update actions](/../api/projects/recurring-orders#update-actions).
*
 */
func (rb *ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder) Post(body RecurringOrderUpdate) *ByProjectKeyRecurringOrdersKeyByKeyRequestMethodPost {
	return &ByProjectKeyRecurringOrdersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-orders/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
