package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder) Products() *ByProjectKeyProductSelectionsKeyByKeyProductsRequestBuilder {
	return &ByProjectKeyProductSelectionsKeyByKeyProductsRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductSelectionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductSelectionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-selections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a ProductSelection exists for a given `key`. Returns a `200 OK` status if the ProductSelection exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductSelectionsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductSelectionsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-selections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder) Post(body ProductSelectionUpdate) *ByProjectKeyProductSelectionsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductSelectionsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-selections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletion will only succeed if the Product Selection is not assigned to any [Store](ctp:api:type:Store).
 */
func (rb *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductSelectionsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductSelectionsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-selections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
