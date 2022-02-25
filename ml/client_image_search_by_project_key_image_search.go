package ml

// Generated file, please do not change!!!

import (
	"fmt"
	"io"
)

type ByProjectKeyImageSearchRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyImageSearchRequestBuilder) Config() *ByProjectKeyImageSearchConfigRequestBuilder {
	return &ByProjectKeyImageSearchConfigRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Accepts an image file and returns similar products from product catalogue.
*
 */
func (rb *ByProjectKeyImageSearchRequestBuilder) Post(body io.Reader) *ByProjectKeyImageSearchRequestMethodPost {
	return &ByProjectKeyImageSearchRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/image-search", rb.projectKey),
		client: rb.client,
	}
}
