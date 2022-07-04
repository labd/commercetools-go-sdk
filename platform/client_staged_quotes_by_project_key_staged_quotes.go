package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyStagedQuotesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyStagedQuotesRequestBuilder) WithKey(key string) *ByProjectKeyStagedQuotesKeyByKeyRequestBuilder {
	return &ByProjectKeyStagedQuotesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStagedQuotesRequestBuilder) WithId(id string) *ByProjectKeyStagedQuotesByIDRequestBuilder {
	return &ByProjectKeyStagedQuotesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyStagedQuotesRequestBuilder) Get() *ByProjectKeyStagedQuotesRequestMethodGet {
	return &ByProjectKeyStagedQuotesRequestMethodGet{
		url:    fmt.Sprintf("/%s/staged-quotes", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyStagedQuotesRequestBuilder) Post(body StagedQuoteDraft) *ByProjectKeyStagedQuotesRequestMethodPost {
	return &ByProjectKeyStagedQuotesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/staged-quotes", rb.projectKey),
		client: rb.client,
	}
}
