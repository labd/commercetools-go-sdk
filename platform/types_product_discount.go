package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type ProductDiscount struct {
	// Unique identifier of the ProductDiscount.
	ID string `json:"id"`
	// Current version of the ProductDiscount.
	Version int `json:"version"`
	// Date and time (UTC) the ProductDiscount was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductDiscount was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the ProductDiscount.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the ProductDiscount.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the ProductDiscount.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the ProductDiscount.
	Key *string `json:"key,omitempty"`
	// Description of the ProductDiscount.
	Description *LocalizedString `json:"description,omitempty"`
	// Type of Discount and its corresponding value.
	Value ProductDiscountValue `json:"value"`
	// Valid [ProductDiscount predicate](/../api/projects/predicates#productdiscount-predicates).
	Predicate string `json:"predicate"`
	// Unique decimal value between 0 and 1 (stored as String literal) defining the order of Product Discounts to apply in case more than one is applicable and active.
	// A Product Discount with a higher value is prioritized.
	SortOrder string `json:"sortOrder"`
	// If `true` the Product Discount is applied to Products matching the `predicate`.
	IsActive bool `json:"isActive"`
	// References of all the resources that are addressed in the `predicate`.
	References []Reference `json:"references"`
	// Date and time (UTC) from which the Discount is effective.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount is effective.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscount) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}
	for i := range obj.References {
		var err error
		obj.References[i], err = mapDiscriminatorReference(obj.References[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ProductDiscountDraft struct {
	// Name of the ProductDiscount.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier for the ProductDiscount.
	Key *string `json:"key,omitempty"`
	// Description of the ProductDiscount.
	Description *LocalizedString `json:"description,omitempty"`
	// Type of Discount and its corresponding value.
	Value ProductDiscountValueDraft `json:"value"`
	// Valid [ProductDiscount predicate](/../api/projects/predicates#productdiscount-predicates).
	Predicate string `json:"predicate"`
	// Decimal value between 0 and 1 (passed as String literal) that defines the order of ProductDiscounts to apply in case more than one is applicable and active. A ProductDiscount with a higher `sortOrder` is prioritized.
	// The value must be **unique** among all ProductDiscounts in the [Project](ctp:api:type:Project).
	SortOrder string `json:"sortOrder"`
	// Set to `true` to activate the ProductDiscount, set to `false` to deactivate it (even though the `predicate` matches).
	IsActive bool `json:"isActive"`
	// Date and time (UTC) from which the Discount is effective.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount is effective.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

type ProductDiscountMatchQuery struct {
	// ID of the specified Product.
	ProductId string `json:"productId"`
	// ID of the specified Product Variant.
	VariantId int `json:"variantId"`
	// Controls which [projected representation](/../api/projects/productProjections#current--staged) is applied for the query.
	// Set to `true` for the `staged` Product Projection of the specified Product Variant, set to `false` for the `current` one.
	Staged bool `json:"staged"`
	// Specified Price of the specified Product Variant.
	Price QueryPrice `json:"price"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ProductDiscount](ctp:api:type:ProductDiscount).
*
 */
type ProductDiscountPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [ProductDiscounts](ctp:api:type:ProductDiscount) matching the query.
	Results []ProductDiscount `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductDiscount](ctp:api:type:ProductDiscount).
*
 */
type ProductDiscountReference struct {
	// Unique identifier of the referenced [ProductDiscount](ctp:api:type:ProductDiscount).
	ID string `json:"id"`
	// Contains the representation of the expanded ProductDiscount. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ProductDiscounts.
	Obj *ProductDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ProductDiscount](ctp:api:type:ProductDiscount). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ProductDiscountResourceIdentifier struct {
	// Unique identifier of the referenced [ProductDiscount](ctp:api:type:ProductDiscount). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ProductDiscount](ctp:api:type:ProductDiscount). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-discount", Alias: (*Alias)(&obj)})
}

type ProductDiscountUpdate struct {
	// Expected version of the ProductDiscount on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the ProductDiscount.
	Actions []ProductDiscountUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorProductDiscountUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ProductDiscountUpdateAction interface{}

func mapDiscriminatorProductDiscountUpdateAction(input interface{}) (ProductDiscountUpdateAction, error) {
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
	case "changeIsActive":
		obj := ProductDiscountChangeIsActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ProductDiscountChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changePredicate":
		obj := ProductDiscountChangePredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSortOrder":
		obj := ProductDiscountChangeSortOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeValue":
		obj := ProductDiscountChangeValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Value != nil {
			var err error
			obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "setDescription":
		obj := ProductDiscountSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ProductDiscountSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := ProductDiscountSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := ProductDiscountSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := ProductDiscountSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ProductDiscountValue interface{}

func mapDiscriminatorProductDiscountValue(input interface{}) (ProductDiscountValue, error) {
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
	case "absolute":
		obj := ProductDiscountValueAbsolute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "external":
		obj := ProductDiscountValueExternal{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := ProductDiscountValueRelative{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Discounts the Product's Price by a fixed amount, defined by the `money` field.
 */
type ProductDiscountValueAbsolute struct {
	// Money values in different currencies. An absolute [ProductDiscount](ctp:api:type:ProductDiscount) will only match a price if this array contains a value with the same currency. For example, if it contains 10€ and 15$, the matching € price will be decreased by 10€ and the matching $ price will be decreased by 15\$.
	Money []CentPrecisionMoney `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsolute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type ProductDiscountValueDraft interface{}

func mapDiscriminatorProductDiscountValueDraft(input interface{}) (ProductDiscountValueDraft, error) {
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
	case "absolute":
		obj := ProductDiscountValueAbsoluteDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "external":
		obj := ProductDiscountValueExternalDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := ProductDiscountValueRelativeDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Discounts the Product Price by a fixed amount, defined by the `money` field.
*
 */
type ProductDiscountValueAbsoluteDraft struct {
	// Money values in different currencies.
	// An absolute Product Discount will match a price only if the array contains a value with the same currency. For example, if it contains 10€ and 15$, the matching € price will be decreased by 10€ and the matching $ price will be decreased by 15$. If the array has multiple values of the same currency, the API returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	//
	// If the array is empty, the discount does not apply.
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueAbsoluteDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueAbsoluteDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

/**
*	Discounts the Product Price by allowing the client to explicitly [set a discounted value](ctp:api:type:ProductSetDiscountedPriceAction).
*	Used when setting discounts using an external service.
*
 */
type ProductDiscountValueExternal struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueExternal) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternal
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

/**
*	Discounts the Product Price by allowing the client to explicitly [set a discounted value](ctp:api:type:ProductSetDiscountedPriceAction).
*	Use this when setting discounts using an external service.
*
 */
type ProductDiscountValueExternalDraft struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueExternalDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueExternalDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "external", Alias: (*Alias)(&obj)})
}

/**
*	Discounts the product price by a percentage, defined by the `permyriad` field.
 */
type ProductDiscountValueRelative struct {
	// Fraction (per ten thousand) the price is reduced by. For example, `1000` will result in a 10% price reduction.
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelative
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

/**
*	Discounts the Product Price by a percentage, defined by the `permyriad` field.
*
 */
type ProductDiscountValueRelativeDraft struct {
	// Fraction (per ten thousand) the price is reduced by. For example, `1000` will result in a 10% price reduction.
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountValueRelativeDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountValueRelativeDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeIsActiveAction struct {
	// New value to set.
	// If set to `true`, the Discount will be applied to Product Prices.
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeNameAction struct {
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangePredicateAction struct {
	// New value to set. Must be a valid [ProductDiscount predicate](/../api/projects/predicates#productdiscount-predicates).
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangePredicateAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangePredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changePredicate", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeSortOrderAction struct {
	// New value to set.
	// Must not be empty.
	// The string value must be a number between `0` and `1`.
	// A Discount with a higher sortOrder is prioritized.
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

type ProductDiscountChangeValueAction struct {
	// New value to set. Must not be empty.
	Value ProductDiscountValueDraft `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias ProductDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorProductDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated discount values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidFromAndUntilAction struct {
	// Value to set.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated undiscounted values.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Value to set.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type ProductDiscountSetValidUntilAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	// Take [Eventual Consistency](/../api/general-concepts#eventual-consistency) into account for calculated undiscounted values.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias ProductDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
