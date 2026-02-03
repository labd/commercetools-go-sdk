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
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Get() *ByProjectKeyRecurrencePoliciesByIDRequestMethodGet {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Recurrence Policy exists with the provided `id`. Returns a `200` status if the Recurrence Policy exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
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
*	The `manage_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `manage_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Post(body RecurrencePolicyUpdate) *ByProjectKeyRecurrencePoliciesByIDRequestMethodPost {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a Recurrence Policy in the Project.
*
*	A Recurrence Policy can be deleted only if it is not referenced by any Embedded Price, Standalone Price, or (Custom) Line Item, otherwise a [ReferenceExists](ctp:api:type:ReferenceExistsError) error is returned.
*
*	The `manage_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `manage_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesByIDRequestBuilder) Delete() *ByProjectKeyRecurrencePoliciesByIDRequestMethodDelete {
	return &ByProjectKeyRecurrencePoliciesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurrence-policies/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
