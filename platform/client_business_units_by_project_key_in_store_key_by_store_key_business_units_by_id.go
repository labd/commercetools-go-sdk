package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists with the provided `id`. Returns a `200 OK` status if the BusinessUnit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder) Post(body BusinessUnitUpdate) *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyBusinessUnitsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/business-units/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
