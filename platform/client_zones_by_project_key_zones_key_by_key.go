// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyZonesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Zone by key
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Get() *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	return &ByProjectKeyZonesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Zone by key
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Post(body ZoneUpdate) *ByProjectKeyZonesKeyByKeyRequestMethodPost {
	return &ByProjectKeyZonesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Zone by key
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Delete() *ByProjectKeyZonesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyZonesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
