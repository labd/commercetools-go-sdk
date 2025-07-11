package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Recurring Order with the provided `id`.
*
 */
func (rb *ByProjectKeyRecurringOrdersByIDRequestBuilder) Get() *ByProjectKeyRecurringOrdersByIDRequestMethodGet {
	return &ByProjectKeyRecurringOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Recurring Order exists with the provided `id`. Returns a `200 OK` status if the Recurring Order exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyRecurringOrdersByIDRequestBuilder) Head() *ByProjectKeyRecurringOrdersByIDRequestMethodHead {
	return &ByProjectKeyRecurringOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurring-orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Recurring Order using one or more [update actions](/../api/projects/recurring-orders#update-actions).
*
 */
func (rb *ByProjectKeyRecurringOrdersByIDRequestBuilder) Post(body RecurringOrderUpdate) *ByProjectKeyRecurringOrdersByIDRequestMethodPost {
	return &ByProjectKeyRecurringOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
