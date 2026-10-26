package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

/**
*	Retrieves a [Category](ctp:api:type:Category) by its `key` in the specified [Store](ctp:api:type:Store).
*
*	For global Categories, use the [Get Category by Key](ctp:api:endpoint:/{projectKey}/categories/key={key}:GET) endpoint.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Category exists with the provided `key` in the specified [Store](ctp:api:type:Store). Returns a `200` status if the Category exists in the Store or is global, or a `404` status otherwise.
*
*	For global Categories, use the [Check if Category exists by Key](ctp:api:endpoint:/{projectKey}/categories/key={key}:HEAD) endpoint.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates a [Category](ctp:api:type:Category) by its `key` in the specified [Store](ctp:api:type:Store).
*
*	To update a global Category, use the [Update Category by Key](ctp:api:endpoint:/{projectKey}/categories/key={key}:POST) endpoint.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes a [Category](ctp:api:type:Category) by its `key` in the specified [Store](ctp:api:type:Store).
*
*	To delete a global Category, use the [Delete Category by Key](ctp:api:endpoint:/{projectKey}/categories/key={key}:DELETE) endpoint.
*
*	If you do not have permissions for a Store the Category is assigned to, an [Unauthorized](ctp:api:type:UnauthorizedError) error is returned.
*
*	If the Category does not exist in the Store, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
