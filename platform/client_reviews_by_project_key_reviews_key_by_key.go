// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyReviewsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get Review by key
 */
func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Get() *ByProjectKeyReviewsKeyByKeyRequestMethodGet {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update Review by key
 */
func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Post(body ReviewUpdate) *ByProjectKeyReviewsKeyByKeyRequestMethodPost {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete Review by key
 */
func (rb *ByProjectKeyReviewsKeyByKeyRequestBuilder) Delete() *ByProjectKeyReviewsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyReviewsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/reviews/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
