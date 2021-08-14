// Generated file, please do not change!!!
package platform

import (
	"fmt"
)

type ByProjectKeyTaxCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTaxCategoriesRequestBuilder) WithKey(key string) *ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder {
	return &ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) WithId(id string) *ByProjectKeyTaxCategoriesByIdRequestBuilder {
	return &ByProjectKeyTaxCategoriesByIdRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}

/**
*	Query tax-categories
 */
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) Get() *ByProjectKeyTaxCategoriesRequestMethodGet {
	return &ByProjectKeyTaxCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/tax-categories", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Create TaxCategory
 */
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) Post(body TaxCategoryDraft) *ByProjectKeyTaxCategoriesRequestMethodPost {
	return &ByProjectKeyTaxCategoriesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/tax-categories", rb.projectKey),
		client: rb.client,
	}
}
