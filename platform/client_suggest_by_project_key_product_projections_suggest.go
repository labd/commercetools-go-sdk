// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyProductProjectionsSuggestRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductProjectionsSuggestRequestBuilder) Get() *ByProjectKeyProductProjectionsSuggestRequestMethodGet {
	return &ByProjectKeyProductProjectionsSuggestRequestMethodGet{
		url:    fmt.Sprintf("/%s/product-projections/suggest", rb.projectKey),
		client: rb.client,
	}
}
