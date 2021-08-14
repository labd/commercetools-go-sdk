// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyReviewsByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get Review by ID
 */
func (rb *ByProjectKeyReviewsByIdRequestBuilder) Get() *ByProjectKeyReviewsByIdRequestMethodGet {
	return &ByProjectKeyReviewsByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update Review by ID
 */
func (rb *ByProjectKeyReviewsByIdRequestBuilder) Post(body ReviewUpdate) *ByProjectKeyReviewsByIdRequestMethodPost {
	return &ByProjectKeyReviewsByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete Review by ID
 */
func (rb *ByProjectKeyReviewsByIdRequestBuilder) Delete() *ByProjectKeyReviewsByIdRequestMethodDelete {
	return &ByProjectKeyReviewsByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
