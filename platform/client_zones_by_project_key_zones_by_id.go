// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyZonesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Zone by ID
 */
func (rb *ByProjectKeyZonesByIdRequestBuilder) Get() *ByProjectKeyZonesByIdRequestMethodGet {
	return &ByProjectKeyZonesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Zone by ID
 */
func (rb *ByProjectKeyZonesByIdRequestBuilder) Post(body ZoneUpdate) *ByProjectKeyZonesByIdRequestMethodPost {
	return &ByProjectKeyZonesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Zone by ID
 */
func (rb *ByProjectKeyZonesByIdRequestBuilder) Delete() *ByProjectKeyZonesByIdRequestMethodDelete {
	return &ByProjectKeyZonesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
