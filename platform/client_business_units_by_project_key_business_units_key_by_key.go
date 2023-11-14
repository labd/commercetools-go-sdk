package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder) Get() *ByProjectKeyBusinessUnitsKeyByKeyRequestMethodGet {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists for a given `key`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder) Head() *ByProjectKeyBusinessUnitsKeyByKeyRequestMethodHead {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder) Delete() *ByProjectKeyBusinessUnitsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
