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
	Key string `json:"key"`
	// The `productType` of a [Product](/../api/projects/products#product).
	// Maps to `Product.productType`.
	// The Reference to the [ProductType](/../api/projects/productTypes#producttype) with which the ProductDraft is associated.
	// If referenced ProductType does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary ProductType is created.
	ProductType ProductTypeKeyReference `json:"productType"`
	Name        LocalizedString         `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related product. Each slug must be unique across a project,
	// but a product can have the same slug for different languages. Allowed are alphabetic, numeric, underscore (_) and hyphen (-) characters.
	Slug LocalizedString `json:"slug"`
	// Maps to `Product.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// The Reference to the [Categories](/../api/projects/categories#category) with which the ProductDraft is associated.
	// If referenced Categories do not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Categories are created.
	Categories []CategoryKeyReference `json:"categories"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// The master Product variant.
	// Required if the `variants` array contains a Product Variant.
	MasterVariant *ProductVariantDraftImport `json:"masterVariant,omitempty"`
	// An array of related Product Variants.
	Variants []ProductVariantDraftImport `json:"variants"`
	// The Reference to the [TaxCategory](/../api/projects/taxCategories#taxcategory) with which the ProductDraft is associated.
	// If referenced TaxCategory does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary TaxCategory is created.
	TaxCategory *TaxCategoryKeyReference `json:"taxCategory,omitempty"`
	// Search keywords are primarily used by the suggester but are also considered for the full-text search. SearchKeywords is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag). The value to a language tag key is an array of SearchKeyword for the specific language.
	// ```json
	// {
	//   "en": [
	//     { "text": "Multi tool" },
	//     { "text": "Swiss Army Knife", "suggestTokenizer": { "type": "whitespace" } }
	//   ],
	//   "de": [
	//     {
	//       "text": "Schweizer Messer",
	//       "suggestTokenizer": {
	//         "type": "custom",
	//         "inputs": ["schweizer messer", "offiziersmesser", "sackmesser"]
	//       }
	//     }
	//   ]
	// }
	// ```
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitempty"`
	// The Reference to the [State](/../api/projects/states#state) with which the ProductDraft is associated.
	// If referenced State does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary State is created.
	State *StateKeyReference `json:"state,omitempty"`
	// If `publish` is set to either `true` or `false`, both staged and current projections are set to the same value provided by the import data.
	// If `publish` is not set, the staged projection is set to the provided import data, but the current projection stays unchanged.
	// However, if the import data contains no update, that is, if it matches the staged projection of the existing Product, the import induces no change in the existing Product whether `publish` is set or not.
	Publish *bool `json:"publish,omitempty"`
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
	Sku        *string            `json:"sku,omitempty"`
	Key        string             `json:"key"`
	Prices     []PriceDraftImport `json:"prices"`
	Attributes []Attribute        `json:"attributes"`
	Images     []Image            `json:"images"`
	Assets     []Asset            `json:"assets"`
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
	Value TypedMoney `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country *string `json:"country,omitempty"`
	// References a customer group by key.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// References a channel by key.
	Channel    *ChannelKeyReference `json:"channel,omitempty"`
	ValidFrom  *time.Time           `json:"validFrom,omitempty"`
	ValidUntil *time.Time           `json:"validUntil,omitempty"`
	// The custom fields for this category.
	Custom *Custom `json:"custom,omitempty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers"`
	Key   *string     `json:"key,omitempty"`
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
