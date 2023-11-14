package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks ifa My Cart exists for a given `id`. Returns a `200 OK` status if the My Cart exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	If the Cart exists in the Project but does not have the `store` field, or the `store` field references a different Store, this method returns a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Post(body MyCartUpdate) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/me/carts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
