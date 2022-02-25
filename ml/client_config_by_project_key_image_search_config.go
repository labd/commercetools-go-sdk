package ml

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyImageSearchConfigRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Get the current image search config.
 */
func (rb *ByProjectKeyImageSearchConfigRequestBuilder) Get() *ByProjectKeyImageSearchConfigRequestMethodGet {
	return &ByProjectKeyImageSearchConfigRequestMethodGet{
		url:    fmt.Sprintf("/%s/image-search/config", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Endpoint to update the image search config.
 */
func (rb *ByProjectKeyImageSearchConfigRequestBuilder) Post(body ImageSearchConfigRequest) *ByProjectKeyImageSearchConfigRequestMethodPost {
	return &ByProjectKeyImageSearchConfigRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/image-search/config", rb.projectKey),
		client: rb.client,
	}
}
