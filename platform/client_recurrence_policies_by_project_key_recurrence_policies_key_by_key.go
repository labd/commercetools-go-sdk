package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Recurrence Policy with the provided `key`.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Get() *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Recurrence Policy exists with the provided `key`. Returns a `200 OK` status if the Recurrence Policy exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Head() *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodHead {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Recurrence Policy using one or more [update actions](/../api/projects/recurrence-policies#update-actions).
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Post(body RecurrencePolicyUpdate) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodPost {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
