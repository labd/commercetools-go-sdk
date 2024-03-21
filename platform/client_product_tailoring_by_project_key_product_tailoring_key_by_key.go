package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTailoringKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

func (rb *ByProjectKeyProductTailoringKeyByKeyRequestBuilder) Get() *ByProjectKeyProductTailoringKeyByKeyRequestMethodGet {
	return &ByProjectKeyProductTailoringKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-tailoring/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTailoringKeyByKeyRequestBuilder) Post(body ProductTailoringUpdate) *ByProjectKeyProductTailoringKeyByKeyRequestMethodPost {
	return &ByProjectKeyProductTailoringKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-tailoring/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Generates the [ProductTailoringDeleted](ctp:api:type:ProductTailoringDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyProductTailoringKeyByKeyRequestBuilder) Delete() *ByProjectKeyProductTailoringKeyByKeyRequestMethodDelete {
	return &ByProjectKeyProductTailoringKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-tailoring/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
