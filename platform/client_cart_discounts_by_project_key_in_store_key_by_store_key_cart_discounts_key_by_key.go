package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder struct {
	projectKey string
	storeKey   string
	key        string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a CartDiscount exists for a given `key`. Returns a `200 OK` status if the CartDiscount exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	To update a CartDiscount, you must have permissions for all Stores the CartDiscount is associated with, except when [removing a Store](ctp:api:type:CartDiscountRemoveStoreAction).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}

/**
*	To delete a CartDiscount, specify the `manage_cart_discounts:{projectKey}:{storeKey}` scope for all Stores associated with the CartDiscount.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/key=%s", rb.projectKey, rb.storeKey, rb.key),
		client: rb.client,
	}
}
