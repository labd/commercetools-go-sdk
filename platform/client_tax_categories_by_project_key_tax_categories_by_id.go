// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTaxCategoriesByIdRequestBuilder struct {
	projectKey string
	id         string
	client     *Client
}

/**
*	Get TaxCategory by ID
 */
func (rb *ByProjectKeyTaxCategoriesByIdRequestBuilder) Get() *ByProjectKeyTaxCategoriesByIdRequestMethodGet {
	return &ByProjectKeyTaxCategoriesByIdRequestMethodGet{
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Update TaxCategory by ID
 */
func (rb *ByProjectKeyTaxCategoriesByIdRequestBuilder) Post(body TaxCategoryUpdate) *ByProjectKeyTaxCategoriesByIdRequestMethodPost {
	return &ByProjectKeyTaxCategoriesByIdRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}

/**
*	Delete TaxCategory by ID
 */
func (rb *ByProjectKeyTaxCategoriesByIdRequestBuilder) Delete() *ByProjectKeyTaxCategoriesByIdRequestMethodDelete {
	return &ByProjectKeyTaxCategoriesByIdRequestMethodDelete{
		url:    fmt.Sprintf("/%s/tax-categories/%s", rb.projectKey, rb.id),
		client: rb.client,
	}
}
