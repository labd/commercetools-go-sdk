package platform

// Generated file, please do not change!!!

import (
	"fmt"
	"io"
)

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestBuilder struct {
	projectKey string
	storeKey   string
	productKey string
	client     *Client
}

/**
*	Uploads a JPEG, PNG and GIF file to a [ProductVariantTailoring](ctp:api:type:ProductVariantTailoring).
*	The maximum file size of the image is **10MB**.
*	Either `variant` or `sku` is required to update a specific ProductVariant.
*	If neither is provided, the image is uploaded to the Master Variant of the Product.
*
*	The response status code depends on the size of the original image.
*	If the image is small, the API responds with `200 OK`, and if the image is larger, it responds with `202 Accepted`.
*	The Product returned with a `202 Accepted` status code contains a `warnings` field with an [ImageProcessingOngoing](ctp:api:type:ImageProcessingOngoingWarning) Warning.
*
*	Produces the [ProductTailoringImageAdded](/projects/messages/product-catalog-messages#product-tailoring-image-added) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/key=%s/product-tailoring/images", rb.projectKey, rb.storeKey, rb.productKey),
		client: rb.client,
	}
}
