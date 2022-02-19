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
	ID              *string `json:"id,omitempty"`
	Name            string  `json:"name"`
	Amount          float64 `json:"amount"`
	IncludedInPrice bool    `json:"includedInPrice"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country  string    `json:"country"`
	State    *string   `json:"state,omitempty"`
	SubRates []SubRate `json:"subRates"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxRate) MarshalJSON() ([]byte, error) {
	type Alias TaxRate
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

	if target["subRates"] == nil {
		delete(target, "subRates")
	}

	return json.Marshal(target)
}

/**
*	The data representation for a Price to be imported that is persisted as a [Price](/../api/projects/products#price) in the Project.
*
 */
type PriceImport struct {
	Key string `json:"key"`
	// Maps to `Price.value`.
	Value TypedMoney `json:"value"`
	// Maps to `Price.county`.
	Country *string `json:"country,omitempty"`
	// Maps to `Price.validFrom`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Maps to `Price.validUntil`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// The Reference to the [CustomerGroup](/../api/projects/customerGroups#customergroup) with which the Price is associated.
	// If referenced CustomerGroup does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// The Reference to the [Channel](/../api/projects/channels#channel) with which the Price is associated.
	// If referenced Channel does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Channel is created.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Only the Price updates will be published to `staged` and `current` projection.
	Publish *bool `json:"publish,omitempty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers"`
	// The custom fields for this price.
	Custom *Custom `json:"custom,omitempty"`
	// The ProductVariant in which this Price is contained.
	// The Reference to the [ProductVariant](/../api/projects/products#productvariant) with which the Price is associated.
	// If referenced ProductVariant does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary ProductVariant is created.
	ProductVariant ProductVariantKeyReference `json:"productVariant"`
	// The Product in which the Product Variant containing this Price is contained. Maps to `ProductVariant.product`.
	// The Reference to the [Product](/../api/projects/products#product) with which the Price is associated.
	// If referenced Product does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Product is created.
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

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PriceImport) MarshalJSON() ([]byte, error) {
	type Alias PriceImport
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

	if target["tiers"] == nil {
		delete(target, "tiers")
	}

	return json.Marshal(target)
}
