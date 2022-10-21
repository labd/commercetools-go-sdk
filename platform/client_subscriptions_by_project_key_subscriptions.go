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
*	A test message is sent to ensure the correct configuration of the Destination. If the message cannot be delivered, the Subscription will not be created. The payload of the test message is a notification of type [ResourceCreated](/../api/projects/subscriptions#resourcecreateddeliverypayload) for the `resourceTypeId` `subscription`.
*
 */
func (rb *ByProjectKeySubscriptionsRequestBuilder) Post(body SubscriptionDraft) *ByProjectKeySubscriptionsRequestMethodPost {
	return &ByProjectKeySubscriptionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/subscriptions", rb.projectKey),
		client: rb.client,
	}
}
