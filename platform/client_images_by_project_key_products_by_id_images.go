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
*	Uploads a JPEG, PNG, or a GIF image file to a [ProductVariant](ctp:api:type:ProductVariant).
*	The maximum file size of the image is **10MB**.
*	Either `variant` or `sku` is required to update a specific ProductVariant.
*	If neither is provided, the image is uploaded to the Master Variant of the Product.
*
*	The response status code depends on the size of the original image.
*	If the image is small, the API responds with `200 OK`, and if the image is larger, it responds with `202 Accepted`.
*	The Product returned with a `202 Accepted` status code contains a `warnings` field with an [ImageProcessingOngoing](ctp:api:type:ImageProcessingOngoingWarning) Warning.
*
*	Produces the [ProductImageAdded](/projects/messages/product-catalog-messages#product-image-added) Message.
*
 */
func (rb *ByProjectKeyProductsByIDImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyProductsByIDImagesRequestMethodPost {
	return &ByProjectKeyProductsByIDImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s/images", rb.projectKey, rb.id),
		client: rb.client,
	}
}
