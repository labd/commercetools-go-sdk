package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	It is used to calculate the [taxPortions](/api/projects/carts#taxedprice) field in a Cart or Order.
 */
type SubRate struct {
	// Name of the SubRate.
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TaxCategory struct {
	// Unique identifier of the TaxCategory.
	ID string `json:"id"`
	// Current version of the TaxCategory.
	Version int `json:"version"`
	// Date and time (UTC) the TaxCategory was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the TaxCategory was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the TaxCategory.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the TaxCategory.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the TaxCategory.
	Name string `json:"name"`
	// Description of the TaxCategory.
	Description *string `json:"description,omitempty"`
	// Tax rates and subrates of states and countries. Each TaxRate in the array has a unique ID.
	Rates []TaxRate `json:"rates"`
	// User-defined unique identifier of the TaxCategory.
	Key *string `json:"key,omitempty"`
}

type TaxCategoryDraft struct {
	// Name of the TaxCategory.
	//
	// If the provided name is used by another TaxCategory in the Project, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	Name string `json:"name"`
	// Description of the TaxCategory.
	Description *string `json:"description,omitempty"`
	// Tax rates and subrates of states and countries.
	Rates []TaxRateDraft `json:"rates"`
	// User-defined unique identifier for the TaxCategory.
	//
	// If the value is used by another TaxCategory in the Project, a [DuplicateField](ctp:api:type:DuplicateFieldError) error is returned.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryDraft) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryDraft
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

	if raw["rates"] == nil {
		delete(raw, "rates")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/api/general-concepts#pagedqueryresult) with `results` containing an array of [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryPagedQueryResponse struct {
	// Number of [results requested](/api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/api/predicates/query), `total` is subject to a [limit](/api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [TaxCategories](ctp:api:type:TaxCategory) matching the query.
	Results []TaxCategory `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [TaxCategory](ctp:api:type:TaxCategory).
*
 */
type TaxCategoryReference struct {
	// Unique identifier of the referenced [TaxCategory](ctp:api:type:TaxCategory).
	ID string `json:"id"`
	// Contains the representation of the expanded TaxCategory. Only present in responses to requests with [Reference Expansion](/api/general-concepts#reference-expansion) for TaxCategories.
	Obj *TaxCategory `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [TaxCategory](ctp:api:type:TaxCategory). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](ctp:api:type:InvalidJsonInputError) error is returned.
*
 */
type TaxCategoryResourceIdentifier struct {
	// Unique identifier of the referenced [TaxCategory](ctp:api:type:TaxCategory). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [TaxCategory](ctp:api:type:TaxCategory). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

type TaxCategoryUpdate struct {
	// Expected version of the TaxCategory on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the TaxCategory.
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorTaxCategoryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type TaxCategoryUpdateAction interface{}

func mapDiscriminatorTaxCategoryUpdateAction(input interface{}) (TaxCategoryUpdateAction, error) {
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
	case "addTaxRate":
		obj := TaxCategoryAddTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := TaxCategoryChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeTaxRate":
		obj := TaxCategoryRemoveTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "replaceTaxRate":
		obj := TaxCategoryReplaceTaxRateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := TaxCategorySetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := TaxCategorySetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type TaxRate struct {
	// Present if the TaxRate is part of a [TaxCategory](ctp:api:type:TaxCategory).
	// Absent for external TaxRates in [LineItem](ctp:api:type:LineItem), [CustomLineItem](ctp:api:type:CustomLineItem), and [ShippingInfo](ctp:api:type:ShippingInfo).
	ID *string `json:"id,omitempty"`
	// User-defined identifier of the TaxRate. Unique within the TaxCategory containing it.
	// Present when set using [TaxRateDraft](ctp:api:type:TaxRateDraft). Not available for external TaxRates created using [ExternalTaxRateDraft](ctp:api:type:ExternalTaxRateDraft).
	Key *string `json:"key,omitempty"`
	// Name of the TaxRate.
	Name string `json:"name"`
	// Tax rate. If subrates are used, the amount is the sum of all rates in `subRates`.
	Amount float64 `json:"amount"`
	// Whether tax is included in [Embedded Prices](ctp:api:type:Price) or [Standalone Prices](ctp:api:type:StandalonePrice), and the `taxedPrice` is present on [LineItems](ctp:api:type:LineItem). In this case, the `totalNet` price on [TaxedPrice](ctp:api:type:TaxedPrice) includes the TaxRate.
	IncludedInPrice bool `json:"includedInPrice"`
	// Country in which the tax rate is applied in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	Country string `json:"country"`
	// State within the country, such as Texas in the United States.
	// The value is case-sensitive and must use the same casing as the `state` value in the Cart `shippingAddress`.
	//
	// A TaxRate whose `state` is omitted does **not** act as a wildcard: it only matches a Cart whose `shippingAddress` also has no `state` value. To apply the same rate across states, either define an individual TaxRate for each state, or use the `region` field or a Custom Field on the shipping address as described in [Address matching](/learning-model-your-business-structure/model-your-taxes/tax-categories-and-tax-rates#address-matching).
	State *string `json:"state,omitempty"`
	// Used when the total tax is a combination of multiple taxes (for example, local, state/provincial, and/or federal taxes). The total of all subrates must equal the TaxRate `amount`.
	// These subrates are used to calculate the `taxPortions` field of a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) and the `taxedPrice` field of [LineItems](ctp:api:type:LineItem), [CustomLineItems](ctp:api:type:CustomLineItem), and [ShippingInfos](ctp:api:type:ShippingInfo).
	SubRates []SubRate `json:"subRates"`
	// Determines which of the net price and the tax amount the `taxRoundingMode` of the Cart or Order is applied to, when this TaxRate is included in the price. Ignored if `includedInPrice` is `false`.
	//
	// Always returned by the API. Can be omitted when a TaxRate is supplied as input, such as in [OrderImportDraft](ctp:api:type:OrderImportDraft), and then defaults to `Net`.
	TaxRoundingTarget *TaxRoundingTarget `json:"taxRoundingTarget,omitempty"`
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

type TaxRateDraft struct {
	// Name of the TaxRate.
	Name string `json:"name"`
	// Tax rate.
	//
	// Either `amount` or `subRates` must be defined. If both are defined, the value of this field must be equal to the sum of the `subRates` amounts; otherwise, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	Amount *float64 `json:"amount,omitempty"`
	// Whether tax is included in [Embedded Prices](ctp:api:type:Price) or [Standalone Prices](ctp:api:type:StandalonePrice), and the `taxedPrice` is present on [LineItems](ctp:api:type:LineItem). In this case, the `totalNet` price on [TaxedPrice](ctp:api:type:TaxedPrice) includes the TaxRate.
	IncludedInPrice bool `json:"includedInPrice"`
	// Country in which the tax rate is applied in [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format.
	//
	// If the provided combination of `country` and `state` exists for the TaxCategory, a [DuplicateField](ctp:api:type:DuplicateFieldError) error is returned.
	Country string `json:"country"`
	// State within the country, such as Texas in the United States.
	// The value is case-sensitive and must use the same casing as the `state` value in the Cart `shippingAddress`. Empty strings are treated as if `state` was omitted.
	//
	// If `state` is omitted, the resulting TaxRate does **not** act as a wildcard: it only matches a Cart whose `shippingAddress` also has no `state` value. To apply the same rate across states, either define an individual TaxRate for each state, or use the `region` field or a Custom Field on the shipping address as described in [Address matching](/learning-model-your-business-structure/model-your-taxes/tax-categories-and-tax-rates#address-matching).
	//
	// If the provided combination of `country` and `state` exists for the TaxCategory, a [DuplicateField](ctp:api:type:DuplicateFieldError) error is returned.
	State *string `json:"state,omitempty"`
	// Used when the total tax is a combination of multiple taxes (for example, local, state/provincial, and/or federal taxes). If `amount` is defined, the total of all subrates must equal `amount`.
	// These subrates are used to calculate the `taxPortions` field of a [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) and the `taxedPrice` field of [LineItems](ctp:api:type:LineItem), [CustomLineItems](ctp:api:type:CustomLineItem), and [ShippingInfos](ctp:api:type:ShippingInfo).
	SubRates []SubRate `json:"subRates"`
	// User-defined identifier of the TaxRate.
	//
	// If the provided key is used by another TaxRate in the TaxCategory, a [DuplicateField](ctp:api:type:DuplicateFieldError) error is returned.
	Key *string `json:"key,omitempty"`
	// Determines whether the `taxRoundingMode` of the Cart or Order is applied to the net price or the tax amount when this TaxRate is included in the price. The field is ignored if `includedInPrice` is `false`.
	TaxRoundingTarget *TaxRoundingTarget `json:"taxRoundingTarget,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxRateDraft) MarshalJSON() ([]byte, error) {
	type Alias TaxRateDraft
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
*	For a [TaxRate](ctp:api:type:TaxRate) that is included in the price, this value determines which of the two derived amounts the [RoundingMode](ctp:api:type:RoundingMode) set by `taxRoundingMode` is applied to. The other amount is the exact difference from the gross price, so `totalNet` plus `totalTax` on [TaxedPrice](ctp:api:type:TaxedPrice) always equals `totalGross`.
*
*	This field has no effect if `includedInPrice` is `false`. In that case, the net price is supplied exactly, so rounding the gross price and rounding the tax amount give the same result.
*
*	For more information, see [Tax rounding target](/api/carts-orders-overview#tax-rounding-target).
*
 */
type TaxRoundingTarget string

const (
	TaxRoundingTargetNet TaxRoundingTarget = "Net"
	TaxRoundingTargetTax TaxRoundingTarget = "Tax"
)

type TaxCategoryAddTaxRateAction struct {
	// Value to append to the `rates` array.
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryAddTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryAddTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryChangeNameAction struct {
	// New value to set.
	//
	// If the provided name is used by another TaxCategory in the Project, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TaxCategoryRemoveTaxRateAction struct {
	// ID of the TaxRate to remove. Either `taxRateId` or `taxRateKey` must be provided.
	//
	// If the referenced TaxRate does not exist within the TaxCategory, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	TaxRateId *string `json:"taxRateId,omitempty"`
	// Key of the TaxRate to remove. Either `taxRateId` or `taxRateKey` must be provided.
	//
	// If the referenced TaxRate does not exist within the TaxCategory, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	TaxRateKey *string `json:"taxRateKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryReplaceTaxRateAction struct {
	// ID of the TaxRate to replace.
	// Either `taxRateId` or `taxRateKey` must be provided.
	//
	// If the referenced TaxRate does not exist within the TaxCategory, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	TaxRateId *string `json:"taxRateId,omitempty"`
	// Key of the TaxRate to replace.
	// Either `taxRateId` or `taxRateKey` must be provided.
	//
	// If the referenced TaxRate does not exist within the TaxCategory, an [InvalidOperation](ctp:api:type:InvalidOperationError) error is returned.
	TaxRateKey *string `json:"taxRateKey,omitempty"`
	// New TaxRate to replace with.
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategorySetDescriptionAction struct {
	// Value to set. If omitted, any existing value is removed.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type TaxCategorySetKeyAction struct {
	// Value to set. If omitted, any existing value is removed.
	//
	// If the value is used by another TaxCategory in the Project, a [DuplicateField](ctp:api:type:DuplicateFieldError) error is returned.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
