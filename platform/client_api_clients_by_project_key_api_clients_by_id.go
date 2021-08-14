// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyApiClientsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get ApiClient by ID
 */
func (rb *ByProjectKeyApiClientsByIdRequestBuilder) Get() *ByProjectKeyApiClientsByIdRequestMethodGet {
	return &ByProjectKeyApiClientsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/api-clients/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete ApiClient by ID
 */
func (rb *ByProjectKeyApiClientsByIdRequestBuilder) Delete() *ByProjectKeyApiClientsByIdRequestMethodDelete {
	return &ByProjectKeyApiClientsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/api-clients/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
