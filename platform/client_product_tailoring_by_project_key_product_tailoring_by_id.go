package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyProductTailoringByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyProductTailoringByIDRequestBuilder) Get() *ByProjectKeyProductTailoringByIDRequestMethodGet {
	return &ByProjectKeyProductTailoringByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-tailoring/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyProductTailoringByIDRequestBuilder) Post(body ProductTailoringUpdate) *ByProjectKeyProductTailoringByIDRequestMethodPost {
	return &ByProjectKeyProductTailoringByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/product-tailoring/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Generates the [ProductTailoringDeleted](ctp:api:type:ProductTailoringDeletedMessage) Message.
*
 */
func (rb *ByProjectKeyProductTailoringByIDRequestBuilder) Delete() *ByProjectKeyProductTailoringByIDRequestMethodDelete {
	return &ByProjectKeyProductTailoringByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/product-tailoring/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
