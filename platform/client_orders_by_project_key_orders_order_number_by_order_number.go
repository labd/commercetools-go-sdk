package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder struct {
	projectKey  string
	orderNumber string
	client      *Client
}

/**
*	Retrieves an Order with the provided `orderNumber`.
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists with the provided `orderNumber`. Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Head() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodHead {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Updates an Order in the Project using one or more [update actions](/../api/projects/orders#update-actions).
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Deletes an Order in the Project.
*	Deleting an Order produces the [OrderDeleted](ctp:api:type:OrderDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Delete() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}
