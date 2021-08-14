// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersEditsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersEditsRequestBuilder) WithKey(key string) *ByProjectKeyOrdersEditsKeyByKeyRequestBuilder {
	return &ByProjectKeyOrdersEditsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersEditsRequestBuilder) WithId(id string) *ByProjectKeyOrdersEditsByIdRequestBuilder {
	return &ByProjectKeyOrdersEditsByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query edits
 */
func (rb *ByProjectKeyOrdersEditsRequestBuilder) Get() *ByProjectKeyOrdersEditsRequestMethodGet {
	return &ByProjectKeyOrdersEditsRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders/edits", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create OrderEdit
 */
func (rb *ByProjectKeyOrdersEditsRequestBuilder) Post(body OrderEditDraft) *ByProjectKeyOrdersEditsRequestMethodPost {
	return &ByProjectKeyOrdersEditsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders/edits", rb.projectKey),
		client: rb.client,
	}
}
