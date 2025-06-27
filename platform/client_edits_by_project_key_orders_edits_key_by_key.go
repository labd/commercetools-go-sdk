package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyOrdersEditsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves an OrderEdit with the provided `key`.
 */
func (rb *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder) Get() *ByProjectKeyOrdersEditsKeyByKeyRequestMethodGet {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if an OrderEdit exists with the provided `key`. Returns a `200 OK` status if the OrderEdit exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder) Head() *ByProjectKeyOrdersEditsKeyByKeyRequestMethodHead {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/orders/edits/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Updates an OrderEdit in the Project using one or more [update actions](/../api/projects/order-edits#update-actions).
 */
func (rb *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder) Post(body OrderEditUpdate) *ByProjectKeyOrdersEditsKeyByKeyRequestMethodPost {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deletes an OrderEdit in the Project.
 */
func (rb *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder) Delete() *ByProjectKeyOrdersEditsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/orders/edits/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
