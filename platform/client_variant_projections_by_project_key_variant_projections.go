package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantProjectionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyVariantProjectionsRequestBuilder) WithKey(key string) *ByProjectKeyVariantProjectionsKeyByKeyRequestBuilder {
	return &ByProjectKeyVariantProjectionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyVariantProjectionsRequestBuilder) WithId(id string) *ByProjectKeyVariantProjectionsByIDRequestBuilder {
	return &ByProjectKeyVariantProjectionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Queries Variant Projections matching the provided predicates.
*
*	By default, this endpoint returns the `current` representation where variants are published.
*
*	Required access scopes:
*
*	- To retrieve the current representation of published Variants, the `view_published_products:{projectKey}` scope is required.
*
*	- To retrieve the staged representation (draft data), the API Client must have the `view_products:{projectKey}` scope.
*
 */
func (rb *ByProjectKeyVariantProjectionsRequestBuilder) Get() *ByProjectKeyVariantProjectionsRequestMethodGet {
	return &ByProjectKeyVariantProjectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/variant-projections", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Variant Projection exists for the provided query predicate.
*	Returns a `200` status if any Variant Projections match, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantProjectionsRequestBuilder) Head() *ByProjectKeyVariantProjectionsRequestMethodHead {
	return &ByProjectKeyVariantProjectionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/variant-projections", rb.projectKey),
		client: rb.client,
	}
}
