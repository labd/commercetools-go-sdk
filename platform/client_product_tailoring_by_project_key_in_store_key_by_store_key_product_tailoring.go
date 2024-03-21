package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

/**
*	Queries Product Tailoring in a specific [Store](ctp:api:type:Store).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-tailoring", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Creates a [ProductTailoring](ctp:api:type:ProductTailoring) in the [Store](ctp:api:type:Store) specified by `storeKey`.
*	When using this endpoint the ProductTailoring's `store` field is always set to the [Store](ctp:api:type:Store) specified in the path parameter.
*
*	Generates the [ProductTailoringCreated](ctp:api:type:ProductTailoringCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestBuilder) Post(body ProductTailoringInStoreDraft) *ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyProductTailoringRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/product-tailoring", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
