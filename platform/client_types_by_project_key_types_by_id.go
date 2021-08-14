// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTypesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Type by ID
 */
func (rb *ByProjectKeyTypesByIdRequestBuilder) Get() *ByProjectKeyTypesByIdRequestMethodGet {
	return &ByProjectKeyTypesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Type by ID
 */
func (rb *ByProjectKeyTypesByIdRequestBuilder) Post(body TypeUpdate) *ByProjectKeyTypesByIdRequestMethodPost {
	return &ByProjectKeyTypesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Type by ID
 */
func (rb *ByProjectKeyTypesByIdRequestBuilder) Delete() *ByProjectKeyTypesByIdRequestMethodDelete {
	return &ByProjectKeyTypesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
