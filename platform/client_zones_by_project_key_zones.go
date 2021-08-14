// Generated file, please do not change!!!
package platform

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
func (rb *ByProjectKeyZonesRequestBuilder) Get() *ByProjectKeyZonesRequestMethodGet {
	return &ByProjectKeyZonesRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyZonesRequestBuilder) Post(body ZoneDraft) *ByProjectKeyZonesRequestMethodPost {
	return &ByProjectKeyZonesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones", rb.projectKey),
		client: rb.client,
	}
}
