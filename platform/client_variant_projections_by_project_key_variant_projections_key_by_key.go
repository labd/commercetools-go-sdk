package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantProjectionsKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Retrieves a Variant Projection by its key.
*
 */
func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestBuilder) Get() *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet {
	return &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/variant-projections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a VariantProjection exists with the provided `key`. Returns a `200` status if the VariantProjection exists, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantProjectionsKeyByKeyRequestBuilder) Head() *ByProjectKeyVariantProjectionsKeyByKeyRequestMethodHead {
	return &ByProjectKeyVariantProjectionsKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/variant-projections/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
