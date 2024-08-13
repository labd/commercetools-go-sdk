package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder struct {
	projectKey string
	storeKey   string
	productId  string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder) Images() *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringImagesRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringImagesRequestBuilder{
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		productId:  rb.productId,
		client:     rb.client,
	}
}

/**
*	Gets the current or staged representation of a [Product Tailoring](ctp:api:type:ProductTailoring) by its Product ID in the specified [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productId),
		client: rb.client,
	}
}

/**
*	Updates the current or staged representation of a [Product Tailoring](ctp:api:type:ProductTailoring) by its Product ID in the specified [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder) Post(body ProductTailoringUpdate) *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productId),
		client: rb.client,
	}
}

/**
*	Generates the [ProductTailoringDeleted](ctp:api:type:ProductTailoringDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsByProductIDProductTailoringRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productId),
		client: rb.client,
	}
}
