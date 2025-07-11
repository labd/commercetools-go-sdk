package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	The representation of a Product Draft for the import purpose.
*
 */
type ProductDraftImport struct {
	// User-defined unique identifier. If a [Product](ctp:api:type:Product) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `Product.productType`. If the referenced [ProductType](ctp:api:type:ProductType) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ProductType is created.
	ProductType ProductTypeKeyReference `json:"productType"`
	// Maps to `ProductData.name`.
	Name LocalizedString `json:"name"`
	// Maps to `ProductData.slug`.
	Slug LocalizedString `json:"slug"`
	// Maps to `ProductData.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `ProductData.categories`. If the referenced [Categories](ctp:api:type:Category) do not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Categories are created.
	Categories []CategoryKeyReference `json:"categories"`
	Attributes []Attribute            `json:"attributes"`
	// Maps to `ProductData.metaTitle`.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Maps to `ProductData.metaDescription`.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Maps to `ProductData.metaKeywords`.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// The master ProductVariant.
	// Required if `variants` contains at least one ProductVariant.
	MasterVariant *ProductVariantDraftImport `json:"masterVariant,omitempty"`
	// An array of related ProductVariants.
	Variants []ProductVariantDraftImport `json:"variants"`
	// Maps to `Product.taxCategory`. If the referenced [TaxCategory](ctp:api:type:TaxCategory) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Maps to `ProductData.searchKeywords`.
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// Maps to `Product.state`. If the referenced [State](ctp:api:type:State) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced State is created.
	State *StateKeyReference `json:"state,omitempty"`
	// Determines the published status and current/staged projection of the Product. For more information, see [Managing the published state of Products](/import-export/best-practices#manage-published-state-of-products).
	Publish *bool `json:"publish,omitempty"`
	// Maps to `Product.priceMode`. If not provided, the existing `Product.priceMode` is not changed.
	PriceMode *ProductPriceModeEnum `json:"priceMode,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDraftImport) UnmarshalJSON(data []byte) error {
	type Alias ProductDraftImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDraftImport) MarshalJSON() ([]byte, error) {
	type Alias ProductDraftImport
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

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["variants"] == nil {
		delete(raw, "variants")
	}

	return json.Marshal(raw)

}

/**
*	The representation of a Product Variant Draft for the import purpose.
*
 */
type ProductVariantDraftImport struct {
	// User-defined unique SKU of the Product Variant.
	Sku *string `json:"sku,omitempty"`
	// User-defined unique identifier for the ProductVariant.
	Key string `json:"key"`
	// The Embedded Prices for the Product Variant.
	// Each Price must have its unique Price scope (with same currency, country, Customer Group, Channel, `validFrom` and `validUntil`).
	Prices []PriceDraftImport `json:"prices"`
	// Attributes according to the respective AttributeDefinition.
	Attributes []Attribute `json:"attributes"`
	// Images for the Product Variant.
	Images []Image `json:"images"`
	// Media assets for the Product Variant.
	Assets []Asset `json:"assets"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantDraftImport) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantDraftImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantDraftImport) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantDraftImport
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
*	The representation of a Price Draft for the import purpose.
*
 */
type PriceDraftImport struct {
	// Money value of this Price.
	Value TypedMoney `json:"value"`
	// Set this field if this Price is only valid for the specified country.
	Country *string `json:"country,omitempty"`
	// Set this field if this Price is only valid for the referenced [CustomerGroup](ctp:api:type:CustomerGroup). If the referenced CustomerGroup does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Set this field if this Price is only valid for the referenced `ProductDistribution` [Channel](ctp:api:type:Channel). If the referenced Channel does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Set this field if this Price is only valid from the specified date and time. Must be at least 1 ms earlier than `validUntil`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Set this field if this Price is only valid until the specified date and time. Must be at least 1 ms later than `validFrom`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Custom Fields for the Embedded Price.
	Custom *Custom `json:"custom,omitempty"`
	// Set this field to add a DiscountedPrice from an **external service**.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers"`
	// User-defined unique identifier for the Embedded Price.
	Key string `json:"key"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PriceDraftImport) UnmarshalJSON(data []byte) error {
	type Alias PriceDraftImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoney(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceDraftImport) MarshalJSON() ([]byte, error) {
	type Alias PriceDraftImport
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

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}
