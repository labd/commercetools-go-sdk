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

func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Get() *ByProjectKeyZonesKeyByKeyRequestMethodGet {
	return &ByProjectKeyZonesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Post(body ZoneUpdate) *ByProjectKeyZonesKeyByKeyRequestMethodPost {
	return &ByProjectKeyZonesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyZonesKeyByKeyRequestBuilder) Delete() *ByProjectKeyZonesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyZonesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/zones/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
