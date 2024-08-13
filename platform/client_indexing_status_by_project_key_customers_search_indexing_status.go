package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersSearchIndexingStatusRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCustomersSearchIndexingStatusRequestBuilder) Get() *ByProjectKeyCustomersSearchIndexingStatusRequestMethodGet {
	return &ByProjectKeyCustomersSearchIndexingStatusRequestMethodGet{
		url:    fmt.Sprintf("/%s/customers/search/indexing-status", rb.projectKey),
		client: rb.client,
	}
}
