package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyZonesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Zone with the provided `key`.
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Get() *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	return &ByProjectKeyZonesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Zone exists with the provided `key`. Returns a `200 OK` status if the Zone exists or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Head() *ByProjectKeyZonesKeyByKeyRequestMethodHead {
	return &ByProjectKeyZonesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Zone in the Project using one or more [update actions](/../api/projects/zones#update-actions).
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Post(body ZoneUpdate) *ByProjectKeyZonesKeyByKeyRequestMethodPost {
	return &ByProjectKeyZonesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a Zone in the Project.
 */
func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Delete() *ByProjectKeyZonesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyZonesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
