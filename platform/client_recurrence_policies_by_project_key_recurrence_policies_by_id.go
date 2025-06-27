package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurrencePoliciesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Recurrence Policy with the provided `id`.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Get() *ByProjectKeyRecurrencePoliciesByIDRequestMethodGet {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Recurrence Policy exists with the provided `id`. Returns a `200 OK` status if the Recurrence Policy exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Head() *ByProjectKeyRecurrencePoliciesByIDRequestMethodHead {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Recurrence Policy using one or more [update actions](/../api/projects/recurrence-policies#update-actions).
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Post(body RecurrencePolicyUpdate) *ByProjectKeyRecurrencePoliciesByIDRequestMethodPost {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
