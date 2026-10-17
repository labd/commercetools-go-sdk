package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

/**
*	Retrieves a [Category](ctp:api:type:Category) by its `id` if it is assigned to the specified [Store](ctp:api:type:Store) or is global.
*
*	For global Categories, use the [Get Category by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:GET) endpoint.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Category exists with the provided `id` in the specified [Store](ctp:api:type:Store). Returns a `200` status if the Category exists in the Store or is global, or a `404` status otherwise.
*
*	For global Categories, use the [Check if Category exists by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:HEAD) endpoint.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Updates a [Category](ctp:api:type:Category) by its `id` in the specified [Store](ctp:api:type:Store).
*
*	To update a global Category, use the [Update Category by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:POST) endpoint.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Deletes a [Category](ctp:api:type:Category) by its `id` in the specified [Store](ctp:api:type:Store).
*
*	To delete a global Category, use the [Delete Category by ID](ctp:api:endpoint:/{projectKey}/categories/{id}:DELETE) endpoint.
*
*	If you do not have permissions for a Store the Category is assigned to, an [Unauthorized](ctp:api:type:UnauthorizedError) error is returned.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
