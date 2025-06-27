package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountGroupsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a DiscountGroup with the provided `id`.
*
 */
func (rb *ByProjectKeyDiscountGroupsByIDRequestBuilder) Get() *ByProjectKeyDiscountGroupsByIDRequestMethodGet {
	return &ByProjectKeyDiscountGroupsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountGroup exists with the provided `id`.
*	Returns a `200 OK` status if a DiscountGroup exists; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyDiscountGroupsByIDRequestBuilder) Head() *ByProjectKeyDiscountGroupsByIDRequestMethodHead {
	return &ByProjectKeyDiscountGroupsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a DiscountGroup in the Project using one or more [update actions](/../api/projects/discount-groups#update-actions).
*
 */
func (rb *ByProjectKeyDiscountGroupsByIDRequestBuilder) Post(body DiscountGroupUpdate) *ByProjectKeyDiscountGroupsByIDRequestMethodPost {
	return &ByProjectKeyDiscountGroupsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-groups/%s", rb.projectKey, rb.id),
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
func (rb *ByProjectKeyDiscountGroupsByIDRequestBuilder) Delete() *ByProjectKeyDiscountGroupsByIDRequestMethodDelete {
	return &ByProjectKeyDiscountGroupsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-groups/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
