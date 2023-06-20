package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMeOrdersRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMeOrdersRequestBuilder) WithId(id string) *ByProjectKeyMeOrdersByIDRequestBuilder {
	return &ByProjectKeyMeOrdersByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeOrdersRequestBuilder) Quotes() *ByProjectKeyMeOrdersQuotesRequestBuilder {
	return &ByProjectKeyMeOrdersQuotesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMeOrdersRequestBuilder) Get() *ByProjectKeyMeOrdersRequestMethodGet {
	return &ByProjectKeyMeOrdersRequestMethodGet{
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyMeOrdersRequestBuilder) Post(body MyOrderFromCartDraft) *ByProjectKeyMeOrdersRequestMethodPost {
	return &ByProjectKeyMeOrdersRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/me/orders", rb.projectKey),
		client: rb.client,
	}
}
