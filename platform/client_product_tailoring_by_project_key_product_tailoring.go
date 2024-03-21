package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTailoringRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductTailoringRequestBuilder) WithKey(key string) *ByProjectKeyProductTailoringKeyByKeyRequestBuilder {
	return &ByProjectKeyProductTailoringKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductTailoringRequestBuilder) WithId(id string) *ByProjectKeyProductTailoringByIDRequestBuilder {
	return &ByProjectKeyProductTailoringByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductTailoringRequestBuilder) Get() *ByProjectKeyProductTailoringRequestMethodGet {
	return &ByProjectKeyProductTailoringRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-tailoring", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Generates the [ProductTailoringCreated](ctp:api:type:ProductTailoringCreatedMessage) Message.
*
 */
func (rb *ByProjectKeyProductTailoringRequestBuilder) Post(body ProductTailoringDraft) *ByProjectKeyProductTailoringRequestMethodPost {
	return &ByProjectKeyProductTailoringRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-tailoring", rb.projectKey),
		client: rb.client,
	}
}
