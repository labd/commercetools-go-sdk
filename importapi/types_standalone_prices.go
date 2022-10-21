package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	The data representation for a Standalone Price to be imported that is persisted as a [Standalone Price](/../api/projects/standalone-prices#standaloneprice) in the Project.
*
 */
type StandalonePriceImport struct {
	// User-defined unique identifier for the Standalone Price.
	Key string `json:"key"`
	// Specifies to which ProductVariant the API associates this Price. It is not validated to exist in product variants.
	Sku string `json:"sku"`
	// Sets the money value of this Price.
	Value TypedMoney `json:"value"`
	// Sets the country for which this Price is valid.
	Country *string `json:"country,omitempty"`
	// Sets the CustomerGroup for which this Price is valid.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Sets the product distribution Channel for which this Price is valid
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Sets the date from which the Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Sets the date until the Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Sets price tiers.
	Tiers []PriceTier `json:"tiers"`
	// Sets a discounted price for this Price that is different from the base price with value.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Custom Fields for the StandalonePrice.
	Custom *Custom `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceImport) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceImport
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
func (obj StandalonePriceImport) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceImport
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
