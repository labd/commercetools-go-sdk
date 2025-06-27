package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRecurringOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRecurringOrdersRequestBuilder) WithId(id string) *ByProjectKeyRecurringOrdersByIDRequestBuilder {
	return &ByProjectKeyRecurringOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRecurringOrdersRequestBuilder) WithKey(key string) *ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder {
	return &ByProjectKeyRecurringOrdersKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Recurring Orders in the Project.
*
 */
func (rb *ByProjectKeyRecurringOrdersRequestBuilder) Get() *ByProjectKeyRecurringOrdersRequestMethodGet {
	return &ByProjectKeyRecurringOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/recurring-orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Recurring Orders exist for the provided query predicate. Returns a `200 OK` status if any Recurring Orders match the query predicate, or a [NotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyRecurringOrdersRequestBuilder) Head() *ByProjectKeyRecurringOrdersRequestMethodHead {
	return &ByProjectKeyRecurringOrdersRequestMethodHead{
		url:    fmt.Sprintf("/%s/recurring-orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a Recurring Order in the Project.
*	The Cart is validated to ensure that it is convertible to an [Order](ctp:api:type:Order). If validation fails, an error is returned.
*
*	Produces the [RecurringOrderCreated](ctp:api:type:RecurringOrderCreatedMessage) message.
*
*	If a server-side problem occurs, indicated by a 500 Internal Server Error HTTP response, the Recurring Order creation may still successfully complete after the error is returned.
*	If you receive this error, you should verify the status of the Recurring Order by querying a unique identifier supplied during the creation request, such as the key.
*
 */
func (rb *ByProjectKeyRecurringOrdersRequestBuilder) Post(body RecurringOrderDraft) *ByProjectKeyRecurringOrdersRequestMethodPost {
	return &ByProjectKeyRecurringOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/recurring-orders", rb.projectKey),
		client: rb.client,
	}
}
