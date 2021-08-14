// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersEditsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsByIdRequestBuilder) Apply() *ByProjectKeyOrdersEditsByIdApplyRequestBuilder {
	return &ByProjectKeyOrdersEditsByIdApplyRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	Get OrderEdit by ID
 */
func (rb *ByProjectKeyOrdersEditsByIdRequestBuilder) Get() *ByProjectKeyOrdersEditsByIdRequestMethodGet {
	return &ByProjectKeyOrdersEditsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update OrderEdit by ID
 */
func (rb *ByProjectKeyOrdersEditsByIdRequestBuilder) Post(body OrderEditUpdate) *ByProjectKeyOrdersEditsByIdRequestMethodPost {
	return &ByProjectKeyOrdersEditsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete OrderEdit by ID
 */
func (rb *ByProjectKeyOrdersEditsByIdRequestBuilder) Delete() *ByProjectKeyOrdersEditsByIdRequestMethodDelete {
	return &ByProjectKeyOrdersEditsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/edits/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
