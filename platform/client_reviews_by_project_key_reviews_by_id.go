package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyReviewsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyReviewsByIDRequestBuilder) Get() *ByProjectKeyReviewsByIDRequestMethodGet {
	return &ByProjectKeyReviewsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a Review exists for a given `id`. Returns a `200 OK` status if the Review exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyReviewsByIDRequestBuilder) Head() *ByProjectKeyReviewsByIDRequestMethodHead {
	return &ByProjectKeyReviewsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyReviewsByIDRequestBuilder) Post(body ReviewUpdate) *ByProjectKeyReviewsByIDRequestMethodPost {
	return &ByProjectKeyReviewsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyReviewsByIDRequestBuilder) Delete() *ByProjectKeyReviewsByIDRequestMethodDelete {
	return &ByProjectKeyReviewsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/reviews/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
