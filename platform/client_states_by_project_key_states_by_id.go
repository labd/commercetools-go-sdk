// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyStatesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get State by ID
 */
func (rb *ByProjectKeyStatesByIdRequestBuilder) Get() *ByProjectKeyStatesByIdRequestMethodGet {
	return &ByProjectKeyStatesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update State by ID
 */
func (rb *ByProjectKeyStatesByIdRequestBuilder) Post(body StateUpdate) *ByProjectKeyStatesByIdRequestMethodPost {
	return &ByProjectKeyStatesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete State by ID
 */
func (rb *ByProjectKeyStatesByIdRequestBuilder) Delete() *ByProjectKeyStatesByIdRequestMethodDelete {
	return &ByProjectKeyStatesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/states/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
