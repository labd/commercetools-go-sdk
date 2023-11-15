package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeySubscriptionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) WithIdHealth() *ByProjectKeySubscriptionsByIDHealthRequestBuilder {
	return &ByProjectKeySubscriptionsByIDHealthRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Get() *ByProjectKeySubscriptionsByIDRequestMethodGet {
	return &ByProjectKeySubscriptionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Subscription exists for a given `id`. Returns a `200 OK` status if the Subscription exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Head() *ByProjectKeySubscriptionsByIDRequestMethodHead {
	return &ByProjectKeySubscriptionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Post(body SubscriptionUpdate) *ByProjectKeySubscriptionsByIDRequestMethodPost {
	return &ByProjectKeySubscriptionsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsByIDRequestBuilder) Delete() *ByProjectKeySubscriptionsByIDRequestMethodDelete {
	return &ByProjectKeySubscriptionsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/subscriptions/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
