// Generated file, please do not change!!!
package ml

import (
	"time"
)

type AttributeCount struct {
	// Number of attributes defined in the product type.
	ProductTypeAttributes int `json:"productTypeAttributes"`
	// Number of attributes defined in the variant.
	VariantAttributes int `json:"variantAttributes"`
	// Number of attributes missing values in the variant.
	MissingAttributeValues int `json:"missingAttributeValues"`
}

type AttributeCoverage struct {
	// The percentage of attributes from the product type defined in the product variant. A value of `1.0` indicates a product variant contains all attributes defined in the product type.
	Names float64 `json:"names"`
	// Represents the percentage of attributes in the product variant that contain values.
	Values float64 `json:"values"`
}

type MissingAttributesDetails struct {
	// Number of products scanned.
	Total int `json:"total"`
	// Number of products missing attribute names.
	MissingAttributeNames int `json:"missingAttributeNames"`
	// Number of products missing attribute values.
	MissingAttributeValues int `json:"missingAttributeValues"`
}

type MissingAttributes struct {
	Product     ProductReference     `json:"product"`
	ProductType ProductTypeReference `json:"productType"`
	// ID of a ProductVariant.
	VariantId int `json:"variantId"`
	// The names of the attributes of the product type that the variant is missing, sorted by attribute importance in descending order.
	MissingAttributeValues []string `json:"missingAttributeValues"`
	// The names of the attributes of the product type that the variant is missing, sorted by attribute importance in descending order.
	MissingAttributeNames []string           `json:"missingAttributeNames,omitempty"`
	AttributeCount        *AttributeCount    `json:"attributeCount,omitempty"`
	AttributeCoverage     *AttributeCoverage `json:"attributeCoverage,omitempty"`
}

type MissingAttributesMeta struct {
	ProductLevel MissingAttributesDetails `json:"productLevel"`
	VariantLevel MissingAttributesDetails `json:"variantLevel"`
	// The IDs of the product types containing the requested `attributeName`.
	ProductTypeIds []string `json:"productTypeIds,omitempty"`
}

type MissingAttributesSearchRequest struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	// If true, searches data from staged products in addition to published products.
	Staged *bool `json:"staged,omitempty"`
	// Maximum number of products to scan.
	ProductSetLimit *int `json:"productSetLimit,omitempty"`
	// If true, searches all product variants. If false, only searches master variants.
	IncludeVariants *bool `json:"includeVariants,omitempty"`
	// Minimum attribute coverage of variants to display, applied to both coverage types.
	CoverageMin *float64 `json:"coverageMin,omitempty"`
	// Maximum attribute coverage of variants to display, applied to both coverage types.
	CoverageMax *float64 `json:"coverageMax,omitempty"`
	// Default value: `coverageAttributeValues` - Allowed values: [`coverageAttributeValues`, `coverageAttributeNames`]
	// `coverageAttributeValues` shows the product variants with the most missing attribute values first and `coverageAttributeNames` the ones with the most missing attribute names.
	SortBy *string `json:"sortBy,omitempty"`
	// If true, the `missingAttributeNames` will be included in the results.
	ShowMissingAttributeNames *bool `json:"showMissingAttributeNames,omitempty"`
	// Filters results by the provided Product IDs.
	// Cannot be applied in combination with any other filter.
	ProductIds []string `json:"productIds,omitempty"`
	// Filters results by the provided product type IDs.
	// Cannot be applied in combination with any other filter.
	ProductTypeIds []string `json:"productTypeIds,omitempty"`
	// Filters results by the provided attribute name. If provided,  products are only checked for this attribute. Therefore, only products of product types which define the attribute name are considered. These product type IDs
	// are then listed in `MissingAttributesMeta`. The  `attributeCount` and `attributeCoverage` fields are not part of the response when using this filter. Cannot be applied in combination with any other filter.
	AttributeName *string `json:"attributeName,omitempty"`
}

type MissingAttributesPagedQueryResult struct {
	Count   int                   `json:"count"`
	Total   int                   `json:"total"`
	Offset  int                   `json:"offset"`
	Results []MissingAttributes   `json:"results"`
	Meta    MissingAttributesMeta `json:"meta"`
}

/**
*	Represents a URL path to poll to get the results of an Asynchronous Request.
 */
type MissingDataTaskStatus struct {
	State TaskStatusEnum `json:"state"`
	// The expiry date of the result. You cannot access the result after the expiry date. Default: 1 day after the result first becomes available. This is only available when the TaskStatus state is SUCCESS.
	Expires time.Time `json:"expires"`
	// The response to an asynchronous request. The type depends on the request initiated. Only populated when the status is `SUCCESS`.
	Result MissingAttributesPagedQueryResult `json:"result"`
}

type MissingImages struct {
	Product ProductReference `json:"product"`
	// ID of the variant
	VariantId int `json:"variantId"`
	// Number of images the variant contains.
	ImageCount int `json:"imageCount"`
}

type MissingImagesCount struct {
	MissingImages int `json:"missingImages"`
	// Number of products scanned.
	Total int `json:"total"`
}

type MissingImagesProductLevel struct {
	// Number of products missing images.
	MissingImages int `json:"missingImages"`
	// Number of products scanned.
	Total int `json:"total"`
}

type MissingImagesVariantLevel struct {
	// Number of product variants missing images.
	MissingImages int `json:"missingImages"`
	// Number of products scanned.
	Total int `json:"total"`
}

type MissingImagesMeta struct {
	ProductLevel MissingImagesProductLevel `json:"productLevel"`
	VariantLevel MissingImagesVariantLevel `json:"variantLevel"`
	// The minimum number of images a product variant must have. Anything below this value is considered a product variant with missing images.
	Threshold int `json:"threshold"`
}

type MissingImagesSearchRequest struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	// If true, searches data from staged products in addition to published products.
	Staged *bool `json:"staged,omitempty"`
	// Maximum number of products to scan.
	ProductSetLimit *int `json:"productSetLimit,omitempty"`
	// If true, searches all product variants. If false, only searches master variants.
	IncludeVariants *bool `json:"includeVariants,omitempty"`
	// If true, uses the median number of images per product variant as a threshold value.
	AutoThreshold *bool `json:"autoThreshold,omitempty"`
	// The minimum number of images a product variant must have. Anything below this value is considered a product variant with missing images.
	Threshold *int `json:"threshold,omitempty"`
	// Filters results by the provided Product IDs. Cannot be applied in combination with any other filter.
	ProductIds []string `json:"productIds,omitempty"`
	// Filters results by the provided product type IDs. It cannot be applied in combination with any other filter.
	ProductTypeIds []string `json:"productTypeIds,omitempty"`
}

type MissingImagesPagedQueryResult struct {
	Count   int               `json:"count"`
	Total   int               `json:"total"`
	Offset  int               `json:"offset"`
	Results []MissingImages   `json:"results"`
	Meta    MissingImagesMeta `json:"meta"`
}

/**
*	Represents a URL path to poll to get the results of an Asynchronous Request.
 */
type MissingImagesTaskStatus struct {
	State TaskStatusEnum `json:"state"`
	// The expiry date of the result. You cannot access the result after the expiry date. Default: 1 day after the result first becomes available. This is only available when the TaskStatus state is SUCCESS.
	Expires time.Time `json:"expires"`
	// The response to an asynchronous request. The type depends on the request initiated. Only populated when the status is `SUCCESS`.
	Result MissingImagesPagedQueryResult `json:"result"`
}

type MissingPrices struct {
	Product ProductReference `json:"product"`
	// Id of the `ProductVariant`.
	VariantId int `json:"variantId"`
}

type MissingPricesProductCount struct {
	Total         int `json:"total"`
	MissingPrices int `json:"missingPrices"`
}

type MissingPricesProductLevel struct {
	// Number of products scanned.
	Total int `json:"total"`
	// Number of products missing prices.
	MissingPrices int `json:"missingPrices"`
}

type MissingPricesVariantLevel struct {
	// Number of product variants scanned.
	Total int `json:"total"`
	// Number of product variants missing prices.
	MissingPrices int `json:"missingPrices"`
}

type MissingPricesMeta struct {
	ProductLevel MissingPricesProductLevel `json:"productLevel"`
	VariantLevel MissingPricesVariantLevel `json:"variantLevel"`
}

type MissingPricesSearchRequest struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
	// If true, searches data from staged products in addition to published products.
	Staged *bool `json:"staged,omitempty"`
	// Maximum number of products to scan.
	ProductSetLimit *int `json:"productSetLimit,omitempty"`
	// If true, searches all product variants. If false, only searches master variants.
	IncludeVariants *bool `json:"includeVariants,omitempty"`
	// If used, only checks if a product variant has a price in the provided currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// If true, checks if there are prices for the specified date range and time.
	CheckDate *bool `json:"checkDate,omitempty"`
	// Starting date of the range to check. If no value is given, checks prices valid at the time the search is initiated.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Ending date of the range to check. If no value is given, it is equal to `validFrom`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Filters results by the provided Product IDs. Cannot be applied in combination with the `productTypeIds` filter.
	ProductIds []string `json:"productIds,omitempty"`
	// Filters results by the provided product type IDs. Cannot be applied in combination with the `productIds` filter.
	ProductTypeIds []string `json:"productTypeIds,omitempty"`
}

type MissingPricesPagedQueryResult struct {
	Count   int               `json:"count"`
	Total   int               `json:"total"`
	Offset  int               `json:"offset"`
	Results []MissingPrices   `json:"results"`
	Meta    MissingPricesMeta `json:"meta"`
}

/**
*	Represents a URL path to poll to get the results of an Asynchronous Request.
 */
type MissingPricesTaskStatus struct {
	State TaskStatusEnum `json:"state"`
	// The expiry date of the result. You cannot access the result after the expiry date. Default: 1 day after the result first becomes available. This is only available when the TaskStatus state is SUCCESS.
	Expires time.Time `json:"expires"`
	// The response to an asynchronous request. The type depends on the request initiated. Only populated when the status is `SUCCESS`.
	Result MissingPricesPagedQueryResult `json:"result"`
}
