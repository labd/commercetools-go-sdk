package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists with the provided `key`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
