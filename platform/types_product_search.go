package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type ProductPagedSearchResponse struct {
	// Total number of results matching the query.
	Total int `json:"total"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Results for [facets](/../api/projects/product-search#facets) when requested.
	Facets []ProductSearchFacetResult `json:"facets"`
	// Search result containing the Products matching the search query.
	Results []ProductSearchResult `json:"results"`
}

type ProductSearchErrorResponse struct {
	// The HTTP status code of the response.
	StatusCode int `json:"statusCode"`
	// Describes the error.
	Message string `json:"message"`
	// The errors that caused this error response.
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductSearchErrorResponse) UnmarshalJSON(data []byte) error {
	type Alias ProductSearchErrorResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSearchErrorResponse) MarshalJSON() ([]byte, error) {
	type Alias ProductSearchErrorResponse
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["errors"] == nil {
		delete(raw, "errors")
	}

	return json.Marshal(raw)

}

func (obj ProductSearchErrorResponse) Error() string {
	if obj.Message != "" {
		return obj.Message
	}
	return "unknown ProductSearchErrorResponse: failed to parse error response"
}

type ProductSearchMatchingVariantEntry struct {
	// `id` of the [ProductVariant](ctp:api:type:ProductVariant) that matches the search query.
	ID int `json:"id"`
	// `sku` of the ProductVariant that matches the search query.
	Sku *string `json:"sku,omitempty"`
}

type ProductSearchMatchingVariants struct {
	// `true` if all Variants of the returned Product match the search query, or if search query does not specify any expression for a [Product Variant field](/../api/projects/product-search#field-levels).
	//
	// `false` if only a subset of the Product Variants match the search query.
	//
	// Is always `false` for query expressions on Product Variant fields.
	AllMatched bool `json:"allMatched"`
	// Identifiers of the Product Variants that match the search query.
	//
	// Empty if all Product Variants of the returned Product match.
	MatchedVariants []ProductSearchMatchingVariantEntry `json:"matchedVariants"`
}

/**
*	The query parameters used for [data integration with Product Projection parameters](/../api/projects/product-search#with-product-projection-parameters).
*
 */
type ProductSearchProjectionParams struct {
	// Expands a `value` of type [Reference](ctp:api:type:Reference) in a [ProductProjection](ctp:api:type:ProductProjection).
	// In case the referenced object does not exist, the API returns the non-expanded reference.
	Expand []string `json:"expand"`
	// Set to `true` to retrieve the [staged](ctp:api:type:CurrentStaged) Product Projection
	Staged *bool `json:"staged,omitempty"`
	// The currency used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	PriceCurrency *string `json:"priceCurrency,omitempty"`
	// The country used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). It can be used only *in conjunction with* the `priceCurrency` parameter.
	PriceCountry *string `json:"priceCountry,omitempty"`
	// `id` of an existing [CustomerGroup](ctp:api:type:CustomerGroup) used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). It can be used only *in conjunction with* the `priceCurrency` parameter.
	PriceCustomerGroup *string `json:"priceCustomerGroup,omitempty"`
	// IDs of existing [CustomerGroups](ctp:api:type:CustomerGroup) used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection), when using [multiple Customer Groups](/../api/customers-overview#customer-groups). It can be used only *in conjunction with* the `priceCurrency` parameter.
	PriceCustomerGroupAssignments []string `json:"priceCustomerGroupAssignments"`
	// `id` of an existing [Channel](ctp:api:type:Channel) used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). It can be used only *in conjunction with* the `priceCurrency` parameter.
	PriceChannel *string `json:"priceChannel,omitempty"`
	// Used for [locale-based projection](ctp:api:type:ProductProjectionLocales).
	LocaleProjection []string `json:"localeProjection"`
	// `key` of an existing [Store](ctp:api:type:Store).
	// If the Store has defined `languages`, `countries`, `distributionChannels`, or `supplyChannels`,
	// they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices),
	// and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
	// For Projects with active [Product Selections](/../api/projects/product-selections), the API does not take the [availability of the Product in the specified Store](/../api/projects/stores#products-available-in-store) into account.
	// [Product Tailoring](/../api/projects/product-tailoring) modifies the product information returned in API responses, but evaluating [query expressions](/../api/search-query-language#simple-expressions) is restricted to the original product information.
	StoreProjection *string `json:"storeProjection,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSearchProjectionParams) MarshalJSON() ([]byte, error) {
	type Alias ProductSearchProjectionParams
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["expand"] == nil {
		delete(raw, "expand")
	}

	if raw["priceCustomerGroupAssignments"] == nil {
		delete(raw, "priceCustomerGroupAssignments")
	}

	if raw["localeProjection"] == nil {
		delete(raw, "localeProjection")
	}

	return json.Marshal(raw)

}

type ProductSearchRequest struct {
	// The search query against [searchable Product fields](/../api/projects/product-search#searchable-product-fields).
	Query *SearchQuery `json:"query,omitempty"`
	// Controls how results to your query are [sorted](/../api/projects/product-search#sorting).
	// If not provided, the results are sorted by relevance score in descending order.
	Sort []SearchSorting `json:"sort"`
	// The maximum number of search results to be returned in one [page](/../api/projects/product-search#pagination).
	Limit *int `json:"limit,omitempty"`
	// The number of search results to be skipped in the response for [pagination](/../api/projects/product-search#pagination).
	Offset *int `json:"offset,omitempty"`
	// If `query` specifies an expression for a Product Variant field,
	// set this to `true` to get additional information for each returned Product about which Product Variants match the search query.
	// For details, see [matching variants](/../api/projects/product-search#matching-variants).
	MarkMatchingVariants *bool `json:"markMatchingVariants,omitempty"`
	// Controls data integration [with Product Projection parameters](/../api/projects/product-search#with-product-projection-parameters).
	// If not set, the result does not include the Product Projection.
	ProductProjectionParameters *ProductSearchProjectionParams `json:"productProjectionParameters,omitempty"`
	// Set this field to request [facets](/../api/projects/product-search#facets).
	Facets []ProductSearchFacetExpression `json:"facets"`
	// Specify an additional filter on the result of the `query` after the API calculated `facets`.
	// This feature assists you in implementing faceted search.
	PostFilter *SearchQuery `json:"postFilter,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSearchRequest) MarshalJSON() ([]byte, error) {
	type Alias ProductSearchRequest
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["sort"] == nil {
		delete(raw, "sort")
	}

	if raw["facets"] == nil {
		delete(raw, "facets")
	}

	return json.Marshal(raw)

}

type ProductSearchResult struct {
	// `id` of the [Product](ctp:api:type:Product) that matches the search query.
	ID string `json:"id"`
	// Information about which Product Variants match the search query.
	// Only present if `markMatchingVariants` is set to `true` in the [ProductSearchRequest](ctp:api:type:ProductSearchRequest).
	MatchingVariants *ProductSearchMatchingVariants `json:"matchingVariants,omitempty"`
	// Projected data of the Product with `id`.
	// Only present if data integration [with Product Projection parameters](/../api/projects/product-search#with-product-projection-parameters) is requested.
	ProductProjection *ProductProjection `json:"productProjection,omitempty"`
}

type ProductSearchFacetCountLevelEnum string

const (
	ProductSearchFacetCountLevelEnumProducts ProductSearchFacetCountLevelEnum = "products"
	ProductSearchFacetCountLevelEnumVariants ProductSearchFacetCountLevelEnum = "variants"
)

type ProductSearchFacetCountValue struct {
	// Name of the count facet to appear in the [ProductSearchFacetResultCount](ctp:api:type:ProductSearchFacetResultCount).
	Name string `json:"name"`
	// Whether the facet must consider only the Products resulting from the search (`query`) or all the Products (`all`).
	Scope *ProductSearchFacetScopeEnum `json:"scope,omitempty"`
	// Additional filtering expression to apply to the facet result before calculating the facet.
	Filter *SearchQuery `json:"filter,omitempty"`
	// Specify whether to count Products (`products`) or Product Variants (`variants`).
	Level *ProductSearchFacetCountLevelEnum `json:"level,omitempty"`
}

type ProductSearchFacetDistinctBucketSortBy string

const (
	ProductSearchFacetDistinctBucketSortByCount ProductSearchFacetDistinctBucketSortBy = "count"
	ProductSearchFacetDistinctBucketSortByKey   ProductSearchFacetDistinctBucketSortBy = "key"
)

type ProductSearchFacetDistinctBucketSortExpression struct {
	// Defines whether to sort by bucket count or key.
	By ProductSearchFacetDistinctBucketSortBy `json:"by"`
	// Defines the sorting order.
	Order SearchSortOrder `json:"order"`
}

type ProductSearchFacetDistinctValue struct {
	// Name of the distinct facet to appear in the [ProductSearchFacetResultBucket](ctp:api:type:ProductSearchFacetResultBucket).
	Name string `json:"name"`
	// Whether the facet must consider only the Products resulting from the search (`query`) or all the Products (`all`).
	Scope *ProductSearchFacetScopeEnum `json:"scope,omitempty"`
	// Additional filtering expression to apply to the facet result before calculating the facet.
	Filter *SearchQuery `json:"filter,omitempty"`
	// Specify whether to count Products (`products`) or Product Variants (`variants`).
	Level *ProductSearchFacetCountLevelEnum `json:"level,omitempty"`
	// The [searchable Product field](/api/projects/product-search#searchable-product-fields) to facet on.
	Field string `json:"field"`
	// Specify which bucket keys the facets results should include.
	Includes []string `json:"includes"`
	// Define how the buckets are sorted.
	Sort *ProductSearchFacetDistinctBucketSortExpression `json:"sort,omitempty"`
	// Maximum number of buckets to return.
	Limit *int `json:"limit,omitempty"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language *string `json:"language,omitempty"`
	// If the `field` is not standard, this must be the Attribute type.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
	// Default value to use if the specified field is not present on some Products.
	Missing *string `json:"missing,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSearchFacetDistinctValue) MarshalJSON() ([]byte, error) {
	type Alias ProductSearchFacetDistinctValue
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["includes"] == nil {
		delete(raw, "includes")
	}

	return json.Marshal(raw)

}

type ProductSearchFacetExpression map[string]interface{}
type ProductSearchFacetCountExpression struct {
	// Definition of the count facet.
	Count ProductSearchFacetCountValue `json:"count"`
}

type ProductSearchFacetDistinctExpression struct {
	// Definition of the distinct facet.
	Distinct ProductSearchFacetDistinctValue `json:"distinct"`
}

type ProductSearchFacetRangesExpression struct {
	// Definition of the ranges facet.
	Ranges ProductSearchFacetRangesValue `json:"ranges"`
}

/**
*	Values for `from` and `to` must be a number or [DateTime](ctp:api:type:DateTime).
 */
type ProductSearchFacetRangesFacetRange struct {
	// Starting value of the bucket (inclusive).
	From interface{} `json:"from,omitempty"`
	// Ending value of the bucket (non-inclusive).
	To interface{} `json:"to,omitempty"`
	// Key to assign the bucket.
	Key *string `json:"key,omitempty"`
}

type ProductSearchFacetRangesValue struct {
	// Name of the ranges facet to appear in the [ProductSearchFacetResultBucket](ctp:api:type:ProductSearchFacetResultBucket).
	Name string `json:"name"`
	// Whether the facet must consider only the Products resulting from the search (`query`) or all the Products (`all`).
	Scope *ProductSearchFacetScopeEnum `json:"scope,omitempty"`
	// Additional filtering expression to apply to the facet result before calculating the facet.
	Filter *SearchQuery `json:"filter,omitempty"`
	// Specify whether to count Products (`products`) or Product Variants (`variants`).
	Level *ProductSearchFacetCountLevelEnum `json:"level,omitempty"`
	// The [searchable Product field](/api/projects/product-search#searchable-product-fields) to facet on.
	Field string `json:"field"`
	// Define ranges for the facet.
	Ranges []ProductSearchFacetRangesFacetRange `json:"ranges"`
	// String value specifying linguistic and regional preferences using the [IETF language tag format](https://en.wikipedia.org/wiki/IETF_language_tag), as described in [BCP 47](https://www.rfc-editor.org/rfc/bcp/bcp47.txt). The format combines language, script, and region using hyphen-separated subtags. For example: `en`, `en-US`, `zh-Hans-SG`.
	Language *string `json:"language,omitempty"`
	// If the `field` is not standard, this must be the Attribute type.
	FieldType *SearchFieldType `json:"fieldType,omitempty"`
}

type ProductSearchFacetResult struct {
	// Name of the facet.
	Name string `json:"name"`
}

/**
*	Result of a [distinct facet](/../api/projects/product-search#distinct-facets) or a [ranges facet](/../api/projects/product-search#ranges-facets).
*
 */
type ProductSearchFacetResultBucket struct {
	// Name of the facet.
	Name string `json:"name"`
	// Contains results of the facet.
	Buckets []ProductSearchFacetResultBucketEntry `json:"buckets"`
}

type ProductSearchFacetResultBucketEntry struct {
	// Key of the bucket.
	Key string `json:"key"`
	// Number of values in the bucket.
	Count int `json:"count"`
}

/**
*	Result of a [count facet](/../api/projects/product-search#count-facets).
*
 */
type ProductSearchFacetResultCount struct {
	// Name of the facet.
	Name string `json:"name"`
	// Number of Products (or Product Variants) matching the query.
	Value int `json:"value"`
}

type ProductSearchFacetScopeEnum string

const (
	ProductSearchFacetScopeEnumAll   ProductSearchFacetScopeEnum = "all"
	ProductSearchFacetScopeEnumQuery ProductSearchFacetScopeEnum = "query"
)
