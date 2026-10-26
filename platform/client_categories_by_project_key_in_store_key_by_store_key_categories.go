package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}

/**
*	Retrieves [Categories](ctp:api:type:Category) that are either assigned to the specified [Store](ctp:api:type:Store) or global.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Categories exist in the [Store](ctp:api:type:Store) for the provided query predicate. Returns a `200` status if any Categories match the query predicate, or a `404` status otherwise.
*
*	For global Categories, use the [Check if Category exists by Query Predicate](ctp:api:endpoint:/{projectKey}/categories:HEAD) endpoint.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a [Category](ctp:api:type:Category) in the specified [Store](ctp:api:type:Store).
*
*	For global Categories, use the [Create Category](ctp:api:endpoint:/{projectKey}/categories:POST) endpoint.
*
*	Creating a Category with a `parent` locks that parent Category. For details, see [Category tree locking](/api/projects/categories#category-tree-locking).
*
*	Creating a Category produces the [CategoryCreated](ctp:api:type:CategoryCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestBuilder) Post(body CategoryDraft) *ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCategoriesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/categories", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
