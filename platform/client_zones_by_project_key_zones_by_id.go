// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyZonesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyZonesByIDRequestBuilder) Get() *ByProjectKeyZonesByIDRequestMethodGet {
	return &ByProjectKeyZonesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyZonesByIDRequestBuilder) Post(body ZoneUpdate) *ByProjectKeyZonesByIDRequestMethodPost {
	return &ByProjectKeyZonesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyZonesByIDRequestBuilder) Delete() *ByProjectKeyZonesByIDRequestMethodDelete {
	return &ByProjectKeyZonesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/zones/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
