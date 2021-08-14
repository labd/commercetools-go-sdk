// Generated file, please do not change!!!
package platform

import (
	"fmt"
	"io"
)

type ByProjectKeyProductsByIdImagesRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Uploads a binary image file to a given product variant. The supported image formats are JPEG, PNG and GIF.
*
 */
func (rb *ByProjectKeyProductsByIdImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyProductsByIdImagesRequestMethodPost {
	return &ByProjectKeyProductsByIdImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s/images", rb.projectKey, rb.id),
		client: rb.client,
	}
}
