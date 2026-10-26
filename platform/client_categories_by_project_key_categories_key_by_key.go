package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCategoriesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Either the [scope](/api/scopes) `view_products:{projectKey}` or `view_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Get() *ByProjectKeyCategoriesKeyByKeyRequestMethodGet {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Category exists with the provided `key`. Returns a `200` status if the Category exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Head() *ByProjectKeyCategoriesKeyByKeyRequestMethodHead {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Either the [scope](/api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
*	The response status code is `200 OK`, unless the request contains a [Set Stores with Descendants](ctp:api:type:CategorySetStoresWithDescendantsAction), [Add Store with Descendants](ctp:api:type:CategoryAddStoreWithDescendantsAction), or [Remove Store with Descendants](ctp:api:type:CategoryRemoveStoreWithDescendantsAction) update action. In that case, the API responds with `202 Accepted`, and the returned Category contains a `warnings` field with a [StoreSettingOngoing](ctp:api:type:StoreSettingOngoingWarning) Warning.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesKeyByKeyRequestMethodPost {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Either the [scope](/api/scopes) `manage_products:{projectKey}` or `manage_categories:{projectKey}` is required.
*
 */
func (rb *ByProjectKeyCategoriesKeyByKeyRequestBuilder) Delete() *ByProjectKeyCategoriesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCategoriesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
