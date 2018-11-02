package products

import (
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
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
	Key                string                             `json:"key,omitempty"`
	Name               commercetools.LocalizedString      `json:"name"`
	ProductType        commercetools.Reference            `json:"productType"`
	Slug               commercetools.LocalizedString      `json:"slug"`
	Description        commercetools.LocalizedString      `json:"description,omitempty"`
	Categories         []commercetools.ResourceIdentifier `json:"categories,omitempty"`
	CategoryOrderHints CategoryOrderHints                 `json:"categoryOrderHints,omitempty"`
	MetaTitle          commercetools.LocalizedString      `json:"metaTitle,omitempty"`
	MetaDescription    commercetools.LocalizedString      `json:"metaDescription,omitempty"`
	MetaKeywords       commercetools.LocalizedString      `json:"metaKeywords,omitempty"`
	MasterVariant      ProductVariantDraft                `json:"masterVariant,omitempty"`
	Variants           []ProductVariantDraft              `json:"variants,omitempty"`
	TaxCategory        commercetools.Reference            `json:"taxCategory,omitempty"`
	SearchKeywords     []SearchKeyword                    `json:"searchKeywords,omitempty"`
	State              *commercetools.Reference           `json:"state,omitempty"`
	Publish            bool                               `json:"publish"`
}

type ProductVariantDraft struct {
	SKU        string                     `json:"sku,omitempty"`
	Key        string                     `json:"key,omitempty"`
	Prices     []PriceDraft               `json:"prices,omitempty"`
	Images     []Image                    `json:"images,omitempty"`
	Assets     []commercetools.AssetDraft `json:"assets,omitempty"`
	Attributes []Attribute                `json:"attributes,omitempty"`
}

type PriceDraft struct {
	Value         commercetools.Money     `json:"value"` // FIXME
	Country       string                  `json:"country"`
	CustomerGroup commercetools.Reference `json:"customerGroup,omitempty"`
	Channel       commercetools.Reference `json:"channel,omitempty"`
	ValidFrom     time.Time               `json:"validFrom,omitempty"`
	ValidUntil    time.Time               `json:"validUntil,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`
	// Custom FIXME
}

type Product struct {
	ID                     string
	Key                    string
	Version                int
	CreatedAt              time.Time
	LastModifiedAt         time.Time
	ProductType            commercetools.Reference
	MasterData             ProductCatalogData
	TaxCategory            commercetools.Reference
	State                  commercetools.Reference
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
	Assets     []commercetools.Asset
}

type ProductCatalogData struct {
	Published        bool
	Current          ProductData
	Staged           ProductData
	HasStagedChanged bool
}

type ProductData struct {
	Name               commercetools.LocalizedString `json:"name"`
	Categories         []commercetools.Reference     `json:"categories"`
	CategoryOrderHints CategoryOrderHints            `json:"categoryOrderHints,omitempty"`
	Description        commercetools.LocalizedString `json:"description,omitempty"`
	Slug               commercetools.LocalizedString `json:"slug"`
	MetaTitle          commercetools.LocalizedString `json:"metaTitle,omitempty"`
	MetaDescription    commercetools.LocalizedString `json:"metaDescription,omitempty"`
	MetaKeywords       commercetools.LocalizedString `json:"metaKeywords,omitempty"`
	MasterVariant      ProductVariant                `json:"masterVariant"`
	Variants           []ProductVariant              `json:"variants"`
}

// TODO
type ReviewRatingStatistics struct {
}

type Price struct {
	ID            string                  `json:"id"`
	Value         commercetools.Money     `json:"value"` // FIXME
	Country       string                  `json:"country"`
	CustomerGroup commercetools.Reference `json:"customerGroup,omitempty"`
	Channel       commercetools.Reference `json:"channel,omitempty"`
	ValidFrom     time.Time               `json:"validFrom,omitempty"`
	ValidUntil    time.Time               `json:"validUntil,omitempty"`
	Tiers         []PriceTier             `json:"tiers,omitempty"`

	//Discounted    DiscountedPrice
	// Custom
}

type PriceTier struct {
	MinimumQuantity int
	Value           commercetools.Money // BaseMoney
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
