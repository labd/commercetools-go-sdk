package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Get() *ByProjectKeyCustomersKeyByKeyRequestMethodGet {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Post(body CustomerUpdate) *ByProjectKeyCustomersKeyByKeyRequestMethodPost {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deleting a Customer produces the [CustomerDeleted](ctp:api:type:CustomerDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyCustomersKeyByKeyRequestBuilder) Delete() *ByProjectKeyCustomersKeyByKeyRequestMethodDelete {
	return &ByProjectKeyCustomersKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/customers/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
