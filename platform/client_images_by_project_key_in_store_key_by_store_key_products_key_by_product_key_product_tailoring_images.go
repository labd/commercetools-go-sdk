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
*	Upload a JPEG, PNG and GIF file to a [ProductTailoringVariant](ctp:api:type:ProductTailoringVariant). The maximum file size of the image is 10MB. `variant` or `sku` is required to update a specific ProductVariant. Produces the [ProductTailoringImageAdded](/projects/messages#product-tailoring-image-added) Message when the `Small` version of the image has been uploaded to the CDN.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestBuilder) Post(body io.Reader) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/key=%s/product-tailoring/images", rb.projectKey, rb.storeKey, rb.productKey),
		client: rb.client,
	}
}
