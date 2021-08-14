// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyMeOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyMeOrdersByIdRequestBuilder {
	return &ByProjectKeyMeOrdersByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query orders
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Get() *ByProjectKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create Order
 */
func (rb *ByProjectKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}
