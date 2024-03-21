package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder struct {
	projectKey string
	storeKey   string
	productKey string
	client     *Client
}

/**
*	Gets the current or staged representation of a [Product Tailoring](ctp:api:type:ProductTailoring) by its Product key in the specified [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/key=%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productKey),
		client: rb.client,
	}
}

/**
*	Updates the current or staged representation of a [Product Tailoring](ctp:api:type:ProductTailoring) by its Product key in the specified [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder) Post(body ProductTailoringUpdate) *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/key=%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productKey),
		client: rb.client,
	}
}

/**
*	Generates the [ProductTailoringDeleted](ctp:api:type:ProductTailoringDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyProductsKeyByProductKeyProductTailoringRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/products/key=%s/product-tailoring", rb.projectKey, rb.storeKey, rb.productKey),
		client: rb.client,
	}
}
