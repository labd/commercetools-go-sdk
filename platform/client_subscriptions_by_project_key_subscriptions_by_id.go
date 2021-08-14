// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeySubscriptionsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves the representation of a subscription by its id.
 */
func (rb *ByProjectKeySubscriptionsByIdRequestBuilder) Get() *ByProjectKeySubscriptionsByIdRequestMethodGet {
	return &ByProjectKeySubscriptionsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Subscription by ID
 */
func (rb *ByProjectKeySubscriptionsByIdRequestBuilder) Post(body SubscriptionUpdate) *ByProjectKeySubscriptionsByIdRequestMethodPost {
	return &ByProjectKeySubscriptionsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Subscription by ID
 */
func (rb *ByProjectKeySubscriptionsByIdRequestBuilder) Delete() *ByProjectKeySubscriptionsByIdRequestMethodDelete {
	return &ByProjectKeySubscriptionsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
