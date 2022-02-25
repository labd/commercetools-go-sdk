package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyReviewsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Get() *ByProjectKeyReviewsKeyByKeyRequestMethodGet {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Post(body ReviewUpdate) *ByProjectKeyReviewsKeyByKeyRequestMethodPost {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Delete() *ByProjectKeyReviewsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
