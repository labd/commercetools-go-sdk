package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeBusinessUnitsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Returns a Business Unit for a given `id`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsByIDRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists with the provided `id`. Returns a `200 OK` status if the BusinessUnit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Head() *ByProjectKeyMeBusinessUnitsByIDRequestMethodHead {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a Business Unit for a given `id`. Returns a `200 OK` status if the Business Unit exists and the Customer has access to it, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeBusinessUnitsByIDRequestBuilder) Post(body MyBusinessUnitUpdate) *ByProjectKeyMeBusinessUnitsByIDRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
