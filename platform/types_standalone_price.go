package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type StagedPriceDraft struct {
	// Money value for the StagedPriceDraft.
	Value TypedMoneyDraft `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedPriceDraft) UnmarshalJSON(data []byte) error {
	type Alias StagedPriceDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorTypedMoneyDraft(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Staged changes on a Standalone Price.
*	To update the `value` property of a Staged Standalone Price, use the [Change Value](ctp:api:type:StandalonePriceChangeValueAction) update action.
*	To apply all staged changes to the Standalone Price, use the [Apply Staged Changes](ctp:api:type:StandalonePriceApplyStagedChangesAction) update action.
*
 */
type StagedStandalonePrice struct {
	// Money value of the StagedStandalonePrice.
	Value TypedMoney `json:"value"`
	// Discounted price for the StagedStandalonePrice.
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StagedStandalonePrice) UnmarshalJSON(data []byte) error {
	type Alias StagedStandalonePrice
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

type StandalonePrice struct {
	// Unique identifier of the StandalonePrice.
	ID string `json:"id"`
	// Current version of the StandalonePrice.
	Version int `json:"version"`
	// Date and time (UTC) the StandalonePrice was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the StandalonePrice was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the StandalonePrice.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the StandalonePrice.
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
	// Date until the Price is valid. Standalone Prices that are no longer valid are not automatically deleted, but they can be [deleted](/../api/projects/standalone-prices#delete-standaloneprice) if necessary.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Price tiers if any are defined.
	//
	// If `discounted` is present, the tiered Price is ignored for a Product Variant.
	Tiers []PriceTier `json:"tiers"`
	// Set if a matching [ProductDiscount](ctp:api:type:ProductDiscount) exists. If set, the API uses the `discounted` value for the [Line Item price selection](/../api/pricing-and-discounts-overview#line-item-price-selection).
	// When a [relative discount](/../api/projects/productDiscounts#productdiscountvaluerelative) is applied and the fraction part of the `discounted` price is 0.5, the discounted price is rounded in favor of the customer with the [half down rounding](https://en.wikipedia.org/wiki/Rounding#Round_half_down).
	Discounted *DiscountedPrice `json:"discounted,omitempty"`
	// Custom Fields for the StandalonePrice.
	Custom *CustomFields `json:"custom,omitempty"`
	// Staged changes of the StandalonePrice. Only present if the StandalonePrice has some changes staged.
	Staged *StagedStandalonePrice `json:"staged,omitempty"`
	// If set to `true`, the StandalonePrice is considered during [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	// If set to `false`, the StandalonePrice is not considered during [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection) and any associated Line Items in a Cart cannot be ordered.
	Active bool `json:"active"`
	// [RecurrencePolicy](ctp:api:type:RecurrencePolicy) for which this Price is valid.
	RecurrencePolicy *RecurrencePolicyReference `json:"recurrencePolicy,omitempty"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

type StandalonePriceDraft struct {
	// User-defined unique identifier for the StandalonePrice.
	Key *string `json:"key,omitempty"`
	// Specifies to which [ProductVariant](ctp:api:type:ProductVariant) the API associates this Price.
	// It is not validated to exist in product variants.
	Sku string `json:"sku"`
	// Sets the money value of this Price.
	//
	// To set the money value in high precision, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft).
	Value Money `json:"value"`
	// Sets the country for which this Price is valid.
	Country *string `json:"country,omitempty"`
	// Sets the [CustomerGroup](ctp:api:type:CustomerGroup) for which this Price is valid.
	CustomerGroup *CustomerGroupResourceIdentifier `json:"customerGroup,omitempty"`
	// Sets the product distribution [Channel](ctp:api:type:Channel) for which this Price is valid.
	Channel *ChannelResourceIdentifier `json:"channel,omitempty"`
	// Sets the date from which the Price is valid. Must be at least 1 ms earlier than `validUntil`.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Sets the date until the Price is valid. Must be at least 1 ms later than `validFrom`. Standalone Prices that are no longer valid are not automatically deleted, but they can be [deleted](/../api/projects/standalone-prices#delete-standaloneprice) if necessary.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Sets price tiers.
	//
	// If `discounted` is set, the tiered Price is ignored for a Product Variant.
	Tiers []PriceTierDraft `json:"tiers"`
	// Sets a discounted price for this Price that is different from the base price with `value`.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
	// Custom Fields for the StandalonePrice.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// [RecurrencePolicy](ctp:api:type:RecurrencePolicy) for which this Price is valid.
	RecurrencePolicy *RecurrencePolicyResourceIdentifier `json:"recurrencePolicy,omitempty"`
	// Staged changes for the StandalonePrice.
	Staged *StagedPriceDraft `json:"staged,omitempty"`
	// Set to `false`, if the StandalonePrice should not be considered during [Product price selection](/../api/pricing-and-discounts-overview#product-price-selection).
	Active *bool `json:"active,omitempty"`
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tiers"] == nil {
		delete(raw, "tiers")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [StandalonePrice](ctp:api:type:StandalonePrice).
*
 */
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
*	[Reference](ctp:api:type:Reference) to a [StandalonePrice](ctp:api:type:StandalonePrice).
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
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [StandalonePrice](ctp:api:type:StandalonePrice).
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
	// Expected version of the StandalonePrice on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
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
	case "addPriceTier":
		obj := StandalonePriceAddPriceTierAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "applyStagedChanges":
		obj := StandalonePriceApplyStagedChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeActive":
		obj := StandalonePriceChangeActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeValue":
		obj := StandalonePriceChangeValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removePriceTier":
		obj := StandalonePriceRemovePriceTierAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeStagedChanges":
		obj := StandalonePriceRemoveStagedChangesAction{}
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
	case "setDiscountedPrice":
		obj := StandalonePriceSetDiscountedPriceAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := StandalonePriceSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPriceTiers":
		obj := StandalonePriceSetPriceTiersAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := StandalonePriceSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := StandalonePriceSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := StandalonePriceSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Adding a [PriceTier](ctp:api:type:PriceTier) to a [StandalonePrice](ctp:api:type:StandalonePrice) produces the [Standalone Price Tier Added](ctp:api:type:StandalonePriceTierAddedMessage) Message.
*
 */
type StandalonePriceAddPriceTierAction struct {
	// The [PriceTier](ctp:api:type:PriceTier) to be added to the `tiers` field of the [StandalonePrice](ctp:api:type:StandalonePrice).
	// The action returns an [InvalidField](ctp:api:type:InvalidFieldError) error in the following cases:
	//
	// * Trying to add a PriceTier with `minimumQuantity` < `2`.
	// * Trying to add a PriceTier with `minimumQuantity` that already exists for the StandalonePrice.
	Tier PriceTierDraft `json:"tier"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceAddPriceTierAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceAddPriceTierAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addPriceTier", Alias: (*Alias)(&obj)})
}

/**
*	Applies all staged changes to the StandalonePrice by overwriting all current values with the values in the [StagedStandalonePrice](ctp:api:type:StagedStandalonePrice). After successfully applied, the [StagedStandalonePrice](ctp:api:type:StagedStandalonePrice) will be removed from the StandalonePrice. An `applyStagedChanges` update action on a StandalonePrice that does not contain any staged changes will return a `400 Bad Request` error. Applying staged changes successfully will produce the [StandalonePriceStagedChangesApplied](ctp:api:type:StandalonePriceStagedChangesAppliedMessage) Message.
*
 */
type StandalonePriceApplyStagedChangesAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceApplyStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceApplyStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "applyStagedChanges", Alias: (*Alias)(&obj)})
}

/**
*	Updating the value of a [StandalonePrice](ctp:api:type:StandalonePrice) produces the [StandalonePriceActiveChanged](ctp:api:type:StandalonePriceActiveChangedMessage) Message.
*
 */
type StandalonePriceChangeActiveAction struct {
	// New value to set for the `active` field of the [StandalonePrice](ctp:api:type:StandalonePrice).
	Active bool `json:"active"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceChangeActiveAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceChangeActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeActive", Alias: (*Alias)(&obj)})
}

/**
*	Updating the value of a [StandalonePrice](ctp:api:type:StandalonePrice) produces the [StandalonePriceValueChangedMessage](ctp:api:type:StandalonePriceValueChangedMessage).
*
 */
type StandalonePriceChangeValueAction struct {
	// New value to set. Must not be empty.
	//
	// To set the money value in high precision, use [HighPrecisionMoneyDraft](ctp:api:type:HighPrecisionMoneyDraft).
	Value Money `json:"value"`
	// If set to `true` the update action applies to the [StagedStandalonePrice](ctp:api:type:StagedStandalonePrice). If set to `false`, the update action applies to the current [StandalonePrice](ctp:api:type:StandalonePrice).
	Staged *bool `json:"staged,omitempty"`
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

/**
*	Removing a [PriceTier](ctp:api:type:PriceTier) from a [StandalonePrice](ctp:api:type:StandalonePrice) produces the [Standalone Price Tier Removed](ctp:api:type:StandalonePriceTierRemovedMessage) Message.
*
 */
type StandalonePriceRemovePriceTierAction struct {
	// The `minimumQuantity` of the [PriceTier](ctp:api:type:PriceTier) to be removed from the `tiers` field of the [StandalonePrice](ctp:api:type:StandalonePrice).
	TierMinimumQuantity int `json:"tierMinimumQuantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceRemovePriceTierAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceRemovePriceTierAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removePriceTier", Alias: (*Alias)(&obj)})
}

/**
*	Removes all staged changes from the StandalonePrice.
*	Removing staged changes successfully produces the [StandalonePriceStagedChangesRemoved](ctp:api:type:StandalonePriceStagedChangesRemovedMessage) Message.
*
 */
type StandalonePriceRemoveStagedChangesAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceRemoveStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceRemoveStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStagedChanges", Alias: (*Alias)(&obj)})
}

type StandalonePriceSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
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

/**
*	Discounts a Standalone Price of a Product Variant on a published [Product](ctp:api:type:Product).
*	If the Product Variant does not exist or if it exists only in the staged representation of a Product, an [InvalidOperationError](ctp:api:type:InvalidOperationError) error is returned.
*
*	Produces the [StandalonePriceExternalDiscountSet](ctp:api:type:StandalonePriceExternalDiscountSetMessage) Message.
*
 */
type StandalonePriceSetDiscountedPriceAction struct {
	// Value to set. If empty, any existing value will be removed.
	//
	// The referenced [ProductDiscount](ctp:api:type:ProductDiscount) must be of type external, active, and its predicate must match the referenced Price.
	Discounted *DiscountedPriceDraft `json:"discounted,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetDiscountedPriceAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetDiscountedPriceAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountedPrice", Alias: (*Alias)(&obj)})
}

/**
*	Sets the key on a Standalone Price. Produces the [StandalonePriceKeySet](ctp:api:type:StandalonePriceKeySetMessage) Message.
*
 */
type StandalonePriceSetKeyAction struct {
	// Value to set. Must be unique. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Sets all [PriceTiers](ctp:api:type:PriceTier) for a [StandalonePrice](ctp:api:type:StandalonePrice) in one action, produces the [Standalone Price Tiers Set](ctp:api:type:StandalonePriceTiersSetMessage) Message.
*
 */
type StandalonePriceSetPriceTiersAction struct {
	// Value to set. If empty, any existing value will be removed.
	// The `minimumQuantity` of the PriceTiers must be unique and greater than `1`, otherwise an [InvalidField](ctp:api:type:InvalidFieldError) error is returned.
	Tiers []PriceTierDraft `json:"tiers"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetPriceTiersAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetPriceTiersAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPriceTiers", Alias: (*Alias)(&obj)})
}

/**
*	Updating the `validFrom` value generates the [StandalonePriceValidFromSet](ctp:api:type:StandalonePriceValidFromSetMessage) Message.
*
*	As the validity dates are part of the price scope and are not allowed to overlap, this update might return the [DuplicateStandalonePriceScope](ctp:api:type:DuplicateStandalonePriceScopeError) and [OverlappingStandalonePriceValidity](ctp:api:type:OverlappingStandalonePriceValidityError) errors, respectively. A Price without validity period does not conflict with a Price defined for a time period.
*
 */
type StandalonePriceSetValidFromAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

/**
*	Updating the `validFrom` and `validUntil` values generates the [StandalonePriceValidFromAndUntilSet](ctp:api:type:StandalonePriceValidFromAndUntilSetMessage) Message.
*
*	As the validity dates are part of the price scope and are not allowed to overlap, this update might return the [DuplicateStandalonePriceScope](ctp:api:type:DuplicateStandalonePriceScopeError) and [OverlappingStandalonePriceValidity](ctp:api:type:OverlappingStandalonePriceValidityError) errors, respectively. A Price without validity period does not conflict with a Price defined for a time period.
*
 */
type StandalonePriceSetValidFromAndUntilAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Value to set.
	// If empty, any existing value is removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

/**
*	Updating the `validUntil` value generates the [StandalonePriceValidUntilSet](ctp:api:type:StandalonePriceValidUntilSetMessage) Message.
*
*	As the validity dates are part of the price scope and are not allowed to overlap, this update might return the [DuplicateStandalonePriceScope](ctp:api:type:DuplicateStandalonePriceScopeError) and [OverlappingStandalonePriceValidity](ctp:api:type:OverlappingStandalonePriceValidityError) errors, respectively. A Price without validity period does not conflict with a Price defined for a time period.
*
 */
type StandalonePriceSetValidUntilAction struct {
	// Value to set.
	// If empty, any existing value is removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandalonePriceSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias StandalonePriceSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
