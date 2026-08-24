package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantProjectionsByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Retrieves a Variant Projection by its ID.
*
 */
func (rb *ByProjectKeyVariantProjectionsByIDRequestBuilder) Get() *ByProjectKeyVariantProjectionsByIDRequestMethodGet {
	return &ByProjectKeyVariantProjectionsByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/variant-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Checks if a VariantProjection exists with the provided `id`. Returns a `200` status if the VariantProjection exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantProjectionsByIDRequestBuilder) Head() *ByProjectKeyVariantProjectionsByIDRequestMethodHead {
	return &ByProjectKeyVariantProjectionsByIDRequestMethodHead{
		url:    fmt.Sprintf("/%s/variant-projections/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
