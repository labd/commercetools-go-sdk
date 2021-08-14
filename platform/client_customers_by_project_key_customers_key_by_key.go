// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCustomersKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Customer by key
 */
func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Get() *ByProjectKeyCustomersKeyByKeyRequestMethodGet {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Customer by key
 */
func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyCustomersKeyByKeyRequestMethodPost {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Customer by key
 */
func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Delete() *ByProjectKeyCustomersKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
