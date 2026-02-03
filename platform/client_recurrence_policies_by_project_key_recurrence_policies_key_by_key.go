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
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Get() *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Recurrence Policy exists with the provided `key`. Returns a `200` status if the Recurrence Policy exists, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
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
*	The `manage_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `manage_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Post(body RecurrencePolicyUpdate) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodPost {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
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
func (rb *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder) Delete() *ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/recurrence-policies/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
