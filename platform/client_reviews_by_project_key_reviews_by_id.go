// Generated file, please do not change!!!
package platform

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
