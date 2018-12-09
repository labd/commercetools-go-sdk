// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Attribute struct {
	Value interface{} `json:"value"`
	Name  string      `json:"name"`
}

type AttributeValue struct{}
type CategoryOrderHints map[string]string

type CustomTokenizer struct {
	Inputs []string `json:"inputs"`
}

func (obj CustomTokenizer) MarshalJSON() ([]byte, error) {
	type Alias CustomTokenizer
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "custom", Alias: (*Alias)(&obj)})
}

type FacetResult interface{}
type AbstractFacetResult struct{}

func AbstractFacetResultDiscriminatorMapping(input FacetResult) FacetResult {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "filter":
		new := FilteredFacetResult{}
		mapstructure.Decode(input, &new)
		return new
	case "range":
		new := RangeFacetResult{}
		mapstructure.Decode(input, &new)
		return new
	case "terms":
		new := TermFacetResult{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type FacetResultRange struct {
	Total        int     `json:"total"`
	ToStr        string  `json:"toStr"`
	To           float64 `json:"to"`
	ProductCount int     `json:"productCount,omitempty"`
	Min          float64 `json:"min"`
	Mean         float64 `json:"mean"`
	Max          float64 `json:"max"`
	FromStr      string  `json:"fromStr"`
	From         float64 `json:"from"`
	Count        int     `json:"count"`
}

type FacetResultTerm struct {
	Term         interface{} `json:"term"`
	ProductCount int         `json:"productCount,omitempty"`
	Count        int         `json:"count"`
}

type FacetResults struct {
}

func (obj *FacetResults) UnmarshalJSON(data []byte) error {
	type Alias FacetResults
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	return nil
}

type FacetTypes string

const (
	FacetTypesTerms  FacetTypes = "terms"
	FacetTypesRange  FacetTypes = "range"
	FacetTypesFilter FacetTypes = "filter"
)

type FilteredFacetResult struct {
	ProductCount int `json:"productCount,omitempty"`
	Count        int `json:"count"`
}

func (obj FilteredFacetResult) MarshalJSON() ([]byte, error) {
	type Alias FilteredFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "filter", Alias: (*Alias)(&obj)})
}

type Product struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	TaxCategory            *TaxCategoryReference   `json:"taxCategory,omitempty"`
	State                  *StateReference         `json:"state,omitempty"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	ProductType            *ProductTypeReference   `json:"productType"`
	MasterData             *ProductCatalogData     `json:"masterData"`
	Key                    string                  `json:"key,omitempty"`
}

type ProductAddAssetAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Staged    bool        `json:"staged,omitempty"`
	Sku       string      `json:"sku,omitempty"`
	Position  float64     `json:"position,omitempty"`
	Asset     *AssetDraft `json:"asset"`
}

func (obj ProductAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

type ProductAddExternalImageAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Image     *Image `json:"image"`
}

func (obj ProductAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

type ProductAddPriceAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Staged    bool        `json:"staged,omitempty"`
	Sku       string      `json:"sku,omitempty"`
	Price     *PriceDraft `json:"price"`
}

func (obj ProductAddPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPrice", Alias: (*Alias)(&obj)})
}

type ProductAddToCategoryAction struct {
	Staged    bool               `json:"staged,omitempty"`
	OrderHint string             `json:"orderHint,omitempty"`
	Category  *CategoryReference `json:"category"`
}

func (obj ProductAddToCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddToCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addToCategory", Alias: (*Alias)(&obj)})
}

type ProductAddVariantAction struct {
	Staged     bool         `json:"staged,omitempty"`
	Sku        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Key        string       `json:"key,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
}

func (obj ProductAddVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductAddVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addVariant", Alias: (*Alias)(&obj)})
}

type ProductCatalogData struct {
	Staged           *ProductData `json:"staged"`
	Published        bool         `json:"published"`
	HasStagedChanges bool         `json:"hasStagedChanges"`
	Current          *ProductData `json:"current"`
}

type ProductChangeAssetNameAction struct {
	VariantID int              `json:"variantId,omitempty"`
	Staged    bool             `json:"staged,omitempty"`
	Sku       string           `json:"sku,omitempty"`
	Name      *LocalizedString `json:"name"`
	AssetKey  string           `json:"assetKey,omitempty"`
	AssetID   string           `json:"assetId,omitempty"`
}

func (obj ProductChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

type ProductChangeAssetOrderAction struct {
	VariantID  int      `json:"variantId,omitempty"`
	Staged     bool     `json:"staged,omitempty"`
	Sku        string   `json:"sku,omitempty"`
	AssetOrder []string `json:"assetOrder"`
}

func (obj ProductChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

type ProductChangeMasterVariantAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
}

func (obj ProductChangeMasterVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeMasterVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMasterVariant", Alias: (*Alias)(&obj)})
}

type ProductChangeNameAction struct {
	Staged bool             `json:"staged,omitempty"`
	Name   *LocalizedString `json:"name"`
}

func (obj ProductChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductChangePriceAction struct {
	Staged  bool        `json:"staged,omitempty"`
	PriceID string      `json:"priceId"`
	Price   *PriceDraft `json:"price"`
}

func (obj ProductChangePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePrice", Alias: (*Alias)(&obj)})
}

type ProductChangeSlugAction struct {
	Staged bool             `json:"staged,omitempty"`
	Slug   *LocalizedString `json:"slug"`
}

func (obj ProductChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type ProductData struct {
	Variants           []ProductVariant    `json:"variants"`
	Slug               *LocalizedString    `json:"slug"`
	SearchKeywords     *SearchKeywords     `json:"searchKeywords"`
	Name               *LocalizedString    `json:"name"`
	MetaTitle          *LocalizedString    `json:"metaTitle,omitempty"`
	MetaKeywords       *LocalizedString    `json:"metaKeywords,omitempty"`
	MetaDescription    *LocalizedString    `json:"metaDescription,omitempty"`
	MasterVariant      *ProductVariant     `json:"masterVariant"`
	Description        *LocalizedString    `json:"description,omitempty"`
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	Categories         []CategoryReference `json:"categories"`
}

type ProductDraft struct {
	Variants           []ProductVariantDraft `json:"variants,omitempty"`
	TaxCategory        *TaxCategoryReference `json:"taxCategory,omitempty"`
	State              *StateReference       `json:"state,omitempty"`
	Slug               *LocalizedString      `json:"slug"`
	SearchKeywords     *SearchKeywords       `json:"searchKeywords,omitempty"`
	Publish            bool                  `json:"publish,omitempty"`
	ProductType        *ProductTypeReference `json:"productType,omitempty"`
	Name               *LocalizedString      `json:"name"`
	MetaTitle          *LocalizedString      `json:"metaTitle,omitempty"`
	MetaKeywords       *LocalizedString      `json:"metaKeywords,omitempty"`
	MetaDescription    *LocalizedString      `json:"metaDescription,omitempty"`
	MasterVariant      *ProductVariantDraft  `json:"masterVariant,omitempty"`
	Key                string                `json:"key,omitempty"`
	Description        *LocalizedString      `json:"description,omitempty"`
	CategoryOrderHints *CategoryOrderHints   `json:"categoryOrderHints,omitempty"`
	Categories         []CategoryReference   `json:"categories,omitempty"`
}

type ProductLegacySetSkuAction struct {
	VariantID int    `json:"variantId"`
	Sku       string `json:"sku,omitempty"`
}

func (obj ProductLegacySetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductLegacySetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "legacySetSku", Alias: (*Alias)(&obj)})
}

type ProductMoveImageToPositionAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Position  int    `json:"position"`
	ImageURL  string `json:"imageUrl"`
}

func (obj ProductMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

type ProductPagedQueryResponse struct {
	Total   int       `json:"total,omitempty"`
	Offset  int       `json:"offset"`
	Count   int       `json:"count"`
	Results []Product `json:"results"`
}

type ProductProjection struct {
	Version                int                     `json:"version"`
	LastModifiedAt         time.Time               `json:"lastModifiedAt"`
	ID                     string                  `json:"id"`
	CreatedAt              time.Time               `json:"createdAt"`
	Variants               []ProductVariant        `json:"variants"`
	TaxCategory            *TaxCategoryReference   `json:"taxCategory,omitempty"`
	State                  *StateReference         `json:"state,omitempty"`
	Slug                   *LocalizedString        `json:"slug"`
	SearchKeywords         *SearchKeywords         `json:"searchKeywords,omitempty"`
	ReviewRatingStatistics *ReviewRatingStatistics `json:"reviewRatingStatistics,omitempty"`
	Published              bool                    `json:"published,omitempty"`
	ProductType            *ProductTypeReference   `json:"productType"`
	Name                   *LocalizedString        `json:"name"`
	MetaTitle              *LocalizedString        `json:"metaTitle,omitempty"`
	MetaKeywords           *LocalizedString        `json:"metaKeywords,omitempty"`
	MetaDescription        *LocalizedString        `json:"metaDescription,omitempty"`
	MasterVariant          *ProductVariant         `json:"masterVariant"`
	Key                    string                  `json:"key,omitempty"`
	HasStagedChanges       bool                    `json:"hasStagedChanges,omitempty"`
	Description            *LocalizedString        `json:"description,omitempty"`
	CategoryOrderHints     *CategoryOrderHints     `json:"categoryOrderHints,omitempty"`
	Categories             []CategoryReference     `json:"categories"`
}

type ProductProjectionPagedQueryResponse struct {
	Total   int                 `json:"total,omitempty"`
	Offset  int                 `json:"offset"`
	Count   int                 `json:"count"`
	Results []ProductProjection `json:"results"`
}

type ProductProjectionPagedSearchResponse struct {
	Total   int                 `json:"total,omitempty"`
	Offset  int                 `json:"offset"`
	Count   int                 `json:"count"`
	Results []ProductProjection `json:"results"`
	Facets  *FacetResults       `json:"facets"`
}

type ProductPublishAction struct {
	Scope ProductPublishScope `json:"scope,omitempty"`
}

func (obj ProductPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

type ProductReference struct {
	Key string   `json:"key,omitempty"`
	ID  string   `json:"id,omitempty"`
	Obj *Product `json:"obj,omitempty"`
}

func (obj ProductReference) MarshalJSON() ([]byte, error) {
	type Alias ProductReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "product", Alias: (*Alias)(&obj)})
}

type ProductRemoveAssetAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	AssetKey  string `json:"assetKey,omitempty"`
	AssetID   string `json:"assetId,omitempty"`
}

func (obj ProductRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

type ProductRemoveFromCategoryAction struct {
	Staged   bool               `json:"staged,omitempty"`
	Category *CategoryReference `json:"category"`
}

func (obj ProductRemoveFromCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveFromCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeFromCategory", Alias: (*Alias)(&obj)})
}

type ProductRemoveImageAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	ImageURL  string `json:"imageUrl"`
}

func (obj ProductRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

type ProductRemovePriceAction struct {
	Staged  bool   `json:"staged,omitempty"`
	PriceID string `json:"priceId"`
}

func (obj ProductRemovePriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemovePriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePrice", Alias: (*Alias)(&obj)})
}

type ProductRemoveVariantAction struct {
	Staged bool   `json:"staged,omitempty"`
	Sku    string `json:"sku,omitempty"`
	ID     int    `json:"id,omitempty"`
}

func (obj ProductRemoveVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRemoveVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeVariant", Alias: (*Alias)(&obj)})
}

type ProductRevertStagedChangesAction struct{}

func (obj ProductRevertStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedChanges", Alias: (*Alias)(&obj)})
}

type ProductRevertStagedVariantChangesAction struct {
	VariantID int `json:"variantId"`
}

func (obj ProductRevertStagedVariantChangesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductRevertStagedVariantChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "revertStagedVariantChanges", Alias: (*Alias)(&obj)})
}

type ProductSetAssetCustomFieldAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Staged    bool        `json:"staged,omitempty"`
	Sku       string      `json:"sku,omitempty"`
	Name      string      `json:"name"`
	AssetKey  string      `json:"assetKey,omitempty"`
	AssetID   string      `json:"assetId,omitempty"`
}

func (obj ProductSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

type ProductSetAssetCustomTypeAction struct {
	VariantID int            `json:"variantId,omitempty"`
	Type      *TypeReference `json:"type,omitempty"`
	Staged    bool           `json:"staged,omitempty"`
	Sku       string         `json:"sku,omitempty"`
	Fields    interface{}    `json:"fields,omitempty"`
	AssetKey  string         `json:"assetKey,omitempty"`
	AssetID   string         `json:"assetId,omitempty"`
}

func (obj ProductSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

type ProductSetAssetDescriptionAction struct {
	VariantID   int              `json:"variantId,omitempty"`
	Staged      bool             `json:"staged,omitempty"`
	Sku         string           `json:"sku,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
	AssetKey    string           `json:"assetKey,omitempty"`
	AssetID     string           `json:"assetId,omitempty"`
}

func (obj ProductSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

type ProductSetAssetKeyAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	AssetKey  string `json:"assetKey,omitempty"`
	AssetID   string `json:"assetId"`
}

func (obj ProductSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

type ProductSetAssetSourcesAction struct {
	VariantID int           `json:"variantId,omitempty"`
	Staged    bool          `json:"staged,omitempty"`
	Sources   []AssetSource `json:"sources"`
	Sku       string        `json:"sku,omitempty"`
	AssetKey  string        `json:"assetKey,omitempty"`
	AssetID   string        `json:"assetId,omitempty"`
}

func (obj ProductSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

type ProductSetAssetTagsAction struct {
	VariantID int      `json:"variantId,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Staged    bool     `json:"staged,omitempty"`
	Sku       string   `json:"sku,omitempty"`
	AssetKey  string   `json:"assetKey,omitempty"`
	AssetID   string   `json:"assetId,omitempty"`
}

func (obj ProductSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

type ProductSetAttributeAction struct {
	VariantID int         `json:"variantId,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Staged    bool        `json:"staged,omitempty"`
	Sku       string      `json:"sku,omitempty"`
	Name      string      `json:"name"`
}

func (obj ProductSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

type ProductSetAttributeInAllVariantsAction struct {
	Value  interface{} `json:"value,omitempty"`
	Staged bool        `json:"staged,omitempty"`
	Name   string      `json:"name"`
}

func (obj ProductSetAttributeInAllVariantsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetAttributeInAllVariantsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributeInAllVariants", Alias: (*Alias)(&obj)})
}

type ProductSetCategoryOrderHintAction struct {
	Staged     bool   `json:"staged,omitempty"`
	OrderHint  string `json:"orderHint,omitempty"`
	CategoryID string `json:"categoryId"`
}

func (obj ProductSetCategoryOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetCategoryOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCategoryOrderHint", Alias: (*Alias)(&obj)})
}

type ProductSetDescriptionAction struct {
	Staged      bool             `json:"staged,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
}

func (obj ProductSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ProductSetDiscountedPriceAction struct {
	Staged     bool             `json:"staged,omitempty"`
	PriceID    string           `json:"priceId"`
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

func (obj ProductSetDiscountedPriceAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetDiscountedPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountedPrice", Alias: (*Alias)(&obj)})
}

type ProductSetImageLabelAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Label     string `json:"label,omitempty"`
	ImageURL  string `json:"imageUrl"`
}

func (obj ProductSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

type ProductSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ProductSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ProductSetMetaDescriptionAction struct {
	Staged          bool             `json:"staged,omitempty"`
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
}

func (obj ProductSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type ProductSetMetaKeywordsAction struct {
	Staged       bool             `json:"staged,omitempty"`
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
}

func (obj ProductSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type ProductSetMetaTitleAction struct {
	Staged    bool             `json:"staged,omitempty"`
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
}

func (obj ProductSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

type ProductSetPricesAction struct {
	VariantID int          `json:"variantId,omitempty"`
	Staged    bool         `json:"staged,omitempty"`
	Sku       string       `json:"sku,omitempty"`
	Prices    []PriceDraft `json:"prices"`
}

func (obj ProductSetPricesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetPricesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPrices", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomFieldAction struct {
	Value   interface{} `json:"value,omitempty"`
	Staged  bool        `json:"staged,omitempty"`
	PriceID string      `json:"priceId"`
	Name    string      `json:"name"`
}

func (obj ProductSetProductPriceCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomField", Alias: (*Alias)(&obj)})
}

type ProductSetProductPriceCustomTypeAction struct {
	Type    *TypeReference  `json:"type,omitempty"`
	Staged  bool            `json:"staged,omitempty"`
	PriceID string          `json:"priceId"`
	Fields  *FieldContainer `json:"fields,omitempty"`
}

func (obj ProductSetProductPriceCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductPriceCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductPriceCustomType", Alias: (*Alias)(&obj)})
}

type ProductSetProductVariantKeyAction struct {
	VariantID int    `json:"variantId,omitempty"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
	Key       string `json:"key,omitempty"`
}

func (obj ProductSetProductVariantKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetProductVariantKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductVariantKey", Alias: (*Alias)(&obj)})
}

type ProductSetSearchKeywordsAction struct {
	Staged         bool            `json:"staged,omitempty"`
	SearchKeywords *SearchKeywords `json:"searchKeywords"`
}

func (obj ProductSetSearchKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSearchKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSearchKeywords", Alias: (*Alias)(&obj)})
}

type ProductSetSkuAction struct {
	VariantID int    `json:"variantId"`
	Staged    bool   `json:"staged,omitempty"`
	Sku       string `json:"sku,omitempty"`
}

func (obj ProductSetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSku", Alias: (*Alias)(&obj)})
}

type ProductSetTaxCategoryAction struct {
	TaxCategory *TaxCategoryReference `json:"taxCategory,omitempty"`
}

func (obj ProductSetTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ProductSetTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTaxCategory", Alias: (*Alias)(&obj)})
}

type ProductTransitionStateAction struct {
	State *StateReference `json:"state,omitempty"`
	Force bool            `json:"force,omitempty"`
}

func (obj ProductTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type ProductUnpublishAction struct{}

func (obj ProductUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}

type ProductUpdate struct {
	Version int                   `json:"version"`
	Actions []ProductUpdateAction `json:"actions"`
}

func (obj *ProductUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractProductUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ProductUpdateAction interface{}
type AbstractProductUpdateAction struct{}

func AbstractProductUpdateActionDiscriminatorMapping(input ProductUpdateAction) ProductUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addAsset":
		new := ProductAddAssetAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addExternalImage":
		new := ProductAddExternalImageAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addPrice":
		new := ProductAddPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addToCategory":
		new := ProductAddToCategoryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "addVariant":
		new := ProductAddVariantAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAssetName":
		new := ProductChangeAssetNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAssetOrder":
		new := ProductChangeAssetOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeMasterVariant":
		new := ProductChangeMasterVariantAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ProductChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changePrice":
		new := ProductChangePriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeSlug":
		new := ProductChangeSlugAction{}
		mapstructure.Decode(input, &new)
		return new
	case "legacySetSku":
		new := ProductLegacySetSkuAction{}
		mapstructure.Decode(input, &new)
		return new
	case "moveImageToPosition":
		new := ProductMoveImageToPositionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "publish":
		new := ProductPublishAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeAsset":
		new := ProductRemoveAssetAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeFromCategory":
		new := ProductRemoveFromCategoryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeImage":
		new := ProductRemoveImageAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removePrice":
		new := ProductRemovePriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeVariant":
		new := ProductRemoveVariantAction{}
		mapstructure.Decode(input, &new)
		return new
	case "revertStagedChanges":
		new := ProductRevertStagedChangesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "revertStagedVariantChanges":
		new := ProductRevertStagedVariantChangesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetCustomField":
		new := ProductSetAssetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetCustomType":
		new := ProductSetAssetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetDescription":
		new := ProductSetAssetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetKey":
		new := ProductSetAssetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetSources":
		new := ProductSetAssetSourcesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetTags":
		new := ProductSetAssetTagsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAttribute":
		new := ProductSetAttributeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAttributeInAllVariants":
		new := ProductSetAttributeInAllVariantsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCategoryOrderHint":
		new := ProductSetCategoryOrderHintAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := ProductSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDiscountedPrice":
		new := ProductSetDiscountedPriceAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setImageLabel":
		new := ProductSetImageLabelAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ProductSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaDescription":
		new := ProductSetMetaDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaKeywords":
		new := ProductSetMetaKeywordsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaTitle":
		new := ProductSetMetaTitleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setPrices":
		new := ProductSetPricesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setProductPriceCustomField":
		new := ProductSetProductPriceCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setProductPriceCustomType":
		new := ProductSetProductPriceCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setProductVariantKey":
		new := ProductSetProductVariantKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setSearchKeywords":
		new := ProductSetSearchKeywordsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setSku":
		new := ProductSetSkuAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTaxCategory":
		new := ProductSetTaxCategoryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionState":
		new := ProductTransitionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "unpublish":
		new := ProductUnpublishAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type ProductVariant struct {
	Sku                   string                      `json:"sku,omitempty"`
	ScopedPriceDiscounted bool                        `json:"scopedPriceDiscounted,omitempty"`
	ScopedPrice           *ScopedPrice                `json:"scopedPrice,omitempty"`
	Prices                []Price                     `json:"prices,omitempty"`
	Price                 *Price                      `json:"price,omitempty"`
	Key                   string                      `json:"key,omitempty"`
	IsMatchingVariant     bool                        `json:"isMatchingVariant,omitempty"`
	Images                []Image                     `json:"images,omitempty"`
	ID                    int                         `json:"id"`
	Availability          *ProductVariantAvailability `json:"availability,omitempty"`
	Attributes            []Attribute                 `json:"attributes,omitempty"`
	Assets                []Asset                     `json:"assets,omitempty"`
}

type ProductVariantAvailability struct {
	RestockableInDays int                                   `json:"restockableInDays,omitempty"`
	IsOnStock         bool                                  `json:"isOnStock,omitempty"`
	Channels          *ProductVariantChannelAvailabilityMap `json:"channels,omitempty"`
	AvailableQuantity int                                   `json:"availableQuantity,omitempty"`
}

type ProductVariantChannelAvailability struct {
	RestockableInDays int  `json:"restockableInDays,omitempty"`
	IsOnStock         bool `json:"isOnStock,omitempty"`
	AvailableQuantity int  `json:"availableQuantity,omitempty"`
}

type ProductVariantChannelAvailabilityMap struct {
}

type ProductVariantDraft struct {
	Sku        string       `json:"sku,omitempty"`
	Prices     []PriceDraft `json:"prices,omitempty"`
	Key        string       `json:"key,omitempty"`
	Images     []Image      `json:"images,omitempty"`
	Attributes []Attribute  `json:"attributes,omitempty"`
	Assets     []AssetDraft `json:"assets,omitempty"`
}

type RangeFacetResult struct {
	Ranges []FacetResultRange `json:"ranges"`
}

func (obj RangeFacetResult) MarshalJSON() ([]byte, error) {
	type Alias RangeFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "range", Alias: (*Alias)(&obj)})
}

type SearchKeyword struct {
	Text             string           `json:"text"`
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer,omitempty"`
}

func (obj *SearchKeyword) UnmarshalJSON(data []byte) error {
	type Alias SearchKeyword
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.SuggestTokenizer != nil {
		obj.SuggestTokenizer = AbstractSuggestTokenizerDiscriminatorMapping(obj.SuggestTokenizer)
	}

	return nil
}

type SearchKeywords struct {
}

type SuggestTokenizer interface{}
type AbstractSuggestTokenizer struct{}

func AbstractSuggestTokenizerDiscriminatorMapping(input SuggestTokenizer) SuggestTokenizer {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "custom":
		new := CustomTokenizer{}
		mapstructure.Decode(input, &new)
		return new
	case "whitespace":
		new := WhitespaceTokenizer{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type Suggestion struct {
	Text string `json:"text"`
}

type SuggestionResult struct {
}

type TermFacetResult struct {
	Total    int                 `json:"total"`
	Terms    []FacetResultTerm   `json:"terms"`
	Other    int                 `json:"other"`
	Missing  int                 `json:"missing"`
	DataType TermFacetResultType `json:"dataType"`
}

func (obj TermFacetResult) MarshalJSON() ([]byte, error) {
	type Alias TermFacetResult
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "terms", Alias: (*Alias)(&obj)})
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

type WhitespaceTokenizer struct{}

func (obj WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type Alias WhitespaceTokenizer
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "whitespace", Alias: (*Alias)(&obj)})
}
