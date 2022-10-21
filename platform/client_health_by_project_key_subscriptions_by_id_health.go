package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeySubscriptionsByIDHealthRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	This endpoint can be polled by a monitoring or alerting system that checks the health of your Subscriptions. To ease integration with such systems this endpoint does not require [Authorization](/../api/authorization).
*
 */
func (rb *ByProjectKeySubscriptionsByIDHealthRequestBuilder) Get() *ByProjectKeySubscriptionsByIDHealthRequestMethodGet {
	return &ByProjectKeySubscriptionsByIDHealthRequestMethodGet{
		url:    fmt.Sprintf("/%s/subscriptions/%s/health", rb.projectKey, rb.id),
		client: rb.client,
	}
}
