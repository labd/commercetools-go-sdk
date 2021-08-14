// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyCategoriesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Get() *ByProjectKeyCategoriesByIDRequestMethodGet {
	return &ByProjectKeyCategoriesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Post(body CategoryUpdate) *ByProjectKeyCategoriesByIDRequestMethodPost {
	return &ByProjectKeyCategoriesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyCategoriesByIDRequestBuilder) Delete() *ByProjectKeyCategoriesByIDRequestMethodDelete {
	return &ByProjectKeyCategoriesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
