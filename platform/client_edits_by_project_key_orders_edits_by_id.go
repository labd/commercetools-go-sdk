package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Apply() *ByProjectKeyOrdersEditsByIDApplyRequestBuilder {
	return &ByProjectKeyOrdersEditsByIDApplyRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	Retrieves an OrderEdit with the provided `id`.
 */
func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Get() *ByProjectKeyOrdersEditsByIDRequestMethodGet {
	return &ByProjectKeyOrdersEditsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an OrderEdit exists with the provided `id`. Returns a `200 OK` status if the OrderEdit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Head() *ByProjectKeyOrdersEditsByIDRequestMethodHead {
	return &ByProjectKeyOrdersEditsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates an OrderEdit in the Project using one or more [update actions](/../api/projects/order-edits#update-actions).
 */
func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Post(body OrderEditUpdate) *ByProjectKeyOrdersEditsByIDRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes an OrderEdit in the Project.
 */
func (rb *ByProjectKeyOrdersEditsByIDRequestBuilder) Delete() *ByProjectKeyOrdersEditsByIDRequestMethodDelete {
	return &ByProjectKeyOrdersEditsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
