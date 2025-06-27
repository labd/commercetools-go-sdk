package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more CartDiscounts exist for the provided query predicate. Returns a `200 OK` status if any CartDiscounts match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	When using the endpoint, the Store specified in the path and the Stores specified in the payload's `stores` field are added to the CartDiscount.
*
*	Creating a Cart Discount produces the [CartDiscountCreated](ctp:api:type:CartDiscountCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestBuilder) Post(body CartDiscountDraft) *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
