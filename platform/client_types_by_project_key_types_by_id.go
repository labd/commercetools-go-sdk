package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTypesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Get() *ByProjectKeyTypesByIDRequestMethodGet {
	return &ByProjectKeyTypesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Type exists for a given `id`. Returns a `200 OK` status if the Type exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyTypesByIDRequestBuilder) Head() *ByProjectKeyTypesByIDRequestMethodHead {
	return &ByProjectKeyTypesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Post(body TypeUpdate) *ByProjectKeyTypesByIDRequestMethodPost {
	return &ByProjectKeyTypesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesByIDRequestBuilder) Delete() *ByProjectKeyTypesByIDRequestMethodDelete {
	return &ByProjectKeyTypesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/types/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
