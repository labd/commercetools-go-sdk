package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCustomersByIDRequestBuilder) Get() *ByProjectKeyCustomersByIDRequestMethodGet {
	return &ByProjectKeyCustomersByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Customer exists for a given `id`. Returns a `200 OK` status if the Customer exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCustomersByIDRequestBuilder) Head() *ByProjectKeyCustomersByIDRequestMethodHead {
	return &ByProjectKeyCustomersByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Simultaneously updating two Customers with the same email address can return a [LockedField](ctp:api:type:LockedFieldError) error.
 */
func (rb *ByProjectKeyCustomersByIDRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyCustomersByIDRequestMethodPost {
	return &ByProjectKeyCustomersByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deleting a Customer produces the [CustomerDeleted](ctp:api:type:CustomerDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyCustomersByIDRequestBuilder) Delete() *ByProjectKeyCustomersByIDRequestMethodDelete {
	return &ByProjectKeyCustomersByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customers/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
