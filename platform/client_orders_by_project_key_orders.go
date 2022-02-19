// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyOrdersRequestBuilder) ImportOrder() *ByProjectKeyOrdersImportRequestBuilder {
	return &ByProjectKeyOrdersImportRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) WithOrderNumber(orderNumber string) *ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder {
	return &ByProjectKeyOrdersOrderNumberByOrderNumberRequestBuilder{
		orderNumber: orderNumber,
		projectKey:  rb.projectKey,
		client:      rb.client,
	}
}

/**
*	OrderEdit are containers for financial changes after an Order has been placed.
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Edits() *ByProjectKeyOrdersEditsRequestBuilder {
	return &ByProjectKeyOrdersEditsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) WithId(id string) *ByProjectKeyOrdersByIDRequestBuilder {
	return &ByProjectKeyOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	This endpoint provides high performance search queries over Orders. The order search allows searching through all orders (currently supporting a limit of the 10.000.000 newest orders) in your project.
*
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Search() *ByProjectKeyOrdersSearchRequestBuilder {
	return &ByProjectKeyOrdersSearchRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyOrdersRequestBuilder) Get() *ByProjectKeyOrdersRequestMethodGet {
	return &ByProjectKeyOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/orders", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates an order from a Cart.
*	The cart must have a shipping address set before creating an order.
*	When using the Platform TaxMode, the shipping address is used for tax calculation.
*
 */
func (rb *ByProjectKeyOrdersRequestBuilder) Post(body OrderFromCartDraft) *ByProjectKeyOrdersRequestMethodPost {
	return &ByProjectKeyOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/orders", rb.projectKey),
		client: rb.client,
	}
}
