package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyZonesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyZonesRequestBuilder) WithKey(key string) *ByProjectKeyZonesKeyByKeyRequestBuilder {
	return &ByProjectKeyZonesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyZonesRequestBuilder) WithId(id string) *ByProjectKeyZonesByIDRequestBuilder {
	return &ByProjectKeyZonesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves all Zones in the Project.
 */
func (rb *ByProjectKeyZonesRequestBuilder) Get() *ByProjectKeyZonesRequestMethodGet {
	return &ByProjectKeyZonesRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Zones exist for the provided query predicate. Returns a `200 OK` status if any Zones match the query predicate, or a [Not Found](/../api/errors#404-not-found) error otherwise.
 */
func (rb *ByProjectKeyZonesRequestBuilder) Head() *ByProjectKeyZonesRequestMethodHead {
	return &ByProjectKeyZonesRequestMethodHead{
		url:    fmt.Sprintf("/%s/zones", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a Zone in the Project.
 */
func (rb *ByProjectKeyZonesRequestBuilder) Post(body ZoneDraft) *ByProjectKeyZonesRequestMethodPost {
	return &ByProjectKeyZonesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones", rb.projectKey),
		client: rb.client,
	}
}
