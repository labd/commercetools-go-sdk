// Generated file, please do not change!!!
package ml

import (
	"fmt"
)

type ByProjectKeyMissingDataAttributesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyMissingDataAttributesRequestBuilder) Status() *ByProjectKeyMissingDataAttributesStatusRequestBuilder {
	return &ByProjectKeyMissingDataAttributesStatusRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyMissingDataAttributesRequestBuilder) Post(body MissingAttributesSearchRequest) *ByProjectKeyMissingDataAttributesRequestMethodPost {
	return &ByProjectKeyMissingDataAttributesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/missing-data/attributes", rb.projectKey),
		client: rb.client,
	}
}
