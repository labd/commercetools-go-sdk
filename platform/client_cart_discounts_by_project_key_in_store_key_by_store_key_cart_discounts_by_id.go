package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder struct {
	projectKey string
	storeKey   string
	id         string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a CartDiscount exists for a given `id`. Returns a `200 OK` status if the CartDiscount exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	To update a CartDiscount, you must have permissions for all Stores the CartDiscount is associated with, except when [removing a Store](ctp:api:type:CartDiscountRemoveStoreAction).
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder) Post(body CartDiscountUpdate) *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}

/**
*	To delete a CartDiscount, specify the `manage_cart_discounts:{projectKey}:{storeKey}` scope for all Stores associated with the CartDiscount.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestBuilder) Delete() *ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodDelete {
	return &ByProjectKeyInStoreKeyByStoreKeyCartDiscountsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/in-store/key=%s/cart-discounts/%s", rb.projectKey, rb.storeKey, rb.id),
		client: rb.client,
	}
}
