package platform

// Generated file, please do not change!!!

import (
	"fmt"
	"io"
)

type ByProjectKeyProductsByIDImagesRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Uploads a binary image file to a given product variant. The supported image formats are JPEG, PNG and GIF.
*
 */
func (rb *ByProjectKeyProductsByIDImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyProductsByIDImagesRequestMethodPost {
	return &ByProjectKeyProductsByIDImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s/images", rb.projectKey, rb.id),
		client: rb.client,
	}
}
