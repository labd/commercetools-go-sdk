package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Attribute struct {
	// Name of the Attribute.
	Name string `json:"name"`
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, use the list of values of all Attributes of the nested Product as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	Value interface{} `json:"value"`
}

/**
*	JSON object where the key is a [Category](ctp:api:type:Category) `id` and the value is an order hint.
*	Allows controlling the order of Products and how they appear in Categories. Products with no order hint have an order score below `0`. Order hints are non-unique.
*	If a subset of Products have the same value for order hint in a specific category, the behavior is undetermined.
 */
type CategoryOrderHints map[string]string
type FacetRange struct {
	From         float64 `json:"from"`
	FromStr      string  `json:"fromStr"`
	To           float64 `json:"to"`
	ToStr        string  `json:"toStr"`
	Count        int     `json:"count"`
	ProductCount *int    `json:"productCount,omitempty"`
	Total        float64 `json:"total"`
	Min          float64 `json:"min"`
	Max          float64 `json:"max"`
	Mean         float64 `json:"mean"`
}

type FacetResult interface{}

func mapDiscriminatorFacetResult(input interface{}) (FacetResult, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "filter":
		obj := FilteredFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "range":
		obj := RangeFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "terms":
		obj := TermFacetResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type FacetResults map[string]FacetResult
type FacetTerm struct {
	Term         interface{} `json:"term"`
	Count        int         `json:"count"`
	ProductCount *int        `json:"productCount,omitempty"`
}

type FacetTypes string

const (
	FacetTypesTerms  FacetTypes = "terms"
	FacetTypesRange  FacetTypes = "range"
	FacetTypesFilter FacetTypes = "filter"
)

type FilteredFacetResult struct {
	Count        int  `json:"count"`
	ProductCount *int `json:"productCount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj FilteredFacetResult) MarshalJSON() ([]byte, error) {
	type Alias FilteredFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "filter", Alias: (*Alias)(&obj)})
}

/**
*	An abstract sellable good with a set of Attributes defined by a Product Type.
*	Products themselves are not sellable. Instead, they act as a parent structure for Product Variants.
*	Each Product must have at least one Product Variant, which is called the Master Variant.
*	A single Product representation contains the _current_ and the _staged_ representation of its product data.
*
 */
type Product struct {
	// Unique identifier of the Product.
	ID string `json:"id"`
	// Current version of the Product.
	Version int `json:"version"`
	// Date and time (UTC) the Product was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Product was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Product.
	//
	// This is different from the `key` of a [ProductVariant](ctp:api:type:ProductVariant).
	Key *string `json:"key,omitempty"`
	// The Product Type defining the Attributes of the Product. Cannot be changed.
	ProductType ProductTypeReference `json:"productType"`
	// Contains the current and the staged representation of the product information.
	MasterData ProductCatalogData `json:"masterData"`
	// The [TaxCategory](ctp:api:type:TaxCategory) of the Product.
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// [State](ctp:api:type:State) of the Product.
	State *StateReference `json:"state,omitempty"`
	// Review statistics of the Product.
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	// Type of Price to be used when looking up a price for the Product.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

/**
*	Contains the `current` and `staged` [ProductData](ctp:api:type:ProductData).
*
 */
type ProductCatalogData struct {
	// `true` if the Product is published.
	Published bool `json:"published"`
	// Current (published) data of the Product.
	Current ProductData `json:"current"`
	// Staged (unpublished) data of the Product.
	Staged ProductData `json:"staged"`
	// `true` if the `staged` data is different from the `current` data.
	HasStagedChanges bool `json:"hasStagedChanges"`
}

/**
*	Contains all the data of a Product and its Product Variants.
*
 */
type ProductData struct {
	// Name of the Product.
	Name LocalizedString `json:"name"`
	// [Categories](ctp:api:type:Category) assigned to the Product.
	Categories []CategoryReference `json:"categories"`
	// Numerical values to allow ordering of Products within a specified Category.
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	// Description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// User-defined identifier used in a deep-link URL for the Product.
	// Must be unique across a Project, but can be the same for Products in different [Locales](ctp:api:type:Locale).
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug LocalizedString `json:"slug"`
	// Title of the Product displayed in search results.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Description of the Product displayed in search results below the meta title.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Keywords that give additional information about the Product to search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// The Master Variant of the Product.
	MasterVariant ProductVariant `json:"masterVariant"`
	// Additional Product Variants.
	Variants []ProductVariant `json:"variants"`
	// Used by [Product Suggestions](ctp:api:type:ProductSuggestions), but is also considered for a full text search.
	SearchKeywords SearchKeywords `json:"searchKeywords"`
}

type ProductDraft struct {
	// The Product Type defining the Attributes for the Product. Cannot be changed later.
	ProductType ProductTypeResourceIdentifier `json:"productType"`
	// Name of the Product.
	Name LocalizedString `json:"name"`
	// User-defined identifier used in a deep-link URL for the Product.
	// It must be unique across a Project, but a Product can have the same slug in different [Locales](ctp:api:type:Locale).
	// It must match the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug LocalizedString `json:"slug"`
	// User-defined unique identifier for the Product.
	Key *string `json:"key,omitempty"`
	// Description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Categories assigned to the Product.
	Categories []CategoryResourceIdentifier `json:"categories"`
	// Numerical values to allow ordering of Products within a specified Category.
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	// Title of the Product displayed in search results.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Description of the Product displayed in search results.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Keywords that give additional information about the Product to search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// The Product Variant to be the Master Variant for the Product. Required if `variants` are provided also.
	MasterVariant *ProductVariantDraft `json:"masterVariant,omitempty"`
	// The additional Product Variants for the Product.
	Variants []ProductVariantDraft `json:"variants"`
	// The Tax Category to be assigned to the Product.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
	// Used by [Product Suggestions](ctp:api:type:ProductSuggestions), but is also considered for a [full text search](/projects/products-search#full-text-search).
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// State to be assigned to the Product.
	State *StateResourceIdentifier `json:"state,omitempty"`
	// If `true`, the Product is published immediately to the current projection.
	Publish *bool `json:"publish,omitempty"`
	// Specifies the type of prices used when looking up a price for the Product.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDraft
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

	if raw["categories"] == nil {
		delete(raw, "categories")
	}

	if raw["variants"] == nil {
		delete(raw, "variants")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [Product](ctp:api:type:Product).
*
 */
type ProductPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](ctp:api:type:QueryPredicate), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Products](ctp:api:type:Product) matching the query.
	Results []Product `json:"results"`
}

/**
*	This mode determines the type of Prices used for [Product Price Selection](ctp:api:type:ProductPriceSelection) as well as for [LineItem Price selection](ctp:api:type:CartLineItemPriceSelection).
*
 */
type ProductPriceModeEnum string

const (
	ProductPriceModeEnumEmbedded   ProductPriceModeEnum = "Embedded"
	ProductPriceModeEnumStandalone ProductPriceModeEnum = "Standalone"
)

type ProductProjection struct {
	// Unique identifier of the [Product](ctp:api:type:Product).
	ID string `json:"id"`
	// Current version of the [Product](ctp:api:type:Product).
	Version int `json:"version"`
	// Date and time (UTC) the ProductProjection was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductProjection was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the [Product](ctp:api:type:Product).
	Key *string `json:"key,omitempty"`
	// The [ProductType](ctp:api:type:ProductType) defining the Attributes of the [Product](ctp:api:type:Product).
	ProductType ProductTypeReference `json:"productType"`
	// Name of the [Product](ctp:api:type:Product).
	Name LocalizedString `json:"name"`
	// Description of the [Product](ctp:api:type:Product).
	Description *LocalizedString `json:"description,omitempty"`
	// User-defined identifier used in a deep-link URL for the [Product](ctp:api:type:Product).
	// Must be unique across a Project, but can be the same for Products in different locales.
	// Matches the pattern `[a-zA-Z0-9_\-]{2,256}`.
	// For [good performance](/../api/predicates/query#performance-considerations), indexes are provided for the first 15 `languages` set in the [Project](ctp:api:type:Project).
	Slug LocalizedString `json:"slug"`
	// [Categories](ctp:api:type:Category) assigned to the [Product](ctp:api:type:Product).
	Categories []CategoryReference `json:"categories"`
	// Order of [Product](ctp:api:type:Product) in [Categories](ctp:api:type:Category).
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	// Title of the [Product](ctp:api:type:Product) displayed in search results.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Description of the [Product](ctp:api:type:Product) displayed in search results below the meta title.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Keywords that give additional information about the [Product](ctp:api:type:Product) to search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// Used by [Product Suggestions](/../api/projects/products-suggestions), but is also considered for a [full text search](ctp:api:type:FullTextSearch).
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// `true` if the staged data is different from the current data.
	HasStagedChanges *bool `json:"hasStagedChanges,omitempty"`
	// `true` if the [Product](ctp:api:type:Product) is [published](ctp:api:type:CurrentStaged).
	Published *bool `json:"published,omitempty"`
	// The Master Variant of the [Product](ctp:api:type:Product).
	MasterVariant ProductVariant `json:"masterVariant"`
	// Additional Product Variants.
	Variants []ProductVariant `json:"variants"`
	// The [TaxCategory](ctp:api:type:TaxCategory) of the [Product](ctp:api:type:Product).
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
	// [State](ctp:api:type:State) of the [Product](ctp:api:type:Product).
	State *StateReference `json:"state,omitempty"`
	// Review statistics of the [Product](ctp:api:type:Product).
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	// Indicates whether the Prices of the Product Projection are [embedded](ctp:api:type:Price) or [standalone](ctp:api:type:StandalonePrice). [Projecting Prices](#prices) only works with `Embedded`, there is currently no support for `Standalone`.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

type ProductProjectionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [ProductProjections](ctp:api:type:ProductProjection) matching the query.
	Results []ProductProjection `json:"results"`
}

type ProductProjectionPagedSearchResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int                 `json:"offset"`
	Results []ProductProjection `json:"results"`
	Facets  FacetResults        `json:"facets"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Product](ctp:api:type:Product).
*
 */
type ProductReference struct {
	// Unique identifier of the referenced [Product](ctp:api:type:Product).
	ID string `json:"id"`
	// Contains the representation of the expanded Product. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Products.
	Obj *Product `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Product](ctp:api:type:Product). Either `id` or `key` is required.
*
 */
type ProductResourceIdentifier struct {
	// Unique identifier of the referenced [Product](ctp:api:type:Product).
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Product](ctp:api:type:Product).
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product", Alias: (*Alias)(&obj)})
}

type ProductUpdate struct {
	// Expected version of the Product on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Product.
	Actions []ProductUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ProductUpdateAction interface{}

func mapDiscriminatorProductUpdateAction(input interface{}) (ProductUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addAsset":
		obj := ProductAddAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addExternalImage":
		obj := ProductAddExternalImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addPrice":
		obj := ProductAddPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addToCategory":
		obj := ProductAddToCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addVariant":
		obj := ProductAddVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetName":
		obj := ProductChangeAssetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetOrder":
		obj := ProductChangeAssetOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeMasterVariant":
		obj := ProductChangeMasterVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePrice":
		obj := ProductChangePriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSlug":
		obj := ProductChangeSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "legacySetSku":
		obj := ProductLegacySetSkuAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "moveImageToPosition":
		obj := ProductMoveImageToPositionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "publish":
		obj := ProductPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAsset":
		obj := ProductRemoveAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeFromCategory":
		obj := ProductRemoveFromCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeImage":
		obj := ProductRemoveImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePrice":
		obj := ProductRemovePriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeVariant":
		obj := ProductRemoveVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "revertStagedChanges":
		obj := ProductRevertStagedChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "revertStagedVariantChanges":
		obj := ProductRevertStagedVariantChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomField":
		obj := ProductSetAssetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomType":
		obj := ProductSetAssetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetDescription":
		obj := ProductSetAssetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetKey":
		obj := ProductSetAssetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetSources":
		obj := ProductSetAssetSourcesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetTags":
		obj := ProductSetAssetTagsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttribute":
		obj := ProductSetAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttributeInAllVariants":
		obj := ProductSetAttributeInAllVariantsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCategoryOrderHint":
		obj := ProductSetCategoryOrderHintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ProductSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDiscountedPrice":
		obj := ProductSetDiscountedPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImageLabel":
		obj := ProductSetImageLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaDescription":
		obj := ProductSetMetaDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaKeywords":
		obj := ProductSetMetaKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaTitle":
		obj := ProductSetMetaTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPriceMode":
		obj := ProductSetPriceModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPrices":
		obj := ProductSetPricesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductPriceCustomField":
		obj := ProductSetProductPriceCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductPriceCustomType":
		obj := ProductSetProductPriceCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductVariantKey":
		obj := ProductSetProductVariantKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSearchKeywords":
		obj := ProductSetSearchKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSku":
		obj := ProductSetSkuAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTaxCategory":
		obj := ProductSetTaxCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := ProductTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unpublish":
		obj := ProductUnpublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	A concrete sellable good for which inventory can be tracked. Product Variants are generally mapped to specific SKUs.
*
 */
type ProductVariant struct {
	// A unique, sequential identifier of the Product Variant within the Product.
	ID int `json:"id"`
	// User-defined unique SKU of the Product Variant.
	Sku *string `json:"sku,omitempty"`
	// User-defined unique identifier of the ProductVariant.
	//
	// This is different from [Product](ctp:api:type:Product) `key`.
	Key *string `json:"key,omitempty"`
	// The Embedded Prices of the Product Variant.
	// Cannot contain two Prices of the same Price scope (with same currency, country, Customer Group, Channel, `validFrom` and `validUntil`).
	Prices []Price `json:"prices"`
	// Attributes of the Product Variant.
	Attributes []Attribute `json:"attributes"`
	// Only available when [Price selection](#price-selection) is used.
	// Cannot be used in a [Query Predicate](ctp:api:type:QueryPredicate).
	Price *Price `json:"price,omitempty"`
	// Images of the Product Variant.
	Images []Image `json:"images"`
	// Media assets of the Product Variant.
	Assets []Asset `json:"assets"`
	// Set if the Product Variant is tracked by [Inventory](ctp:api:type:InventoryEntry).
	// Can be used as an optimization to reduce calls to the Inventory service.
	// May not contain the latest Inventory State (it is [eventually consistent](/general-concepts#eventual-consistency)).
	Availability *ProductVariantAvailability `json:"availability,omitempty"`
	// `true` if the Product Variant matches the search query.
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearch) request.
	IsMatchingVariant *bool `json:"isMatchingVariant,omitempty"`
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearch) request
	// with [price selection](ctp:api:type:ProductPriceSelection).
	// Can be used to sort, [filter](ctp:api:type:ProductProjectionSearchFilterScopedPrice), and facet.
	ScopedPrice *ScopedPrice `json:"scopedPrice,omitempty"`
	// Only available in response to a [Product Projection Search](ctp:api:type:ProductProjectionSearchFilterScopedPrice) request
	// with [price selection](ctp:api:type:ProductPriceSelection).
	ScopedPriceDiscounted *bool `json:"scopedPriceDiscounted,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariant) MarshalJSON() ([]byte, error) {
	type Alias ProductVariant
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

	if raw["prices"] == nil {
		delete(raw, "prices")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}

/**
*	The [InventoryEntry](ctp:api:type:InventoryEntry) information of the Product Variant. If there is a supply [Channel](ctp:api:type:Channel) for the InventoryEntry, then `channels` is returned. If not, then `isOnStock`, `restockableInDays`, and `quantityOnStock` are returned.
*
 */
type ProductVariantAvailability struct {
	// For each [InventoryEntry](ctp:api:type:InventoryEntry) with a supply Channel, an entry is added to `channels`.
	Channels *ProductVariantChannelAvailabilityMap `json:"channels,omitempty"`
	// Indicates whether a Product Variant is in stock.
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock a Product Variant once it is out of stock.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of the Product Variant that are in stock.
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
}

type ProductVariantChannelAvailability struct {
	// Indicates whether a Product Variant is in stock in a specified [Channel](ctp:api:type:Channel).
	IsOnStock *bool `json:"isOnStock,omitempty"`
	// Number of days to restock a Product Variant once it is out of stock in a specified [Channel](ctp:api:type:Channel).
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Number of items of this Product Variant that are in stock in a specified [Channel](ctp:api:type:Channel).
	AvailableQuantity *int `json:"availableQuantity,omitempty"`
	// Unique identifier of the [InventoryEntry](ctp:api:type:InventoryEntry).
	ID string `json:"id"`
	// Current version of the [InventoryEntry](ctp:api:type:InventoryEntry).
	Version int `json:"version"`
}

/**
*	JSON object where the key is a supply [Channel](ctp:api:type:Channel) `id` and the value is the [ProductVariantChannelAvailability](ctp:api:type:ProductVariantChannelAvailability) of the [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type ProductVariantChannelAvailabilityMap map[string]ProductVariantChannelAvailability

/**
*	Creates a Product Variant when included in the `masterVariant` and `variants` fields of the [ProductDraft](ctp:api:type:ProductDraft).
*
 */
type ProductVariantDraft struct {
	// User-defined unique SKU of the Product Variant.
	Sku *string `json:"sku,omitempty"`
	// User-defined unique identifier for the ProductVariant.
	Key *string `json:"key,omitempty"`
	// The Embedded Prices for the Product Variant.
	// Each Price must have its unique Price scope (with same currency, country, Customer Group, Channel, `validFrom` and `validUntil`).
	Prices []PriceDraft `json:"prices"`
	// Attributes according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Attributes []Attribute `json:"attributes"`
	// Images for the Product Variant.
	Images []Image `json:"images"`
	// Media assets for the Product Variant.
	Assets []AssetDraft `json:"assets"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDraft
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

	if raw["prices"] == nil {
		delete(raw, "prices")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}

type RangeFacetResult struct {
	Ranges []FacetRange `json:"ranges"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RangeFacetResult) MarshalJSON() ([]byte, error) {
	type Alias RangeFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "range", Alias: (*Alias)(&obj)})
}

type SearchKeyword struct {
	// Text to return in the result of a [suggest query](ctp:api:type:ProductSuggestionsSuggestQuery).
	Text string `json:"text"`
	// If no tokenizer is defined, the `text` is used as a single token.
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *SearchKeyword) UnmarshalJSON(data []byte) error {
	type Alias SearchKeyword
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SuggestTokenizer != nil {
		var err error
		obj.SuggestTokenizer, err = mapDiscriminatorSuggestTokenizer(obj.SuggestTokenizer)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Search keywords are JSON objects primarily used by [Product Suggestions](ctp:api:type:ProductSuggestions), but are also considered for a full text search. The keys are of type [Locale](ctp:api:type:Locale), and the values are an array of [SearchKeyword](ctp:api:type:SearchKeyword).
*
 */
type SearchKeywords map[string][]SearchKeyword
type SuggestTokenizer interface{}

func mapDiscriminatorSuggestTokenizer(input interface{}) (SuggestTokenizer, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "custom":
		obj := CustomTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "whitespace":
		obj := WhitespaceTokenizer{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Define arbitrary tokens that are used to match the input.
*
 */
type CustomTokenizer struct {
	// Contains custom tokens.
	Inputs []string `json:"inputs"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CustomTokenizer) MarshalJSON() ([]byte, error) {
	type Alias CustomTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "custom", Alias: (*Alias)(&obj)})
}

type Suggestion struct {
	// The suggested text.
	Text string `json:"text"`
}

type SuggestionResult map[string][]Suggestion
type TermFacetResult struct {
	DataType TermFacetResultType `json:"dataType"`
	Missing  int                 `json:"missing"`
	Total    int                 `json:"total"`
	Other    int                 `json:"other"`
	Terms    []FacetTerm         `json:"terms"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TermFacetResult) MarshalJSON() ([]byte, error) {
	type Alias TermFacetResult
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "terms", Alias: (*Alias)(&obj)})
}

type TermFacetResultType string

const (
	TermFacetResultTypeText     TermFacetResultType = "text"
	TermFacetResultTypeDate     TermFacetResultType = "date"
	TermFacetResultTypeTime     TermFacetResultType = "time"
	TermFacetResultTypeDatetime TermFacetResultType = "datetime"
	TermFacetResultTypeBoolean  TermFacetResultType = "boolean"
	TermFacetResultTypeNumber   TermFacetResultType = "number"
)

/**
*	Creates tokens by splitting the `text` field in [SearchKeyword](ctp:api:type:SearchKeyword) by whitespaces.
*
 */
type WhitespaceTokenizer struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type Alias WhitespaceTokenizer
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "whitespace", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductAddAssetAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged `assets` are updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// Value to append.
	Asset AssetDraft `json:"asset"`
	// Position in `assets` where the Asset should be put. When specified, the value must be between `0` and the total number of Assets minus `1`.
	Position *int `json:"position,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. Produces the [ProductImageAdded](/projects/messages#product-image-added) Message.
*
 */
type ProductAddExternalImageAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Value to add to `images`.
	Image Image `json:"image"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

/**
*	Adds the given Price to the `prices` array of the [ProductVariant](ctp:api:type:ProductVariant).
*	Either `variantId` or `sku` is required.
*
 */
type ProductAddPriceAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Embedded Price to add to the Product Variant.
	Price PriceDraft `json:"price"`
	// If `true`, only the staged `prices` is updated. If `false`, both the current and staged `prices` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPrice", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ProductAddedToCategory](/projects/messages#product-added-to-category) Message.
 */
type ProductAddToCategoryAction struct {
	// The Category to add.
	Category CategoryResourceIdentifier `json:"category"`
	// A string representing a number between 0 and 1. Must start with `0.` and cannot end with `0`. If empty, any existing value will be removed.
	OrderHint *string `json:"orderHint,omitempty"`
	// If `true`, only the staged `categories` and `categoryOrderHints` are updated. If `false`, both the current and staged `categories` and `categoryOrderHints` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddToCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddToCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addToCategory", Alias: (*Alias)(&obj)})
}

type ProductAddVariantAction struct {
	// Value to set. Must be unique.
	Sku *string `json:"sku,omitempty"`
	// Value to set. Must be unique.
	Key *string `json:"key,omitempty"`
	// Embedded Prices for the Product Variant.
	Prices []PriceDraft `json:"prices"`
	// Images for the Product Variant.
	Images []Image `json:"images"`
	// Attributes for the Product Variant.
	Attributes []Attribute `json:"attributes"`
	// If `true` the new Product Variant is only staged. If `false` the new Product Variant is both current and staged.
	Staged *bool `json:"staged,omitempty"`
	// Media assets for the Product Variant.
	Assets []Asset `json:"assets"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductAddVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddVariantAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addVariant", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["prices"] == nil {
		delete(raw, "prices")
	}

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}

/**
*	Either `variantId` or `sku` is required. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductChangeAssetNameAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductChangeAssetOrderAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged `assets` is updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// All existing Asset `id`s of the ProductVariant in the desired new order.
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

/**
*	Assigns the specified Product Variant to the `masterVariant` and removes the same from `variants` at the same time. The current Master Variant becomes part of the `variants` array.
*	Either `variantId` or `sku` is required.
*
 */
type ProductChangeMasterVariantAction struct {
	// The `id` of the ProductVariant to become the Master Variant.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to become the Master Variant.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Master Variant is changed. If `false`, both the current and staged Master Variant are changed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangeMasterVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeMasterVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMasterVariant", Alias: (*Alias)(&obj)})
}

type ProductChangeNameAction struct {
	// Value to set. Must not be empty.
	Name LocalizedString `json:"name"`
	// If `true`, only the staged `name` is updated. If `false`, both the current and staged `name` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductChangePriceAction struct {
	// The `id` of the Embedded Price to update.
	PriceId string `json:"priceId"`
	// Value to set.
	Price PriceDraft `json:"price"`
	// If `true`, only the staged Embedded Price is updated. If `false`, both the current and staged Embedded Price are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePrice", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ProductSlugChanged](ctp:api:type:ProductSlugChangedMessage) Message.
 */
type ProductChangeSlugAction struct {
	// Value to set. Must not be empty. A Product can have the same slug for different [Locales](ctp:api:type:Locale), but it must be unique across the [Project](ctp:api:type:Project). Must match the pattern `^[A-Za-z0-9_-]{2,256}+$`.
	Slug LocalizedString `json:"slug"`
	// If `true`, only the staged `slug` is updated. If `false`, both the current and staged `slug` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type ProductLegacySetSkuAction struct {
	Sku       *string `json:"sku,omitempty"`
	VariantId int     `json:"variantId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductLegacySetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductLegacySetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "legacySetSku", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductMoveImageToPositionAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The URL of the image to update.
	ImageUrl string `json:"imageUrl"`
	// Position in `images` where the image should be moved. Must be between `0` and the total number of images minus `1`.
	Position int `json:"position"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

/**
*	Publishes product data from the Product's staged projection to its current projection.
*	Produces the [ProductPublished](ctp:api:type:ProductPublishedMessage) Message.
 */
type ProductPublishAction struct {
	// `All` or `Prices`
	Scope *ProductPublishScope `json:"scope,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The Asset to remove must be specified using either `assetId` or `assetKey`.
*
 */
type ProductRemoveAssetAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is removed. If `false`, both the current and staged Asset is removed.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to remove.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to remove.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ProductRemovedFromCategory](ctp:api:type:ProductRemovedFromCategoryMessage) Message.
 */
type ProductRemoveFromCategoryAction struct {
	// The Category to remove.
	Category CategoryResourceIdentifier `json:"category"`
	// If `true`, only the staged `categories` and `categoryOrderHints` are removed. If `false`, both the current and staged `categories` and `categoryOrderHints` are removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemoveFromCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveFromCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFromCategory", Alias: (*Alias)(&obj)})
}

/**
*	Removes a Product image and deletes it from the Content Delivery Network (external images are not deleted). Deletion from the CDN is not instant, which means the image file itself will stay available for some time after the deletion. Either `variantId` or `sku` is required.
*
 */
type ProductRemoveImageAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The URL of the image to remove.
	ImageUrl string `json:"imageUrl"`
	// If `true`, only the staged image is removed. If `false`, both the current and staged image is removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

type ProductRemovePriceAction struct {
	// The `id` of the Embedded Price to remove.
	PriceId string `json:"priceId"`
	// If `true`, only the staged Embedded Price is removed. If `false`, both the current and staged Embedded Price are removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemovePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePrice", Alias: (*Alias)(&obj)})
}

/**
*	Either `id` or `sku` is required. Produces the [ProductVariantDeleted](ctp:api:type:ProductVariantDeletedMessage) Message.
*
 */
type ProductRemoveVariantAction struct {
	// The `id` of the ProductVariant to remove.
	ID *int `json:"id,omitempty"`
	// The `sku` of the ProductVariant to remove.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged ProductVariant is removed. If `false`, both the current and staged ProductVariant is removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRemoveVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeVariant", Alias: (*Alias)(&obj)})
}

/**
*	Reverts the staged version of a Product to the current version. Produces the [ProductRevertedStagedChanges](ctp:api:type:ProductRevertedStagedChangesMessage) Message.
*
 */
type ProductRevertStagedChangesAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRevertStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedChanges", Alias: (*Alias)(&obj)})
}

/**
*	Reverts the staged version of a ProductVariant to the current version.
*
 */
type ProductRevertStagedVariantChangesAction struct {
	// The `id` of the ProductVariant to revert.
	VariantId int `json:"variantId"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductRevertStagedVariantChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedVariantChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedVariantChanges", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductSetAssetCustomFieldAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductSetAssetCustomTypeAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Asset with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Asset.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Asset.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductSetAssetDescriptionAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductSetAssetKeyAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId string `json:"assetId"`
	// Value to set. If empty, any existing value will be removed.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductSetAssetSourcesAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false` both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set.
	Sources []AssetSource `json:"sources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductSetAssetTagsAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetTagsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tags"] == nil {
		delete(raw, "tags")
	}

	return json.Marshal(raw)

}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductSetAttributeAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attribute. If empty, any existing value will be removed.
	//
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, use the list of values of all Attributes of the nested Product as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	Value interface{} `json:"value,omitempty"`
	// If `true`, only the staged Attribute is set. If `false`, both current and staged Attribute is set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

/**
*	Adds, removes, or changes a Product Attribute in all Product Variants at the same time.
*	This action is useful for setting values for Attributes with the [Constraint](ctp:api:type:AttributeConstraintEnum) `SameForAll`.
 */
type ProductSetAttributeInAllVariantsAction struct {
	// The name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attributes. If empty, any existing value will be removed.
	//
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, use the list of values of all Attributes of the nested Product as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	Value interface{} `json:"value,omitempty"`
	// If `true`, only the staged Attributes are set. If `false`, both the current and staged Attributes are set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetAttributeInAllVariantsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeInAllVariantsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributeInAllVariants", Alias: (*Alias)(&obj)})
}

type ProductSetCategoryOrderHintAction struct {
	// The `id` of the Category to add the `orderHint`.
	CategoryId string `json:"categoryId"`
	// A string representing a number between 0 and 1. Must start with `0.` and cannot end with `0`. If empty, any existing value will be removed.
	OrderHint *string `json:"orderHint,omitempty"`
	// If `true`, only the staged `categoryOrderHints` is updated. If `false`, both the current and staged `categoryOrderHints` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetCategoryOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetCategoryOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCategoryOrderHint", Alias: (*Alias)(&obj)})
}

type ProductSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
	// If `true`, only the staged `description` is updated. If `false`, both the current and staged `description` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Produces the [ProductPriceExternalDiscountSet](ctp:api:type:ProductPriceExternalDiscountSetMessage) Message.
*
 */
type ProductSetDiscountedPriceAction struct {
	// The `id` of the [Embedded Price](ctp:api:type:Price) to set the Discount.
	PriceId string `json:"priceId"`
	// If `true`, only the staged Embedded Price is updated. If `false`, both the current and staged Embedded Price are updated.
	Staged *bool `json:"staged,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	// The referenced [ProductDiscount](ctp:api:type:ProductDiscount) must have the Type `external`, be active, and its predicate must match the referenced Price.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetDiscountedPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDiscountedPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountedPrice", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductSetImageLabelAction struct {
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The URL of the image to set the label.
	ImageUrl string `json:"imageUrl"`
	// Value to set. If empty, any existing value will be removed.
	Label *string `json:"label,omitempty"`
	// If `true`, only the staged image is updated. If `false`, both the current and staged image is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

type ProductSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ProductSetMetaDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// If `true`, only the staged `metaDescription` is updated. If `false`, both the current and staged `metaDescription` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type ProductSetMetaKeywordsAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// If `true`, only the staged `metaKeywords` is updated. If `false`, both the current and staged `metaKeywords` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type ProductSetMetaTitleAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// If `true`, only the staged `metaTitle` is updated. If `false`, both the current and staged `metaTitle` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

/**
*	Controls whether the Prices of a Product Variant are embedded into the Product or standalone.
*
 */
type ProductSetPriceModeAction struct {
	// Specifies which type of Prices should be used when looking up a price for the Product.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetPriceModeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetPriceModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPriceMode", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductSetPricesAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The Embedded Prices to set.
	// Each Price must have its unique Price scope (with same currency, country, Customer Group, Channel, `validFrom` and `validUntil`).
	Prices []PriceDraft `json:"prices"`
	// If `true`, only the staged ProductVariant is updated. If `false`, both the current and staged ProductVariant are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetPricesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetPricesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPrices", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomFieldAction struct {
	// The `id` of the Embedded Price to update.
	PriceId string `json:"priceId"`
	// If `true`, only the staged Embedded Price Custom Field is updated. If `false`, both the current and staged Embedded Price Custom Field are updated.
	Staged *bool `json:"staged,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetProductPriceCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomField", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomTypeAction struct {
	// The `id` of the Embedded Price to update.
	PriceId string `json:"priceId"`
	// If `true`, only the staged Embedded Price is updated. If `false`, both the current and staged Embedded Price is updated.
	Staged *bool `json:"staged,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Price with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Embedded Price.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Embedded Price.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetProductPriceCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductSetProductVariantKeyAction struct {
	// The `id` of the ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Value to set. Must be unique. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
	// If `true`, only the staged `key` is set. If `false`, both the current and staged `key` are set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetProductVariantKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductVariantKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductVariantKey", Alias: (*Alias)(&obj)})
}

type ProductSetSearchKeywordsAction struct {
	// Value to set.
	SearchKeywords SearchKeywords `json:"searchKeywords"`
	// If `true`, only the staged `searchKeywords` is updated. If `false`, both the current and staged `searchKeywords` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetSearchKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSearchKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSearchKeywords", Alias: (*Alias)(&obj)})
}

/**
*	SKU cannot be changed or removed if it is associated with an [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type ProductSetSkuAction struct {
	// The `id` of the ProductVariant to update.
	VariantId int `json:"variantId"`
	// Value to set. Must be unique. If empty, any existing value will be removed.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged `sku` is updated. If `false`, both the current and staged `sku` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSku", Alias: (*Alias)(&obj)})
}

/**
*	Cannot be staged. Published Products are immediately updated.
*
 */
type ProductSetTaxCategoryAction struct {
	// The Tax Category to set. If empty, any existing value will be removed.
	TaxCategory *TaxCategoryResourceIdentifier `json:"taxCategory,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSetTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTaxCategory", Alias: (*Alias)(&obj)})
}

/**
*	If the existing [State](ctp:api:type:State) has set `transitions`, there must be a direct transition to the new State. If `transitions` is not set, no validation is performed. Produces the [ProductStateTransition](ctp:api:type:ProductStateTransitionMessage) Message.
*
 */
type ProductTransitionStateAction struct {
	// The State to transition to. If there is no existing State, this must be an initial State.
	State *StateResourceIdentifier `json:"state,omitempty"`
	// If `true`, validations are disabled.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

/**
*	Removes the current projection of the Product. The staged projection is unaffected. Unpublished Products only appear in query/search results with `staged=false`. Produces the [ProductUnpublished](ctp:api:type:ProductUnpublishedMessage) Message.
*
 */
type ProductUnpublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}
