// Generated file, please do not change!!!
package importapi

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
	// If referenced ProductType does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary ProductType is created.
	ProductType ProductTypeKeyReference `json:"productType"`
	Name        LocalizedString         `json:"name"`
	// Human-readable identifiers usually used as deep-link URL to the related product. Each slug must be unique across a project,
	// but a product can have the same slug for different languages. Allowed are alphabetic, numeric, underscore (_) and hyphen (-) characters.
	Slug LocalizedString `json:"slug"`
	// Maps to `Product.description`.
	Description *LocalizedString `json:"description,omitEmpty"`
	// The Reference to the [Categories](/../api/projects/categories#category) with which the ProductDraft is associated.
	// If referenced Categories do not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Categories are created.
	Categories []CategoryKeyReference `json:"categories,omitEmpty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaTitle *LocalizedString `json:"metaTitle,omitEmpty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaDescription *LocalizedString `json:"metaDescription,omitEmpty"`
	// A localized string is a JSON object where the keys are of [IETF language tag](https://en.wikipedia.org/wiki/IETF_language_tag), and the values the corresponding strings used for that language.
	// ```json
	// {
	//   "de": "Hundefutter",
	//   "en": "dog food"
	// }
	// ```
	MetaKeywords *LocalizedString `json:"metaKeywords,omitEmpty"`
	// The master Product variant.
	// Required if the `variants` array contains a Product Variant.
	MasterVariant *ProductVariantDraftImport `json:"masterVariant,omitEmpty"`
	// An array of related Product Variants.
	Variants []ProductVariantDraftImport `json:"variants,omitEmpty"`
	// The Reference to the [TaxCategory](/../api/projects/taxCategories#taxcategory) with which the ProductDraft is associated.
	// If referenced TaxCategory does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary TaxCategory is created.
	TaxCategory TaxCategoryKeyReference `json:"taxCategory,omitEmpty"`
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
	SearchKeywords *SearchKeywords `json:"searchKeywords,omitEmpty"`
	// The Reference to the [State](/../api/projects/states#state) with which the ProductDraft is associated.
	// If referenced State does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary State is created.
	State StateKeyReference `json:"state,omitEmpty"`
	// If there were updates, only the updates will be published to `staged` and `current` projection.
	Publish bool `json:"publish,omitEmpty"`
}

/**
*	The representation of a Product Variant Draft for the import purpose.
*
 */
type ProductVariantDraftImport struct {
	Sku        string             `json:"sku,omitEmpty"`
	Key        string             `json:"key,omitEmpty"`
	Prices     []PriceDraftImport `json:"prices,omitEmpty"`
	Attributes []Attribute        `json:"attributes,omitEmpty"`
	Images     []Image            `json:"images,omitEmpty"`
	Assets     []Asset            `json:"assets,omitEmpty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantDraftImport) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantDraftImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

/**
*	The representation of a Price Draft for the import purpose.
*
 */
type PriceDraftImport struct {
	// TypedMoney is what is called BaseMoney in the HTTP API.
	Value TypedMoney `json:"value"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country,omitEmpty"`
	// References a customer group by its key.
	CustomerGroup CustomerGroupKeyReference `json:"customerGroup,omitEmpty"`
	// References a channel by its key.
	Channel    ChannelKeyReference `json:"channel,omitEmpty"`
	ValidFrom  time.Time           `json:"validFrom,omitEmpty"`
	ValidUntil time.Time           `json:"validUntil,omitEmpty"`
	// The custom fields for this category.
	Custom *Custom `json:"custom,omitEmpty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitEmpty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers,omitEmpty"`
	Key   string      `json:"key,omitEmpty"`
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
