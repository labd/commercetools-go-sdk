package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeActiveCartRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves the Customer's most recently modified active Cart.
*	Carts with `Merchant` or `Quote` [CartOrigin](ctp:api:type:CartOrigin) are ignored.
*
*	If no active Cart exists, a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned.
*
 */
func (rb *ByProjectKeyMeActiveCartRequestBuilder) Get() *ByProjectKeyMeActiveCartRequestMethodGet {
	return &ByProjectKeyMeActiveCartRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/active-cart", rb.projectKey),
		client: rb.client,
	}
}
