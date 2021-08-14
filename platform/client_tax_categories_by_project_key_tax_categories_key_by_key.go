// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Get TaxCategory by key
 */
func (rb *ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder) Get() *ByProjectKeyTaxCategoriesKeyByKeyRequestMethodGet {
	return &ByProjectKeyTaxCategoriesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/tax-categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Update TaxCategory by key
 */
func (rb *ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder) Post(body TaxCategoryUpdate) *ByProjectKeyTaxCategoriesKeyByKeyRequestMethodPost {
	return &ByProjectKeyTaxCategoriesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/tax-categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Delete TaxCategory by key
 */
func (rb *ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder) Delete() *ByProjectKeyTaxCategoriesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyTaxCategoriesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/tax-categories/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
