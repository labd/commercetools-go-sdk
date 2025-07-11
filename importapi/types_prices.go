package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

type SubRate struct {
	// Name of the SubRate.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["subRates"] == nil {
		delete(raw, "subRates")
	}

	return json.Marshal(raw)

}

/**
*	Represents the data used to import an [Embedded Price](/../api/pricing-and-discounts-overview#embedded-prices) . Once imported, this data is persisted as a [Price](/../api/types#price) in a Product Variant.
*
 */
type PriceImport struct {
	// User-defined unique identifier for the Embedded Price. If a [Price](/../api/types#price) with this `key` exists on the specified `productVariant`, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `Price.value`.
	Value TypedMoney `json:"value"`
	// Maps to `Price.county`.
	Country *string `json:"country,omitempty"`
	// Maps to `Price.validFrom`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Maps to `Price.validUntil`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Maps to `Price.customerGroup`. If the referenced [CustomerGroup](ctp:api:type:CustomerGroup) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CustomerGroup is created.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `Price.channel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Sets a discounted price from an external service.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Only the [Embedded Price](/types#price) updates will be published to `staged` and `current` projection.
	Publish *bool `json:"publish,omitempty"`
	// - Set to `false` to update both the [current and staged projections](/../api/projects/productProjections#current--staged) of the [Product](ctp:api:type:Product) with the new Price data.
	// - Leave empty or set to `true` to only update the staged projection.
	Staged *bool `json:"staged,omitempty"`
	// The tiered prices for this price.
	Tiers []PriceTier `json:"tiers"`
	// Maps to `Price.custom`.
	Custom *Custom `json:"custom,omitempty"`
	// The [ProductVariant](ctp:api:type:ProductVariant) which contains this Embedded Price. If the referenced ProductVariant does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced ProductVariant is created.
	ProductVariant ProductVariantKeyReference `json:"productVariant"`
	// The [Product](ctp:api:type:Product) which contains the `productVariant`. If the referenced Product does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Product is created.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}
