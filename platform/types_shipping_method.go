package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type PriceFunction struct {
	// Currency code compliant to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	CurrencyCode string `json:"currencyCode"`
	// To calculate a Price based on the score, use `+`, `-`, `*` and parentheses. The score is inserted with `x`. The function returns the cent amount.
	//
	// For example, to charge $1.99 for a score of `1`, $3.99 for a score of `2`, \$5.99 for a score of `3` and onwards, the function is: `(200 * x) - 1)`. To charge $4.50, $6.00, and \$7.50 for express shipping, the function is: `(150 * x) + 300`.
	Function string `json:"function"`
}

type ShippingMethod struct {
	// Unique identifier of the ShippingMethod.
	ID string `json:"id"`
	// Current version of the ShippingMethod.
	Version int `json:"version"`
	// Date and time (UTC) the ShippingMethod was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ShippingMethod was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ShippingMethod.
	Key *string `json:"key,omitempty"`
	// Name of the ShippingMethod.
	Name string `json:"name"`
	// Localized name of the ShippingMethod.
	LocalizedName *LocalizedString `json:"localizedName,omitempty"`
	// Description of the ShippingMethod.
	Description *string `json:"description,omitempty"`
	// Localized description of the ShippingMethod.
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
	// [TaxCategory](ctp:api:type:TaxCategory) of all ZoneRates of the ShippingMethod.
	TaxCategory TaxCategoryReference `json:"taxCategory"`
	// Defines [ShippingRates](ctp:api:type:ShippingRate) (prices) for specific Zones.
	ZoneRates []ZoneRate `json:"zoneRates"`
	// If `true` this ShippingMethod is the [Project](ctp:api:type:Project)'s default ShippingMethod.
	IsDefault bool `json:"isDefault"`
	// Valid [Cart predicate](/projects/predicates#cart-predicates) to select a ShippingMethod for a Cart.
	Predicate *string `json:"predicate,omitempty"`
	// Custom Fields of the ShippingMethod.
	Custom *CustomFields `json:"custom,omitempty"`
}

type ShippingMethodDraft struct {
	// User-defined unique identifier for the ShippingMethod.
	Key *string `json:"key,omitempty"`
	// Name of the ShippingMethod.
	Name string `json:"name"`
	// Localized name of the ShippingMethod.
	LocalizedName *LocalizedString `json:"localizedName,omitempty"`
	// Description of the ShippingMethod.
	Description *string `json:"description,omitempty"`
	// Localized description of the ShippingMethod.
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
	// [TaxCategory](ctp:api:type:TaxCategory) for all ZoneRates of the ShippingMethod.
	TaxCategory TaxCategoryResourceIdentifier `json:"taxCategory"`
	// Defines [ShippingRates](ctp:api:type:ShippingRate) (prices) for specific zones.
	ZoneRates []ZoneRateDraft `json:"zoneRates"`
	// If `true` the ShippingMethod will be the [Project](ctp:api:type:Project)'s default ShippingMethod.
	IsDefault bool `json:"isDefault"`
	// Valid [Cart predicate](/projects/predicates#cart-predicates) to select a ShippingMethod for a Cart.
	Predicate *string `json:"predicate,omitempty"`
	// Custom Fields for the ShippingMethod.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [ShippingMethod](ctp:api:type:ShippingMethod).
*
 */
type ShippingMethodPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit *int `json:"limit,omitempty"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset *int `json:"offset,omitempty"`
	// [Shipping Methods](ctp:api:type:ShippingMethod) matching the query.
	Results []ShippingMethod `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ShippingMethod](ctp:api:type:ShippingMethod).
*
 */
type ShippingMethodReference struct {
	// Unique identifier of the referenced [ShippingMethod](ctp:api:type:ShippingMethod).
	ID string `json:"id"`
	// Contains the representation of the expanded ShippingMethod. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ShippingMethods.
	Obj *ShippingMethod `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodReference) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ShippingMethod](ctp:api:type:ShippingMethod).
*
 */
type ShippingMethodResourceIdentifier struct {
	// Unique identifier of the referenced [ShippingMethod](ctp:api:type:ShippingMethod). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ShippingMethod](ctp:api:type:ShippingMethod). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "shipping-method", Alias: (*Alias)(&obj)})
}

type ShippingMethodUpdate struct {
	// Expected version of the ShippingMethod on which the changes should be applied. If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the [ShippingMethod](/projects/shippingMethods#shippingmethod).
	Actions []ShippingMethodUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingMethodUpdate) UnmarshalJSON(data []byte) error {
	type Alias ShippingMethodUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorShippingMethodUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ShippingMethodUpdateAction interface{}

func mapDiscriminatorShippingMethodUpdateAction(input interface{}) (ShippingMethodUpdateAction, error) {
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
	case "addShippingRate":
		obj := ShippingMethodAddShippingRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addZone":
		obj := ShippingMethodAddZoneAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsDefault":
		obj := ShippingMethodChangeIsDefaultAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ShippingMethodChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTaxCategory":
		obj := ShippingMethodChangeTaxCategoryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeShippingRate":
		obj := ShippingMethodRemoveShippingRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeZone":
		obj := ShippingMethodRemoveZoneAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ShippingMethodSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ShippingMethodSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ShippingMethodSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ShippingMethodSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocalizedDescription":
		obj := ShippingMethodSetLocalizedDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocalizedName":
		obj := ShippingMethodSetLocalizedNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setPredicate":
		obj := ShippingMethodSetPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ShippingRate struct {
	// Currency amount of the ShippingRate.
	Price TypedMoney `json:"price"`
	// Shipping is free if the sum of the (Custom) Line Item Prices reaches the specified value.
	FreeAbove TypedMoney `json:"freeAbove,omitempty"`
	// `true` if the ShippingRate matches given [Cart](ctp:api:type:Cart) or [Location](ctp:api:type:Location).
	// Only appears in response to requests for [Get ShippingMethods for a Cart](#get-shippingmethods-for-a-cart) or
	// [Get ShippingMethods for a Location](#get-shippingmethods-for-a-location).
	IsMatching *bool `json:"isMatching,omitempty"`
	// Price tiers for the ShippingRate.
	Tiers []ShippingRatePriceTier `json:"tiers"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRate) UnmarshalJSON(data []byte) error {
	type Alias ShippingRate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Price != nil {
		var err error
		obj.Price, err = mapDiscriminatorTypedMoney(obj.Price)
		if err != nil {
			return err
		}
	}
	if obj.FreeAbove != nil {
		var err error
		obj.FreeAbove, err = mapDiscriminatorTypedMoney(obj.FreeAbove)
		if err != nil {
			return err
		}
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ShippingRateDraft struct {
	// Money value of the ShippingRate.
	Price Money `json:"price"`
	// Shipping is free if the sum of the (Custom) Line Item Prices reaches the specified value.
	FreeAbove *Money `json:"freeAbove,omitempty"`
	// Price tiers for the ShippingRate.
	Tiers []ShippingRatePriceTier `json:"tiers"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ShippingRateDraft) UnmarshalJSON(data []byte) error {
	type Alias ShippingRateDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Tiers {
		var err error
		obj.Tiers[i], err = mapDiscriminatorShippingRatePriceTier(obj.Tiers[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingRateDraft) MarshalJSON() ([]byte, error) {
	type Alias ShippingRateDraft
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

type ShippingRatePriceTier interface{}

func mapDiscriminatorShippingRatePriceTier(input interface{}) (ShippingRatePriceTier, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "CartClassification":
		obj := CartClassificationTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartScore":
		obj := CartScoreTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CartValue":
		obj := CartValueTier{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Used when the ShippingRate maps to an abstract Cart categorization expressed by strings (for example, `Light`, `Medium`, or `Heavy`).
*
 */
type CartClassificationTier struct {
	// `key` selected from the `values` of the [CartClassificationType](/projects/project#cartclassificationtype) configured in the Project.
	Value string `json:"value"`
	// Fixed shipping rate for the selected classification.
	Price Money `json:"price"`
	// Appears in response to [Get ShippingMethods for a Cart](#get-shippingmethods-for-a-cart) if the shipping rate matches the search query.
	IsMatching *bool `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartClassificationTier) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartClassification", Alias: (*Alias)(&obj)})
}

/**
*	Used when the ShippingRate maps to an abstract Cart categorization expressed by integers (such as shipping scores or weight ranges).
*	Either `price` or `priceFunction` is required.
*
 */
type CartScoreTier struct {
	// Abstract value for categorizing a Cart. The range starts at `0`. The default price covers `0`, tiers start at `1`. See [Using Tiered Shipping Rates](/../tutorials/shipping-rate) for details and examples.
	Score int `json:"score"`
	// Defines a fixed price for the `score`.
	Price *Money `json:"price,omitempty"`
	// Dynamically calculates a Price for a range of scores.
	PriceFunction *PriceFunction `json:"priceFunction,omitempty"`
	// Appears in response to [Get ShippingMethods for a Cart](#get-shippingmethods-for-a-cart) if the shipping rate matches the search query.
	IsMatching *bool `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartScoreTier) MarshalJSON() ([]byte, error) {
	type Alias CartScoreTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartScore", Alias: (*Alias)(&obj)})
}

/**
*	Used when the ShippingRate maps to the sum of [LineItem](ctp:api:type:LineItem) Prices.
*	The value of the Cart is used to select a tier.
*	If chosen, it is not possible to set a value for the `shippingRateInput` on the [Cart](ctp:api:type:Cart).
*	Tiers contain the `centAmount` (a value of `100` in the currency `USD` corresponds to `$ 1.00`), and start at `1`.'
*
 */
type CartValueTier struct {
	// Minimum total price of a Cart for which a shipping rate applies.
	MinimumCentAmount int `json:"minimumCentAmount"`
	// Fixed shipping rate Price for a CartValue.
	Price Money `json:"price"`
	// Appears in response to [Get ShippingMethods for a Cart](#get-shippingmethods-for-a-cart) if the shipping rate matches the search query.
	IsMatching *bool `json:"isMatching,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartValueTier) MarshalJSON() ([]byte, error) {
	type Alias CartValueTier
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CartValue", Alias: (*Alias)(&obj)})
}

type ShippingRateTierType string

const (
	ShippingRateTierTypeCartValue          ShippingRateTierType = "CartValue"
	ShippingRateTierTypeCartClassification ShippingRateTierType = "CartClassification"
	ShippingRateTierTypeCartScore          ShippingRateTierType = "CartScore"
)

/**
*	Defines shipping rates in different currencies for a specific [Zone](ctp:api:type:Zone).
*
 */
type ZoneRate struct {
	// [Zone](ctp:api:type:Zone) for which the shipping rates are valid.
	Zone ZoneReference `json:"zone"`
	// Shipping rates defined per currency.
	ShippingRates []ShippingRate `json:"shippingRates"`
}

type ZoneRateDraft struct {
	// Sets the [Zone](ctp:api:type:Zone) for which the shippng rates are valid.
	Zone ZoneResourceIdentifier `json:"zone"`
	// Shipping rates for the `currencies` configured in the [Project](ctp:api:type:Project). The array must not contain two ShippingRates with the same [CurrencyCode](ctp:api:type:CurrencyCode).
	ShippingRates []ShippingRateDraft `json:"shippingRates"`
}

type ShippingMethodAddShippingRateAction struct {
	// [Zone](ctp:api:type:Zone) to which the ShippingRate should be added.
	Zone ZoneResourceIdentifier `json:"zone"`
	// Value to add to `shippingRates`.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodAddShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodAddZoneAction struct {
	// Value to add to `zoneRates`.
	Zone ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodAddZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodAddZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addZone", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeIsDefaultAction struct {
	// Value to set. Only one ShippingMethod can be default in a [Project](ctp:api:type:Project).
	IsDefault bool `json:"isDefault"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodChangeIsDefaultAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeIsDefaultAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsDefault", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeNameAction struct {
	// Value to set. Must not be empty.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ShippingMethodChangeTaxCategoryAction struct {
	// Value to set.
	TaxCategory TaxCategoryResourceIdentifier `json:"taxCategory"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodChangeTaxCategoryAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodChangeTaxCategoryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTaxCategory", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveShippingRateAction struct {
	// [Zone](ctp:api:type:Zone) from which the ShippingRate should be removed.
	Zone ZoneResourceIdentifier `json:"zone"`
	// Value to remove from `shippingRates`.
	ShippingRate ShippingRateDraft `json:"shippingRate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodRemoveShippingRateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveShippingRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeShippingRate", Alias: (*Alias)(&obj)})
}

type ShippingMethodRemoveZoneAction struct {
	// Value to remove from `zoneRates`.
	Zone ZoneResourceIdentifier `json:"zone"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodRemoveZoneAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodRemoveZoneAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeZone", Alias: (*Alias)(&obj)})
}

/**
*	This action sets, overwrites, or removes any existing [Custom Field](/projects/custom-fields) for an existing ShippingMethod.
*
 */
type ShippingMethodSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the ShippingMethod with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the ShippingMethod.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the ShippingMethod.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetLocalizedDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	LocalizedDescription *LocalizedString `json:"localizedDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetLocalizedDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedDescription", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetLocalizedNameAction struct {
	// Value to set. If empty, any existing value will be removed.
	LocalizedName *LocalizedString `json:"localizedName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetLocalizedNameAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetLocalizedNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocalizedName", Alias: (*Alias)(&obj)})
}

type ShippingMethodSetPredicateAction struct {
	// A valid [Cart predicate](/projects/predicates#cart-predicates). If `predicate` is absent or `null`, it is removed if it exists.
	Predicate *string `json:"predicate,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ShippingMethodSetPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ShippingMethodSetPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setPredicate", Alias: (*Alias)(&obj)})
}
