// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTaxCategoriesByIDRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestBuilder) Get() *ByProjectKeyTaxCategoriesByIDRequestMethodGet {
	return &ByProjectKeyTaxCategoriesByIDRequestMethodGet{
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestBuilder) Post(body TaxCategoryUpdate) *ByProjectKeyTaxCategoriesByIDRequestMethodPost {
	return &ByProjectKeyTaxCategoriesByIDRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTaxCategoriesByIDRequestBuilder) Delete() *ByProjectKeyTaxCategoriesByIDRequestMethodDelete {
	return &ByProjectKeyTaxCategoriesByIDRequestMethodDelete{
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
