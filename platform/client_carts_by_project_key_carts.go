package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCartsRequestBuilder) Replicate() *ByProjectKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithCustomerId(customerId string) *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder{
		customerId: customerId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyCartsByIDRequestBuilder {
	return &ByProjectKeyCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) Get() *ByProjectKeyCartsRequestMethodGet {
	return &ByProjectKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a cart can fail with an InvalidOperation if the referenced shipping method in the
*	CartDraft has a predicate which does not match the cart.
*
 */
func (rb *ByProjectKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyCartsRequestMethodPost {
	return &ByProjectKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}
