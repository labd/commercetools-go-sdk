package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a DiscountGroup with the provided `key`.
*
 */
func (rb *ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder) Get() *ByProjectKeyDiscountGroupsKeyByKeyRequestMethodGet {
	return &ByProjectKeyDiscountGroupsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountGroup exists with the provided `key`.
*	Returns a `200 OK` status if a DiscountGroup exists; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder) Head() *ByProjectKeyDiscountGroupsKeyByKeyRequestMethodHead {
	return &ByProjectKeyDiscountGroupsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a DiscountGroup in the Project using one or more [update actions](/../api/projects/discount-groups#update-actions).
*
 */
func (rb *ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder) Post(body DiscountGroupUpdate) *ByProjectKeyDiscountGroupsKeyByKeyRequestMethodPost {
	return &ByProjectKeyDiscountGroupsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a DiscountGroup in the Project.
*	This request generates the [DiscountGroupDeleted](ctp:api:type:DiscountGroupDeletedMessage) Message.
*
*	If the DiscountGroup is referenced by a CartDiscount, a [ReferenceExists](ctp:api:type:ReferenceExistsError) error is returned.
*
 */
func (rb *ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder) Delete() *ByProjectKeyDiscountGroupsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyDiscountGroupsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-groups/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
