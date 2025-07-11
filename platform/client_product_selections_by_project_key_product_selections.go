package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductSelectionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductSelectionsRequestBuilder) WithKey(key string) *ByProjectKeyProductSelectionsKeyByKeyRequestBuilder {
	return &ByProjectKeyProductSelectionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductSelectionsRequestBuilder) WithId(id string) *ByProjectKeyProductSelectionsByIDRequestBuilder {
	return &ByProjectKeyProductSelectionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductSelectionsRequestBuilder) Get() *ByProjectKeyProductSelectionsRequestMethodGet {
	return &ByProjectKeyProductSelectionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-selections", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more ProductSelections exist for the provided query predicate. Returns a `200 OK` status if any ProductSelections match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyProductSelectionsRequestBuilder) Head() *ByProjectKeyProductSelectionsRequestMethodHead {
	return &ByProjectKeyProductSelectionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/product-selections", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductSelectionsRequestBuilder) Post(body ProductSelectionDraft) *ByProjectKeyProductSelectionsRequestMethodPost {
	return &ByProjectKeyProductSelectionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-selections", rb.projectKey),
		client: rb.client,
	}
}
