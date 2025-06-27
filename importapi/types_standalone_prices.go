package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	Represents the data used to import a StandalonePrice. Once imported, this data is persisted as a [StandalonePrice](ctp:api:type:StandalonePrice)) in the Project.
*
 */
type StandalonePriceImport struct {
	// User-defined unique identifier for the StandalonePrice. If a [StandalonePrice](ctp:api:type:StandalonePrice)) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `StandalonePrice.sku`. This value is not validated to exist in Product Variants.
	Sku string `json:"sku"`
	// Maps to `StandalonePrice.value`.
	Value TypedMoney `json:"value"`
	// Maps to `StandalonePrice.country`. This value cannot be updated. Attempting to update this value will result in an [InvalidFieldsUpdate](/import-export/error#invalidfieldsupdateerror) error.
	Country *string `json:"country,omitempty"`
	// Maps to `StandalonePrice.customerGroup`. If the referenced [CustomerGroup](ctp:api:type:CustomerGroup) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced CustomerGroup is created.
	//
	// This value cannot be updated. Attempting to update this value will result in an [InvalidFieldsUpdate](/import-export/error#invalidfieldsupdateerror) error.
	CustomerGroup *CustomerGroupKeyReference `json:"customerGroup,omitempty"`
	// Maps to `StandalonePrice.channel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	//
	// This value cannot be updated. Attempting to update this value will result in an [InvalidFieldsUpdate](/import-export/error#invalidfieldsupdateerror) error.
	Channel *ChannelKeyReference `json:"channel,omitempty"`
	// Maps to `StandalonePrice.validFrom`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Maps to `StandalonePrice.validUntil`.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Maps to `StandalonePrice.tiers`.
	Tiers []PriceTier `json:"tiers"`
	// Sets a discounted price for this Price that is different from the base price with value.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Maps to `StandalonePrice.custom`.
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
