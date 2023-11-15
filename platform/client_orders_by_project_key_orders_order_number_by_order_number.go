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

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Get() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `orderNumber`. Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Head() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodHead {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}

/**
*	Deleting an Order produces the [OrderDeleted](ctp:api:type:OrderDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder) Delete() *ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/order-number=%s", rb.projectKey, rb.orderNumber),
		client: rb.client,
	}
}
