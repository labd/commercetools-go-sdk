package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyApiClientsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyApiClientsByIDRequestBuilder) Get() *ByProjectKeyApiClientsByIDRequestMethodGet {
	return &ByProjectKeyApiClientsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/api-clients/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if an API Client exists with the provided `id`. Returns a `200 OK` status if the API Client exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyApiClientsByIDRequestBuilder) Head() *ByProjectKeyApiClientsByIDRequestMethodHead {
	return &ByProjectKeyApiClientsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/api-clients/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyApiClientsByIDRequestBuilder) Delete() *ByProjectKeyApiClientsByIDRequestMethodDelete {
	return &ByProjectKeyApiClientsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/api-clients/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
