package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type StandalonePrice struct {
	// Unique identifier of the StandalonePrice.
	ID string `json:"id"`
	// Current version of the StandalonePrice.
	Version int `json:"version"`
	// Date and time (UTC) the StandalonePrice was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the StandalonePrice was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the StandalonePrice.
	Key *string `json:"key,omitempty"`
	// SKU of the [ProductVariant](ctp:api:type:ProductVariant) to which this Price is associated.
	Sku string `json:"sku"`
	// Money value of this Price.
	Value TypedMoney `json:"value"`
	// Country for which this Price is valid.
	Country *string `json:"country,omitempty"`
	// [CustomerGroup](ctp:api:type:CustomerGroup) for which this Price is valid.
	CustomerGroup *CustomerGroupReference `json:"customerGroup,omitempty"`
	// Product distribution [Channel](ctp:api:type:Channel) for which this Price is valid.
	Channel *ChannelReference `json:"channel,omitempty"`
	// Date from which the Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date until the Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Price tiers if any are defined.
	Tiers []PriceTier `json:"tiers"`
	// Set if a matching [ProductDiscount](ctp:api:type:ProductDiscount) exists. If set, the API uses the `discounted` value for the [LineItem Price selection](/../api/projects/carts#lineitem-price-selection).
	// When a [relative discount](/../api/projects/productDiscounts#productdiscountvaluerelative) is applied and the fraction part of the `discounted` price is 0.5, the discounted price is rounded in favor of the customer with the [half down rounding](https://en.wikipedia.org/wiki/Rounding#Round_half_down).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Custom Fields for the StandalonePrice.
	Custom *CustomFields `json:"custom,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePrice) UnmarshalJSON(data []byte) error {
	type Alias StandalonePrice
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
func (obj StandalonePrice) MarshalJSON() ([]byte, error) {
	type Alias StandalonePrice
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

/**
*	Standalone Prices are defined with a scope consisting of `currency` and optionally `country`, `customerGroup`, and `channel` and/or a validity period (`validFrom` and/or `validTo`). For more information see [price selection](/../api/projects/products#price-selection).
*	Creating a Standalone Price is rejected if there already exists a Standalone Price for the same SKU with exactly the same price scope, or with overlapping validity periods within the same price scope. A Price without validity period does not conflict with a Price defined for a time period.
 */
type StandalonePriceDraft struct {
	// User-defined unique identifier for the StandalonePrice.
	Key *string `json:"key,omitempty"`
	// Specifies to which [ProductVariant](ctp:api:type:ProductVariant) the API associates this Price.
	// It is not validated to exist in product variants.
	Sku string `json:"sku"`
	// Sets the money value of this Price.
	Value Money `json:"value"`
	// Sets the country for which this Price is valid.
	Country *string `json:"country,omitempty"`
	// Sets the [CustomerGroup](ctp:api:type:CustomerGroup) for which this Price is valid.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// Sets the product distribution [Channel](ctp:api:type:Channel) for which this Price is valid.
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Sets the date from which the Price is valid.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Sets the date until the Price is valid.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Sets price tiers.
	Tiers []PriceTierDraft `json:"tiers"`
	// Sets a discounted price for this Price that is different from the base price with `value`.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
	// Custom Fields for the StandalonePrice.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceDraft) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceDraft
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

type StandalonePricePagedQueryResponse struct {
	// Number of requested results.
	Limit int `json:"limit"`
	// Offset supplied by the client or server default. It is the number of elements skipped, not a page number.
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [StandalonePrices](ctp:api:type:StandalonePrice) matching the query.
	Results []StandalonePrice `json:"results"`
}

/**
*	[Reference](/../api/types#reference) to a [StandalonePrice](ctp:api:type:StandalonePrice).
*
 */
type StandalonePriceReference struct {
	// Unique ID of the referenced resource.
	ID string `json:"id"`
	// Contains the representation of the expanded StandalonePrice. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for StandalonePrice.
	Obj *StandalonePrice `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceReference) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "standalone-price", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](/../api/types#resourceidentifier) to a [StandalonePrice](ctp:api:type:StandalonePrice).
*
 */
type StandalonePriceResourceIdentifier struct {
	// Unique identifier of the referenced resource. Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced resource. Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "standalone-price", Alias: (*Alias)(&obj)})
}

type StandalonePriceUpdate struct {
	// Expected version of the StandalonePrice on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the StandalonePrice.
	Actions []StandalonePriceUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StandalonePriceUpdate) UnmarshalJSON(data []byte) error {
	type Alias StandalonePriceUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStandalonePriceUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type StandalonePriceUpdateAction interface{}

func mapDiscriminatorStandalonePriceUpdateAction(input interface{}) (StandalonePriceUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "changeValue":
		obj := StandalonePriceChangeValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := StandalonePriceSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := StandalonePriceSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type StandalonePriceChangeValueAction struct {
	// New value to set. Must not be empty.
	Value Money `json:"value"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

type StandalonePriceSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type StandalonePriceSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the StandalonePrice with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the StandalonePrice.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the StandalonePrice.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}
