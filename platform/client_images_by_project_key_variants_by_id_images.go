package platform

// Generated file, please do not change!!!

import (
	"fmt"
	"io"
)

type ByProjectKeyVariantsByIDImagesRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Uploads a JPEG, PNG, or a GIF image file to a [Variant](ctp:api:type:Variant) identified by its ID.
*	The maximum file size of the image is **10MB**.
*
*	The response status code depends on the size of the original image.
*	If the image is small, the API responds with `200 OK`, and if the image is larger, it responds with `202 Accepted`.
*	The Variant returned with a `202 Accepted` status code contains a `warnings` field with an [ImageProcessingOngoing](ctp:api:type:ImageProcessingOngoingWarning) Warning.
*
*	Produces the [VariantImageAdded](/projects/messages/product-catalog-messages#variant-image-added) Message.
*
 */
func (rb *ByProjectKeyVariantsByIDImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyVariantsByIDImagesRequestMethodPost {
	return &ByProjectKeyVariantsByIDImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants/%s/images", rb.projectKey, rb.id),
		client: rb.client,
	}
}
