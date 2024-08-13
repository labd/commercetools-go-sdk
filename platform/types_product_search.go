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
	// Unique identifier of the variant.
	ID int `json:"id"`
	// SKU of the matching variant.
	Sku *string `json:"sku,omitempty"`
}

type ProductSearchMatchingVariants struct {
	// Whether the search criteria definitely matches all Variants of the returned Product, like for Product-level fields. Is always `false` for search expressions on Variant-level fields.
	AllMatched bool `json:"allMatched"`
	// The variants matching the search criteria or empty if all matched.
	MatchedVariants []ProductSearchMatchingVariantEntry `json:"matchedVariants"`
}

type ProductSearchProjectionParams struct {
	// Expands a `value` of type [Reference](ctp:api:type:Reference) in a [ProductProjection](ctp:api:type:ProductProjection).
	// In case the referenced object does not exist, the API returns the non-expanded reference.
	Expand []string `json:"expand"`
	// Set to `true` to retrieve the [staged](ctp:api:type:CurrentStaged) Product Projection
	Staged *bool `json:"staged,omitempty"`
	// The currency used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	PriceCurrency *string `json:"priceCurrency,omitempty"`
	// The country used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). Can only be used **in conjunction with** the `priceCurrency` parameter.
	PriceCountry *string `json:"priceCountry,omitempty"`
	// `id` of an existing [CustomerGroup](ctp:api:type:CustomerGroup) used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). Can only be used **in conjunction with** the `priceCurrency` parameter.
	PriceCustomerGroup *string `json:"priceCustomerGroup,omitempty"`
	// `id` of an existing [Channel](ctp:api:type:Channel) used for [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection). Can only be used **in conjunction with** the `priceCurrency` parameter.
	PriceChannel *string `json:"priceChannel,omitempty"`
	// Used for [locale-based projection](ctp:api:type:ProductProjectionLocales).
	LocaleProjection []string `json:"localeProjection"`
	// `key` of an existing [Store](ctp:api:type:Store).
	// If the Store has defined some languages, countries, distribution or supply Channels,
	// they are used for projections based on [locale](ctp:api:type:ProductProjectionLocales), [price](ctp:api:type:ProductProjectionPrices),
	// and [inventory](ctp:api:type:ProductProjectionInventoryEntries).
	// If the Store has defined [Product Selections](ctp:api:type:ProductSelection) or [Product Tailoring](ctp:api:type:ProductTailoring), they have no effect on the results of this query.
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

	if raw["localeProjection"] == nil {
		delete(raw, "localeProjection")
	}

	return json.Marshal(raw)

}

type ProductSearchRequest struct {
	// The search query against [searchable Product fields](/../api/projects/product-search#searchable-product-fields).
	Query *SearchQuery `json:"query,omitempty"`
	// Controls how results to your query are sorted. If not provided, the results are sorted by relevance in descending order.
	Sort []SearchSorting `json:"sort"`
	// The maximum number of search results to be returned.
	Limit *int `json:"limit,omitempty"`
	// The number of search results to be skipped in the response for pagination.
	Offset *int `json:"offset,omitempty"`
	// The search can return Products where not all Product Variants match the search criteria. If `true`, the response will include a field called `matchingVariants` that contains the `sku` of Product Variants that match the search query. If the query does not specify any variant-level criteria, `matchingVariants` will be null signifying that all Product Variants are a match.
	MarkMatchingVariants *bool `json:"markMatchingVariants,omitempty"`
	// Set this field to `{}` to get the [ProductProjection](ctp:api:type:ProductProjection) included in the [ProductSearchResult](ctp:api:type:ProductSearchResult).
	// Include query parameters for controlling [Reference Expansion](/../api/general-concepts#reference-expansion) or [projections](/../api/projects/productProjections#projection-dimensions) according to your needs.
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
	// Unique identifier of the Product.
	ID string `json:"id"`
	// Contains Product Projection data for Products matching the `projection` field in the Search Products request.
	ProductProjection *ProductProjection `json:"productProjection,omitempty"`
	// Describes the variants that matched the search criteria.
	MatchingVariants *ProductSearchMatchingVariants `json:"matchingVariants,omitempty"`
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
	// Additional filtering expression to apply to the search result before calculating the facet.
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
	// Additional filtering expression to apply to the search result before calculating the facet.
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
	// Additional filtering expression to apply to the search result before calculating the facet.
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

type ProductSearchFacetScope string

const (
	ProductSearchFacetScopeAll   ProductSearchFacetScope = "all"
	ProductSearchFacetScopeQuery ProductSearchFacetScope = "query"
)

type ProductSearchFacetScopeEnum string

const (
	ProductSearchFacetScopeEnumAll   ProductSearchFacetScopeEnum = "all"
	ProductSearchFacetScopeEnumQuery ProductSearchFacetScopeEnum = "query"
)
