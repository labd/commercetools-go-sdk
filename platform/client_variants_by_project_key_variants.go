package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyVariantsRequestBuilder) Bulk() *ByProjectKeyVariantsBulkRequestBuilder {
	return &ByProjectKeyVariantsBulkRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyVariantsRequestBuilder) WithKey(key string) *ByProjectKeyVariantsKeyByKeyRequestBuilder {
	return &ByProjectKeyVariantsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyVariantsRequestBuilder) WithId(id string) *ByProjectKeyVariantsByIDRequestBuilder {
	return &ByProjectKeyVariantsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Queries Variants matching the provided predicates.
*	This endpoint returns both current (published) and staged (draft) data for each Variant.
*
*	When querying an Attribute, you must always include both the `attribute-name` and `attribute-value` details in the QueryPredicate. Otherwise, a [400 Bad Request](/errors#400-bad-request) error will occur.
*
 */
func (rb *ByProjectKeyVariantsRequestBuilder) Get() *ByProjectKeyVariantsRequestMethodGet {
	return &ByProjectKeyVariantsRequestMethodGet{
		url:    fmt.Sprintf("/%s/variants", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creates a new Variant with only the current representation. Produces the [VariantCreated](ctp:api:type:VariantCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyVariantsRequestBuilder) Post(body VariantDraft) *ByProjectKeyVariantsRequestMethodPost {
	return &ByProjectKeyVariantsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Variant exists for the provided query predicate.
*	Returns a `200` status if any Variants match, or a `404` status otherwise.
*
 */
func (rb *ByProjectKeyVariantsRequestBuilder) Head() *ByProjectKeyVariantsRequestMethodHead {
	return &ByProjectKeyVariantsRequestMethodHead{
		url:    fmt.Sprintf("/%s/variants", rb.projectKey),
		client: rb.client,
	}
}
