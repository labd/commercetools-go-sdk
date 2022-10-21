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
*	Upload a JPEG, PNG and GIF file to a [ProductVariant](ctp:api:type:ProductVariant). The maximum file size of the image is 10MB. `variant` or `sku` is required to update a specific ProductVariant. The image is uploaded to the Master Variant if `variant` or `sku` are not included. Produces the [ProductImageAdded](/projects/messages#product-image-added) Message when the `Small` version of the image has been uploaded to the CDN.
*
 */
func (rb *ByProjectKeyProductsByIDImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyProductsByIDImagesRequestMethodPost {
	return &ByProjectKeyProductsByIDImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s/images", rb.projectKey, rb.id),
		client: rb.client,
	}
}
