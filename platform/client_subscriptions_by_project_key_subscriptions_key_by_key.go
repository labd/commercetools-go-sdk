package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeySubscriptionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeySubscriptionsKeyByKeyRequestBuilder) Get() *ByProjectKeySubscriptionsKeyByKeyRequestMethodGet {
	return &ByProjectKeySubscriptionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Subscription exists for a given `key`. Returns a `200 OK` status if the Subscription exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeySubscriptionsKeyByKeyRequestBuilder) Head() *ByProjectKeySubscriptionsKeyByKeyRequestMethodHead {
	return &ByProjectKeySubscriptionsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/subscriptions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsKeyByKeyRequestBuilder) Post(body SubscriptionUpdate) *ByProjectKeySubscriptionsKeyByKeyRequestMethodPost {
	return &ByProjectKeySubscriptionsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeySubscriptionsKeyByKeyRequestBuilder) Delete() *ByProjectKeySubscriptionsKeyByKeyRequestMethodDelete {
	return &ByProjectKeySubscriptionsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/subscriptions/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
