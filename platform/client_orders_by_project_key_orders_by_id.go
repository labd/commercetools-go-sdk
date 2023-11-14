package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersByIDRequestBuilder) Get() *ByProjectKeyOrdersByIDRequestMethodGet {
	return &ByProjectKeyOrdersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an Order exists for a given `id`. Returns a `200 OK` status if the Order exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersByIDRequestBuilder) Head() *ByProjectKeyOrdersByIDRequestMethodHead {
	return &ByProjectKeyOrdersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyOrdersByIDRequestBuilder) Post(body OrderUpdate) *ByProjectKeyOrdersByIDRequestMethodPost {
	return &ByProjectKeyOrdersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deleting an Order produces the [OrderDeleted](ctp:api:type:OrderDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyOrdersByIDRequestBuilder) Delete() *ByProjectKeyOrdersByIDRequestMethodDelete {
	return &ByProjectKeyOrdersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
