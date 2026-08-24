package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStoresKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Get() *ByProjectKeyStoresKeyByKeyRequestMethodGet {
	return &ByProjectKeyStoresKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Store exists with the provided `key`. Returns a `200` status if the Store exists, or a `404` status otherwise.
 */
func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Head() *ByProjectKeyStoresKeyByKeyRequestMethodHead {
	return &ByProjectKeyStoresKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Post(body StoreUpdate) *ByProjectKeyStoresKeyByKeyRequestMethodPost {
	return &ByProjectKeyStoresKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	If the Store is referenced by a resource, a [ReferenceExists](ctp:api:type:ReferenceExistsError) error is returned.
*
*	To avoid dependency conflicts, first delete all active and anonymous Carts that reference the Store, then delete all Orders.
*
 */
func (rb *ByProjectKeyStoresKeyByKeyRequestBuilder) Delete() *ByProjectKeyStoresKeyByKeyRequestMethodDelete {
	return &ByProjectKeyStoresKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/stores/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
