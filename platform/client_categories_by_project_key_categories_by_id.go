package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Either the [scope](/api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Get() *ByProjectKeyCategoriesByIDRequestMethodGet {
	return &ByProjectKeyCategoriesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Category exists with the provided `id`. Returns a `200` status if the Category exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Head() *ByProjectKeyCategoriesByIDRequestMethodHead {
	return &ByProjectKeyCategoriesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Either the [scope](/api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
*	The response status code is `200 OK`, unless the request contains a [Set Stores with Descendants](ctp:api:type:CategorySetStoresWithDescendantsAction), [Add Store with Descendants](ctp:api:type:CategoryAddStoreWithDescendantsAction), or [Remove Store with Descendants](ctp:api:type:CategoryRemoveStoreWithDescendantsAction) update action. In that case, the API responds with `202 Accepted`, and the returned Category contains a `warnings` field with a [StoreSettingOngoing](ctp:api:type:StoreSettingOngoingWarning) Warning.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesByIDRequestMethodPost {
	return &ByProjectKeyCategoriesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Either the [scope](/api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Delete() *ByProjectKeyCategoriesByIDRequestMethodDelete {
	return &ByProjectKeyCategoriesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
