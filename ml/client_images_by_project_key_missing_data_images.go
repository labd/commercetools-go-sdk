package ml

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyMissingDataImagesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataImagesRequestBuilder) Status() *ByProjectKeyMissingDataImagesStatusRequestBuilder {
	return &ByProjectKeyMissingDataImagesStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMissingDataImagesRequestBuilder) Post(body MissingImagesSearchRequest) *ByProjectKeyMissingDataImagesRequestMethodPost {
	return &ByProjectKeyMissingDataImagesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/missing-data/images", rb.projectKey),
		client: rb.client,
	}
}
