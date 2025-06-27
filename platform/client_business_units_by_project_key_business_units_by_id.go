package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsByIDRequestBuilder) Get() *ByProjectKeyBusinessUnitsByIDRequestMethodGet {
	return &ByProjectKeyBusinessUnitsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists with the provided `id`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyBusinessUnitsByIDRequestBuilder) Head() *ByProjectKeyBusinessUnitsByIDRequestMethodHead {
	return &ByProjectKeyBusinessUnitsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsByIDRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyBusinessUnitsByIDRequestMethodPost {
	return &ByProjectKeyBusinessUnitsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsByIDRequestBuilder) Delete() *ByProjectKeyBusinessUnitsByIDRequestMethodDelete {
	return &ByProjectKeyBusinessUnitsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
