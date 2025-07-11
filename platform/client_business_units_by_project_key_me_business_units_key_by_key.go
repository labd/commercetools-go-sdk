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

/**
*	Returns a Business Unit for a given `key`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists with the provided `key`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Head() *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodHead {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a Business Unit for a given `key`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder) Post(body MyBusinessUnitUpdate) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
