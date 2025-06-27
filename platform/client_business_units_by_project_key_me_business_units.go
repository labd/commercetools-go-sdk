package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeBusinessUnitsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyMeBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyMeBusinessUnitsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyMeBusinessUnitsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves Business Units for the authenticated Customer. Returns a `200 OK` status if successful, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) Get() *ByProjectKeyMeBusinessUnitsRequestMethodGet {
	return &ByProjectKeyMeBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/business-units", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more BusinessUnits exist for the provided query predicate. Returns a `200 OK` status if any BusinessUnits match the Query Predicate and the Customer has access to them, or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
*
 */
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) Head() *ByProjectKeyMeBusinessUnitsRequestMethodHead {
	return &ByProjectKeyMeBusinessUnitsRequestMethodHead{
		url:    fmt.Sprintf("/%s/me/business-units", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Automatically assigns the Associate to the Business Unit in the default [Associate Role](ctp:api:type:AssociateRole) defined in [BusinessUnitConfiguration](ctp:api:type:BusinessUnitConfiguration). If there is no default Associate Role configured, this request fails with an [InvalidOperation](ctp:api:type:InvalidOperationError) error. When creating a Division, the Associate must have the `AddChildUnits` [Permission](ctp:api:type:Permission) in the parent unit. If the required [Permission](/projects/associate-roles#permission) is missing, an [AssociateMissingPermission](/errors#associatemissingpermission) error is returned.
*
 */
func (rb *ByProjectKeyMeBusinessUnitsRequestBuilder) Post(body MyBusinessUnitDraft) *ByProjectKeyMeBusinessUnitsRequestMethodPost {
	return &ByProjectKeyMeBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/business-units", rb.projectKey),
		client: rb.client,
	}
}
