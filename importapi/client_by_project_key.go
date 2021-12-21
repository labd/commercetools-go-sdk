// Generated file, please do not change!!!
package importapi

import ()

type ByProjectKeyRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyRequestBuilder) ImportSinks() *ByProjectKeyImportSinksRequestBuilder {
	return &ByProjectKeyImportSinksRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ImportSummaries() *ByProjectKeyImportSummariesRequestBuilder {
	return &ByProjectKeyImportSummariesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ImportContainers() *ByProjectKeyImportContainersRequestBuilder {
	return &ByProjectKeyImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ImportOperations() *ByProjectKeyImportOperationsRequestBuilder {
	return &ByProjectKeyImportOperationsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Categories() *ByProjectKeyCategoriesRequestBuilder {
	return &ByProjectKeyCategoriesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Prices() *ByProjectKeyPricesRequestBuilder {
	return &ByProjectKeyPricesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Products() *ByProjectKeyProductsRequestBuilder {
	return &ByProjectKeyProductsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ProductDrafts() *ByProjectKeyProductDraftsRequestBuilder {
	return &ByProjectKeyProductDraftsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ProductTypes() *ByProjectKeyProductTypesRequestBuilder {
	return &ByProjectKeyProductTypesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ProductVariants() *ByProjectKeyProductVariantsRequestBuilder {
	return &ByProjectKeyProductVariantsRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) ProductVariantPatches() *ByProjectKeyProductVariantPatchesRequestBuilder {
	return &ByProjectKeyProductVariantPatchesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Orders() *ByProjectKeyOrdersRequestBuilder {
	return &ByProjectKeyOrdersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) OrderPatches() *ByProjectKeyOrderPatchesRequestBuilder {
	return &ByProjectKeyOrderPatchesRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyRequestBuilder) Customers() *ByProjectKeyCustomersRequestBuilder {
	return &ByProjectKeyCustomersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
