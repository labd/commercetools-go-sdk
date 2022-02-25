package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeySubscriptionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeySubscriptionsRequestBuilder) WithKey(key string) *ByProjectKeySubscriptionsKeyByKeyRequestBuilder {
	return &ByProjectKeySubscriptionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeySubscriptionsRequestBuilder) WithId(id string) *ByProjectKeySubscriptionsByIDRequestBuilder {
	return &ByProjectKeySubscriptionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeySubscriptionsRequestBuilder) Get() *ByProjectKeySubscriptionsRequestMethodGet {
	return &ByProjectKeySubscriptionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions", rb.projectKey),
		client: rb.client,
	}
}

/**
*	The creation of a Subscription is eventually consistent, it may take up to a minute before it becomes fully active.
*	In order to test that the destination is correctly configured, a test message will be put into the queue.
*	If the message could not be delivered, the subscription will not be created.
*	The payload of the test message is a notification of type ResourceCreated for the resourceTypeId subscription.
*	Currently, a maximum of 25 subscriptions can be created per project.
*
 */
func (rb *ByProjectKeySubscriptionsRequestBuilder) Post(body SubscriptionDraft) *ByProjectKeySubscriptionsRequestMethodPost {
	return &ByProjectKeySubscriptionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions", rb.projectKey),
		client: rb.client,
	}
}
