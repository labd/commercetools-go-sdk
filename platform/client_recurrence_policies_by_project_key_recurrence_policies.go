package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurrencePoliciesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecurrencePoliciesRequestBuilder) WithKey(key string) *ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder {
	return &ByProjectKeyRecurrencePoliciesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRecurrencePoliciesRequestBuilder) WithId(id string) *ByProjectKeyRecurrencePoliciesByIDRequestBuilder {
	return &ByProjectKeyRecurrencePoliciesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Recurrence Policies in the Project.
*
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesRequestBuilder) Get() *ByProjectKeyRecurrencePoliciesRequestMethodGet {
	return &ByProjectKeyRecurrencePoliciesRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurrence-policies", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Recurrence Policies exist for the provided query predicate. Returns a `200` status if any Recurrence Policies match the query predicate, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
*	The `view_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `view_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesRequestBuilder) Head() *ByProjectKeyRecurrencePoliciesRequestMethodHead {
	return &ByProjectKeyRecurrencePoliciesRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurrence-policies", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a Recurrence Policy in the Project.
*
*	The `manage_recurring_orders:{projectKey}` scope is deprecated for use on this endpoint. Update your clients to use the `manage_recurrence_policies:{projectKey}` scope instead. For more information, see the [Deprecations and removals](/api/deprecations-and-removals#recurrence-policies) list.
*
 */
func (rb *ByProjectKeyRecurrencePoliciesRequestBuilder) Post(body RecurrencePolicyDraft) *ByProjectKeyRecurrencePoliciesRequestMethodPost {
	return &ByProjectKeyRecurrencePoliciesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurrence-policies", rb.projectKey),
		client: rb.client,
	}
}
