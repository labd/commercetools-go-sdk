package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists for a given `key`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Head() *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Post(body MyBusinessUnitUpdate) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Delete() *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
