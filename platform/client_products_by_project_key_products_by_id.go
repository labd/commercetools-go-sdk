package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductsByIDRequestBuilder) Images() *ByProjectKeyProductsByIDImagesRequestBuilder {
	return &ByProjectKeyProductsByIDImagesRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductsByIDRequestBuilder) ProductSelections() *ByProjectKeyProductsByIDProductSelectionsRequestBuilder {
	return &ByProjectKeyProductsByIDProductSelectionsRequestBuilder{
		projectKey: rb.projectKey,
		id:         rb.id,
		client:     rb.client,
	}
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Get() *ByProjectKeyProductsByIDRequestMethodGet {
	return &ByProjectKeyProductsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Check if a Product exists with a specified `id`. Responds with a `200 OK` status if the Product exists or `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Head() *ByProjectKeyProductsByIDRequestMethodHead {
	return &ByProjectKeyProductsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Post(body ProductUpdate) *ByProjectKeyProductsByIDRequestMethodPost {
	return &ByProjectKeyProductsByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	If [Price selection](ctp:api:type:ProductPriceSelection) query parameters are provided, the selected Prices are added to the response.
*	Produces the [ProductDeletedMessage](/message-types#productdeletedmessage).
 */
func (rb *ByProjectKeyProductsByIDRequestBuilder) Delete() *ByProjectKeyProductsByIDRequestMethodDelete {
	return &ByProjectKeyProductsByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/products/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
