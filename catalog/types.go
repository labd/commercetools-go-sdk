package catalog

import (
	"time"

	"github.com/labd/commercetools-go-sdk/common"
)

type CategoryOrderHints map[string]float64 // FIXME

type SearchKeyword struct {
	Text             string           `json:"text"`
	SuggestTokenizer SuggestTokenizer `json:"suggestTokenizer"`
}

type SuggestTokenizer struct {
	Type   string   `json:"type"`
	Inputs []string `json:"inputs,omitempty"`
}

// ProductDraft is the representation to be sent to the server when creating a
// new product.
type ProductDraft struct {
	Key                string                      `json:"key,omitempty"`
	Name               common.LocalizedString      `json:"name"`
	ProductType        common.Reference            `json:"productType"`
	Slug               common.LocalizedString      `json:"slug"`
	Description        common.LocalizedString      `json:"description,omitempty"`
	Categories         []common.ResourceIdentifier `json:"categories,omitempty"`
	CategoryOrderHints CategoryOrderHints          `json:"categoryOrderHints,omitempty"`
	MetaTitle          common.LocalizedString      `json:"metaTitle,omitempty"`
	MetaDescription    common.LocalizedString      `json:"metaDescription,omitempty"`
	MetaKeywords       common.LocalizedString      `json:"metaKeywords,omitempty"`
	MasterVariant      ProductVariantDraft         `json:"masterVariant,omitempty"`
	Variants           []ProductVariantDraft       `json:"variants,omitempty"`
	TaxCategory        common.Reference            `json:"taxCategory,omitempty"`
	SearchKeywords     []SearchKeyword             `json:"searchKeywords,omitempty"`
	State              common.Reference            `json:"state,omitempty"`
	Publish            bool                        `json:"publish"`
}

type ProductVariantDraft struct {
	SKU        string              `json:"sku,omitempty"`
	Key        string              `json:"key,omitempty"`
	Prices     []PriceDraft        `json:"prices,omitempty"`
	Images     []Image             `json:"images,omitempty"`
	Assets     []common.AssetDraft `json:"assets,omitempty"`
	Attributes []Attribute         `json:"attributes,omitempty"`
}

type PriceDraft struct {
	Value         common.Money     `json:"value"` // FIXME
	Country       string           `json:"country"`
	CustomerGroup common.Reference `json:"customerGroup,omitempty"`
	Channel       common.Reference `json:"channel,omitempty"`
	ValidFrom     time.Time        `json:"validFrom,omitempty"`
	ValidUntil    time.Time        `json:"validUntil,omitempty"`
	Tiers         []PriceTier      `json:"tiers,omitempty"`
	// Custom FIXME
}

type Product struct {
	ID                     string
	Key                    string
	Version                int
	CreatedAt              time.Time
	LastModifiedAt         time.Time
	ProductType            common.Reference
	MasterData             ProductCatalogData
	TaxCategory            common.Reference
	State                  common.Reference
	ReviewRatingStatistics ReviewRatingStatistics
}

type ProductVariant struct {
	ID         int
	Sku        string
	Key        string
	Prices     []Price `json:"prices"`
	Attributes []Attribute
	Price      Price `json:"price"`
	Images     []Image
	Assets     []common.Asset
}

type ProductCatalogData struct {
	Published        bool
	Current          ProductData
	Staged           ProductData
	HasStagedChanged bool
}

type ProductData struct {
	Name               common.LocalizedString `json:"name"`
	Categories         []common.Reference     `json:"categories"`
	CategoryOrderHints CategoryOrderHints     `json:"categoryOrderHints,omitempty"`
	Description        common.LocalizedString `json:"description,omitempty"`
	Slug               common.LocalizedString `json:"slug"`
	MetaTitle          common.LocalizedString `json:"metaTitle,omitempty"`
	MetaDescription    common.LocalizedString `json:"metaDescription,omitempty"`
	MetaKeywords       common.LocalizedString `json:"metaKeywords,omitempty"`
	MasterVariant      ProductVariant         `json:"masterVariant"`
	Variants           []ProductVariant       `json:"variants"`
}

// TODO
type ReviewRatingStatistics struct {
}

type Price struct {
	ID            string           `json:"id"`
	Value         common.Money     `json:"value"` // FIXME
	Country       string           `json:"country"`
	CustomerGroup common.Reference `json:"customerGroup,omitempty"`
	Channel       common.Reference `json:"channel,omitempty"`
	ValidFrom     time.Time        `json:"validFrom,omitempty"`
	ValidUntil    time.Time        `json:"validUntil,omitempty"`
	Tiers         []PriceTier      `json:"tiers,omitempty"`

	//Discounted    DiscountedPrice
	// Custom
}

type PriceTier struct {
	MinimumQuantity int
	Value           common.Money // BaseMoney
}

type Image struct {
	URL        string          `json:"url"`
	Dimensions ImageDimensions `json:"dimensions"`
	Label      string          `json:"label"`
}

type ImageDimensions struct {
	H int `json:"h"`
	W int `json:"w"`
}

type Attribute struct{}
