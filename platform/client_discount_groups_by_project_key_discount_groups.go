package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountGroupsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyDiscountGroupsRequestBuilder) WithKey(key string) *ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder {
	return &ByProjectKeyDiscountGroupsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyDiscountGroupsRequestBuilder) WithId(id string) *ByProjectKeyDiscountGroupsByIDRequestBuilder {
	return &ByProjectKeyDiscountGroupsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves all DiscountGroups in the Project.
*
 */
func (rb *ByProjectKeyDiscountGroupsRequestBuilder) Get() *ByProjectKeyDiscountGroupsRequestMethodGet {
	return &ByProjectKeyDiscountGroupsRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-groups", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more DiscountGroups exist for the provided query predicate.
*	Returns a `200 OK` status if any DiscountGroups match the query predicate; otherwise, returns a [Not Found](/../api/errors#404-not-found).
*
 */
func (rb *ByProjectKeyDiscountGroupsRequestBuilder) Head() *ByProjectKeyDiscountGroupsRequestMethodHead {
	return &ByProjectKeyDiscountGroupsRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-groups", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a DiscountGroup in the Project.
*	This request generates the [DiscountGroupCreated](ctp:api:type:DiscountGroupCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyDiscountGroupsRequestBuilder) Post(body DiscountGroupDraft) *ByProjectKeyDiscountGroupsRequestMethodPost {
	return &ByProjectKeyDiscountGroupsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-groups", rb.projectKey),
		client: rb.client,
	}
}
