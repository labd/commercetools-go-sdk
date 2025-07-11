package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTypesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTypesRequestBuilder) WithKey(key string) *ByProjectKeyTypesKeyByKeyRequestBuilder {
	return &ByProjectKeyTypesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTypesRequestBuilder) WithId(id string) *ByProjectKeyTypesByIDRequestBuilder {
	return &ByProjectKeyTypesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTypesRequestBuilder) Get() *ByProjectKeyTypesRequestMethodGet {
	return &ByProjectKeyTypesRequestMethodGet{
		url:    fmt.Sprintf("/%s/types", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if one or more Types exist for the provided query predicate. Returns a `200 OK` status if any Types match the query predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyTypesRequestBuilder) Head() *ByProjectKeyTypesRequestMethodHead {
	return &ByProjectKeyTypesRequestMethodHead{
		url:    fmt.Sprintf("/%s/types", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTypesRequestBuilder) Post(body TypeDraft) *ByProjectKeyTypesRequestMethodPost {
	return &ByProjectKeyTypesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/types", rb.projectKey),
		client: rb.client,
	}
}
