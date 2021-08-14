// Generated file, please do not change!!!
package importapi

import (
	"encoding/json"
	"time"
)

type SubRate struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TaxRate struct {
	Id              string  `json:"id,omitEmpty"`
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	IncludedInPrice bool    `json:"includedInPrice"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country  string    `json:"country"`
	State    string    `json:"state,omitEmpty"`
	SubRates []SubRate `json:"subRates,omitEmpty"`
}

/**
*	The data representation for a Price to be imported that is persisted as a [Price](/../api/projects/products#price) in the Project.
*
 */
type PriceImport struct {
	Key string `json:"key"`
	// Maps to `Price.value`. TypedMoney is what is called BaseMoney in the HTTP API.
	Value TypedMoney `json:"value"`
	// Maps to `Price.county`.
	Country string `json:"country,omitEmpty"`
	// Maps to `Price.validFrom`.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// Maps to `Price.validUntil`.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
	// The Reference to the [CustomerGroup](/../api/projects/customerGroups#customergroup) with which the Price is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary CustomerGroup is created.
	CustomerGroup CustomerGroupKeyReference `json:"customerGroup,omitEmpty"`
	// The Reference to the [Channel](/../api/projects/channels#channel) with which the Price is associated.
	// If referenced Channel does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Channel is created.
	Channel ChannelKeyReference `json:"channel,omitEmpty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitEmpty"`
	// Only the Price updates will be published to `staged` and `current` projection.
	Publish bool `json:"publish,omitEmpty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers,omitEmpty"`
	// The custom fields for this price.
	Custom *Custom `json:"custom,omitEmpty"`
	// The ProductVariant in which this Price is contained.
	// The Reference to the [ProductVariant](/../api/projects/products#productvariant) with which the Price is associated.
	// If referenced ProductVariant does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary ProductVariant is created.
	ProductVariant ProductVariantKeyReference `json:"productVariant"`
	// The Product in which the Product Variant containing this Price is contained. Maps to `ProductVariant.product`.
	// The Reference to the [Product](/../api/projects/products#product) with which the Price is associated.
	// If referenced Product does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `Unresolved` until the necessary Product is created.
	Product ProductKeyReference `json:"product"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *PriceImport) UnmarshalJSON(data []byte) error {
	type Alias PriceImport
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
