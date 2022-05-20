package ml

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	A set of ProductData for comparison. If no optional attributes are specified, all `current` ProductData are selected for comparison.
 */
type ProductSetSelector struct {
	// The project containing the project set.
	ProjectKey string `json:"projectKey"`
	// An array of Product IDs to compare. If unspecified, no Product ID filter is applied.
	ProductIds []string `json:"productIds"`
	// An array of product type IDs. Only products with product types in this array are compared. If unspecified, no product type filter is applied.
	ProductTypeIds []string `json:"productTypeIds"`
	// Specifies use of staged or current product data.
	Staged *bool `json:"staged,omitempty"`
	// Specifies use of product variants. If set to `true`, all product variants are compared, not just the master variant.
	IncludeVariants *bool `json:"includeVariants,omitempty"`
	// Maximum number of products to check (if unspecified, all products are considered). Note that the maximum number of product comparisons between two productSets is 20,000,000. This limit cannot be exceeded. If you need a higher limit, contact https://support.commercetools.com
	ProductSetLimit *int `json:"productSetLimit,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetSelector) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSelector
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["productIds"] == nil {
		delete(target, "productIds")
	}

	if target["productTypeIds"] == nil {
		delete(target, "productTypeIds")
	}

	return json.Marshal(target)
}

/**
*	Specify which ProductData attributes to use for estimating similarity and how to weigh them. An attribute's weight can be any whole positive integer, starting with 0. The larger the integer, the higher its weight.
 */
type SimilarityMeasures struct {
	// Importance of the `name` attribute in overall similarity.
	Name *int `json:"name,omitempty"`
	// Importance of the `description` attribute in overall similarity.
	Description *int `json:"description,omitempty"`
	// Importance of the `description` attribute in overall similarity.
	Attribute *int `json:"attribute,omitempty"`
	// Importance of the number of product variants in overall similarity.
	VariantCount *int `json:"variantCount,omitempty"`
	// Importance of the `price` attribute in overall similarity.
	Price *int `json:"price,omitempty"`
}

type SimilarProductSearchRequest struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit *int `json:"limit,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset *int `json:"offset,omitempty"`
	// language tag used to prioritize language for text comparisons.
	Language *string `json:"language,omitempty"`
	// The three-digit  currency code to compare prices in. When a product has multiple prices, all prices for the product are converted to the currency provided by the currency attribute and the median price is calculated for comparison. Currencies are converted using the ECB currency exchange rates at the time the request is made. Of the currency codes, only currencies with currency exchange rates provided by the ECB are supported.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// `similarityMeasures` defines the attributes taken into account to measure product similarity.
	SimilarityMeasures *SimilarityMeasures `json:"similarityMeasures,omitempty"`
	// Array of length 2 of ProductSetSelector
	ProductSetSelectors []ProductSetSelector `json:"productSetSelectors"`
	ConfidenceMin       *float64             `json:"confidenceMin,omitempty"`
	ConfidenceMax       *float64             `json:"confidenceMax,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj SimilarProductSearchRequest) MarshalJSON() ([]byte, error) {
	type Alias SimilarProductSearchRequest
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["productSetSelectors"] == nil {
		delete(target, "productSetSelectors")
	}

	return json.Marshal(target)
}

/**
*	One part of a SimilarProductPair. Refers to a specific ProductVariant.
 */
type SimilarProduct struct {
	// Reference to Product
	Product *ProductReference `json:"product,omitempty"`
	// ID of the ProductVariant that was compared.
	VariantId *int `json:"variantId,omitempty"`
	// Supplementary information about the data used for similarity estimation. This information helps you understand the estimated confidence score, but it should not be used to identify a product.
	Meta *SimilarProductMeta `json:"meta,omitempty"`
}

type SimilarProductMeta struct {
	// Localized product name used for similarity estimation.
	Name *LocalizedString `json:"name,omitempty"`
	// Localized product description used for similarity estimation.
	Description *LocalizedString `json:"description,omitempty"`
	// The product price in cents using the currency defined in SimilarProductSearchRequest If multiple prices exist, the median value is taken as a representative amount.
	Price *Money `json:"price,omitempty"`
	// Total number of variants associated with the product.
	VariantCount *int `json:"variantCount,omitempty"`
}

/**
*	A pair of SimilarProducts
 */
type SimilarProductPair struct {
	// The probability of product similarity.
	Confidence float64          `json:"confidence"`
	Products   []SimilarProduct `json:"products"`
}

type SimilarProductSearchRequestMeta struct {
	// The SimilarityMeasures used in this search.
	SimilarityMeasures SimilarityMeasures `json:"similarityMeasures"`
}

type SimilarProductsPagedQueryResult struct {
	Count int `json:"count"`
	Total int `json:"total"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int                             `json:"offset"`
	Results []SimilarProductPair            `json:"results"`
	Meta    SimilarProductSearchRequestMeta `json:"meta"`
}

/**
*	Represents a URL path to poll to get the results of an Asynchronous Request.
 */
type SimilarProductsTaskStatus struct {
	State TaskStatusEnum `json:"state"`
	// The expiry date of the result. You cannot access the result after the expiry date. Default: 1 day after the result first becomes available. This is only available when the TaskStatus state is SUCCESS.
	Expires *time.Time `json:"expires,omitempty"`
	// The response to an asynchronous request. The type depends on the request initiated. Only populated when the status is `SUCCESS`.
	Result SimilarProductsPagedQueryResult `json:"result"`
}
