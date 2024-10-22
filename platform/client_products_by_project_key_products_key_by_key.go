package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) ProductSelections() *ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder {
	return &ByProjectKeyProductsKeyByKeyProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		key:        rb.key,
		client:     rb.client,
	}
}

/**
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Get() *ByProjectKeyProductsKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a Product exists for a given `key`. Returns a `200 OK` status if the Product exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Head() *ByProjectKeyProductsKeyByKeyRequestMethodHead {
	return &ByProjectKeyProductsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	A failed response can return a [DuplicatePriceScope](ctp:api:type:DuplicatePriceScopeError), [DuplicateVariantValues](ctp:api:type:DuplicateVariantValuesError), [DuplicateAttributeValue](ctp:api:type:DuplicateAttributeValueError), or [DuplicateAttributeValues](ctp:api:type:DuplicateAttributeValuesError) error.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductsKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	If [Product price selection query parameters](/../api/pricing-and-discounts-overview#product-price-selection) are provided, the selected Prices are added to the response.
*	Produces the [ProductDeleted](/projects/messages/product-catalog-messages#product-deleted) Message.
 */
func (rb *ByProjectKeyProductsKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductsKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductsKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
