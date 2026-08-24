package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyVariantsBulkRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Applies update actions to multiple Variants in a single request.
*	Returns a `207 Multi-Status` response with the result for each Variant.
*	Each Variant is updated independently, so some may succeed while others fail.
*
 */
func (rb *ByProjectKeyVariantsBulkRequestBuilder) Post(body VariantBulkUpdate) *ByProjectKeyVariantsBulkRequestMethodPost {
	return &ByProjectKeyVariantsBulkRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/variants/bulk", rb.projectKey),
		client: rb.client,
	}
}
