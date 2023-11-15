package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyReviewsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyReviewsRequestBuilder) WithKey(key string) *ByProjectKeyReviewsKeyByKeyRequestBuilder {
	return &ByProjectKeyReviewsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyReviewsRequestBuilder) WithId(id string) *ByProjectKeyReviewsByIDRequestBuilder {
	return &ByProjectKeyReviewsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyReviewsRequestBuilder) Get() *ByProjectKeyReviewsRequestMethodGet {
	return &ByProjectKeyReviewsRequestMethodGet{
		url:    fmt.Sprintf("/%s/reviews", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Review exists for a given Query Predicate. Returns a `200 OK` status if any Reviews match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyReviewsRequestBuilder) Head() *ByProjectKeyReviewsRequestMethodHead {
	return &ByProjectKeyReviewsRequestMethodHead{
		url:    fmt.Sprintf("/%s/reviews", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyReviewsRequestBuilder) Post(body ReviewDraft) *ByProjectKeyReviewsRequestMethodPost {
	return &ByProjectKeyReviewsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/reviews", rb.projectKey),
		client: rb.client,
	}
}
