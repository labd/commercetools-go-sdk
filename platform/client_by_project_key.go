package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Categories are used to organize products in a hierarchical structure.
 */
func (rb *ByProjectKeyRequestBuilder) Categories() *ByProjectKeyCategoriesRequestBuilder {
	return &ByProjectKeyCategoriesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	A shopping cart holds product variants and can be ordered.
 */
func (rb *ByProjectKeyRequestBuilder) Carts() *ByProjectKeyCartsRequestBuilder {
	return &ByProjectKeyCartsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Cart discounts are used to change the prices of different elements within a cart.
 */
func (rb *ByProjectKeyRequestBuilder) CartDiscounts() *ByProjectKeyCartDiscountsRequestBuilder {
	return &ByProjectKeyCartDiscountsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Channels represent a source or destination of different entities. They can be used to model warehouses or stores.
*
 */
func (rb *ByProjectKeyRequestBuilder) Channels() *ByProjectKeyChannelsRequestBuilder {
	return &ByProjectKeyChannelsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	A customer is a person purchasing products. customers, Orders, Comments and Reviews can be associated to a customer.
*
 */
func (rb *ByProjectKeyRequestBuilder) Customers() *ByProjectKeyCustomersRequestBuilder {
	return &ByProjectKeyCustomersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	customer-groups are used to evaluate products and channels.
 */
func (rb *ByProjectKeyRequestBuilder) CustomerGroups() *ByProjectKeyCustomerGroupsRequestBuilder {
	return &ByProjectKeyCustomerGroupsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Store custom JSON values.
 */
func (rb *ByProjectKeyRequestBuilder) CustomObjects() *ByProjectKeyCustomObjectsRequestBuilder {
	return &ByProjectKeyCustomObjectsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Discount codes can be added to a discount-code to enable certain discount-code discounts.
 */
func (rb *ByProjectKeyRequestBuilder) DiscountCodes() *ByProjectKeyDiscountCodesRequestBuilder {
	return &ByProjectKeyDiscountCodesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The commercetools™ platform provides a GraphQL API
 */
func (rb *ByProjectKeyRequestBuilder) Graphql() *ByProjectKeyGraphqlRequestBuilder {
	return &ByProjectKeyGraphqlRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Inventory allows you to track stock quantities.
 */
func (rb *ByProjectKeyRequestBuilder) Inventory() *ByProjectKeyInventoryRequestBuilder {
	return &ByProjectKeyInventoryRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Retrieves the authenticated customer.
 */
func (rb *ByProjectKeyRequestBuilder) Login() *ByProjectKeyLoginRequestBuilder {
	return &ByProjectKeyLoginRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	A message represents a change or an action performed on a resource (like an Order or a Product).
 */
func (rb *ByProjectKeyRequestBuilder) Messages() *ByProjectKeyMessagesRequestBuilder {
	return &ByProjectKeyMessagesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	An order can be created from a order, usually after a checkout process has been completed.
 */
func (rb *ByProjectKeyRequestBuilder) Orders() *ByProjectKeyOrdersRequestBuilder {
	return &ByProjectKeyOrdersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Payments hold information about the current state of receiving and/or refunding money
 */
func (rb *ByProjectKeyRequestBuilder) Payments() *ByProjectKeyPaymentsRequestBuilder {
	return &ByProjectKeyPaymentsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Products are the sellable goods in an e-commerce project on CTP. This document explains some design concepts
*	of products on CTP and describes the available HTTP APIs for working with them.
*
 */
func (rb *ByProjectKeyRequestBuilder) Products() *ByProjectKeyProductsRequestBuilder {
	return &ByProjectKeyProductsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Product discounts are used to change certain product prices.
 */
func (rb *ByProjectKeyRequestBuilder) ProductDiscounts() *ByProjectKeyProductDiscountsRequestBuilder {
	return &ByProjectKeyProductDiscountsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	A projected representation of a product shows the product with its current or staged data. The current or staged
*	representation of a product in a catalog is called a product projection.
*
 */
func (rb *ByProjectKeyRequestBuilder) ProductProjections() *ByProjectKeyProductProjectionsRequestBuilder {
	return &ByProjectKeyProductProjectionsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Manage individual Store assortments through Product Selections.
*
*	After you have created Product Selections and populated them by Products,
*	you can manage Store assortments by assigning Product Selections to Stores.
*	Product Selections may be used by a single Store or shared across several Stores.
*
 */
func (rb *ByProjectKeyRequestBuilder) ProductSelections() *ByProjectKeyProductSelectionsRequestBuilder {
	return &ByProjectKeyProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Product Types are used to describe common characteristics, most importantly common custom attributes,
*	of many concrete products.
*
 */
func (rb *ByProjectKeyRequestBuilder) ProductTypes() *ByProjectKeyProductTypesRequestBuilder {
	return &ByProjectKeyProductTypesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Reviews are used to evaluate products and channels.
 */
func (rb *ByProjectKeyRequestBuilder) Reviews() *ByProjectKeyReviewsRequestBuilder {
	return &ByProjectKeyReviewsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Shipping Methods define where orders can be shipped and what the costs are.
 */
func (rb *ByProjectKeyRequestBuilder) ShippingMethods() *ByProjectKeyShippingMethodsRequestBuilder {
	return &ByProjectKeyShippingMethodsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	shopping-lists e.g. for wishlist support
 */
func (rb *ByProjectKeyRequestBuilder) ShoppingLists() *ByProjectKeyShoppingListsRequestBuilder {
	return &ByProjectKeyShoppingListsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	The commercetools platform allows you to model states of certain objects, such as orders, line items, products,
*	reviews, and payments in order to define finite state machines reflecting the business logic you'd like to
*	implement.
*
 */
func (rb *ByProjectKeyRequestBuilder) States() *ByProjectKeyStatesRequestBuilder {
	return &ByProjectKeyStatesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Subscriptions allow you to be notified of new messages or changes via a Message Queue of your choice
 */
func (rb *ByProjectKeyRequestBuilder) Subscriptions() *ByProjectKeySubscriptionsRequestBuilder {
	return &ByProjectKeySubscriptionsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Tax Categories define how products are to be taxed in different countries.
 */
func (rb *ByProjectKeyRequestBuilder) TaxCategories() *ByProjectKeyTaxCategoriesRequestBuilder {
	return &ByProjectKeyTaxCategoriesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Types define custom fields that are used to enhance resources as you need.
 */
func (rb *ByProjectKeyRequestBuilder) Types() *ByProjectKeyTypesRequestBuilder {
	return &ByProjectKeyTypesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Zones allow defining ShippingRates for specific Locations.
 */
func (rb *ByProjectKeyRequestBuilder) Zones() *ByProjectKeyZonesRequestBuilder {
	return &ByProjectKeyZonesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Me() *ByProjectKeyMeRequestBuilder {
	return &ByProjectKeyMeRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Extend the behavior of an API with your business logic
 */
func (rb *ByProjectKeyRequestBuilder) Extensions() *ByProjectKeyExtensionsRequestBuilder {
	return &ByProjectKeyExtensionsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Manage your API Clients via an API. Useful for Infrastructure-as-Code tooling, and regularly rotating API secrets.
*
 */
func (rb *ByProjectKeyRequestBuilder) ApiClients() *ByProjectKeyApiClientsRequestBuilder {
	return &ByProjectKeyApiClientsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Stores let you model the context your customers shop in.
 */
func (rb *ByProjectKeyRequestBuilder) Stores() *ByProjectKeyStoresRequestBuilder {
	return &ByProjectKeyStoresRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) InStoreKeyWithStoreKeyValue(storeKey string) *ByProjectKeyInStoreKeyByStoreKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyRequestBuilder{
		storeKey:   storeKey,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Get() *ByProjectKeyRequestMethodGet {
	return &ByProjectKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyRequestBuilder) Post(body ProjectUpdate) *ByProjectKeyRequestMethodPost {
	return &ByProjectKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s", rb.projectKey),
		client: rb.client,
	}
}
