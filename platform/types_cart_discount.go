package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type CartDiscount struct {
	// Unique identifier of the CartDiscount.
	ID string `json:"id"`
	// Current version of the CartDiscount.
	Version int `json:"version"`
	// Date and time (UTC) the CartDiscount was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the CartDiscount was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the CartDiscount.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the CartDiscount.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the CartDiscount.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the CartDiscount.
	Key *string `json:"key,omitempty"`
	// Description of the CartDiscount.
	Description *LocalizedString `json:"description,omitempty"`
	// Effect of the CartDiscount on the `target`.
	Value CartDiscountValue `json:"value"`
	// Valid [Cart Predicate](/../api/projects/predicates#cart-predicates).
	CartPredicate string `json:"cartPredicate"`
	// Segment of the Cart that is discounted.
	//
	// Empty, if the `value` is `giftLineItem`.
	Target CartDiscountTarget `json:"target,omitempty"`
	// Value between `0` and `1` that determines the order in which the CartDiscounts are applied; a CartDiscount with a higher value is prioritized.
	//
	// It is unique among all CartDiscounts and DiscountGroups.
	//
	// If the CartDiscount is part of a DiscountGroup, it uses the sort order of the DiscountGroup.
	SortOrder string `json:"sortOrder"`
	// - If a value exists, the Cart Discount applies on [Carts](ctp:api:type:Cart) having a [Store](ctp:api:type:Store) matching any Store defined for this field.
	// - If empty, the Cart Discount applies on all [Carts](ctp:api:type:Cart), irrespective of a Store.
	Stores []StoreKeyReference `json:"stores"`
	// Indicates if the CartDiscount is active and can be applied to the Cart.
	IsActive bool `json:"isActive"`
	// Date and time (UTC) from which the Discount is effective.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount is effective.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Indicates if the Discount is used in connection with a [DiscountCode](ctp:api:type:DiscountCode).
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
	// References of all resources that are addressed in the predicate.
	// The API generates this array from the predicate.
	References []Reference `json:"references"`
	// Indicates whether the application of the CartDiscount causes other discounts to be ignored.
	StackingMode StackingMode `json:"stackingMode"`
	// Custom Fields of the CartDiscount.
	Custom *CustomFields `json:"custom,omitempty"`
	// Reference to a DiscountGroup that the CartDiscount belongs to.
	DiscountGroup *DiscountGroupReference `json:"discountGroup,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscount) UnmarshalJSON(data []byte) error {
	type Alias CartDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorCartDiscountValue(obj.Value)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
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

type CartDiscountDraft struct {
	// Name of the CartDiscount.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier for the CartDiscount.
	Key *string `json:"key,omitempty"`
	// Description of the CartDiscount.
	Description *LocalizedString `json:"description,omitempty"`
	// Effect of the CartDiscount on the `target`.
	Value CartDiscountValueDraft `json:"value"`
	// Valid [Cart Predicate](/../api/projects/predicates#cart-predicates).
	CartPredicate string `json:"cartPredicate"`
	// Segment of the Cart that will be discounted.
	//
	// Must not be set if the `value` is `giftLineItem`.
	Target CartDiscountTarget `json:"target,omitempty"`
	// Value between `0` and `1` that determines the order in which the CartDiscounts will be applied; a CartDiscount with a higher value will be prioritized.
	//
	// It must be unique among all CartDiscounts and DiscountGroups.
	//
	// If the CartDiscount is part of a DiscountGroup, it will use the sort order of the DiscountGroup.
	SortOrder *string `json:"sortOrder,omitempty"`
	// - If defined, the Cart Discount applies on [Carts](ctp:api:type:Cart) having a [Store](ctp:api:type:Store) matching any Store defined for this field.
	// - If not defined, the Cart Discount applies on all Carts, irrespective of a Store.
	//
	// If the referenced Stores exceed the [limit](/../api/limits#cart-discounts-stores), a [MaxStoreReferencesReached](ctp:api:type:MaxStoreReferencesReachedError) error is returned.
	//
	// If the referenced Stores exceed the [limit](/../api/limits#cart-discounts) for Cart Discounts that do not require a Discount Code, a [StoreCartDiscountsLimitReached](ctp:api:type:StoreCartDiscountsLimitReachedError) error is returned.
	Stores []StoreResourceIdentifier `json:"stores"`
	// Only active Discounts can be applied to the Cart.
	// If the [limit](/../api/limits#cart-discounts) for active Cart Discounts is reached, a [MaxCartDiscountsReached](ctp:api:type:MaxCartDiscountsReachedError) error is returned.
	IsActive *bool `json:"isActive,omitempty"`
	// Date and time (UTC) from which the Discount is effective.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount is effective.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// States whether the Discount can only be used in a connection with a [DiscountCode](ctp:api:type:DiscountCode).
	RequiresDiscountCode *bool `json:"requiresDiscountCode,omitempty"`
	// Specifies whether the application of this discount causes the following discounts to be ignored.
	StackingMode *StackingMode `json:"stackingMode,omitempty"`
	// Custom Fields of the CartDiscount.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Reference to a DiscountGroup that the CartDiscount must belong to.
	DiscountGroup *DiscountGroupResourceIdentifier `json:"discountGroup,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorCartDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountDraft
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

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [CartDiscount](ctp:api:type:CartDiscount).
*
 */
type CartDiscountPagedQueryResponse struct {
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
	// [CartDiscounts](ctp:api:type:CartDiscount) matching the query.
	Results []CartDiscount `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [CartDiscount](ctp:api:type:CartDiscount).
*
 */
type CartDiscountReference struct {
	// Unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount).
	ID string `json:"id"`
	// Contains the representation of the expanded CartDiscount. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for CartDiscounts.
	Obj *CartDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CartDiscount](ctp:api:type:CartDiscount). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type CartDiscountResourceIdentifier struct {
	// Unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

type CartDiscountTarget interface{}

func mapDiscriminatorCartDiscountTarget(input interface{}) (CartDiscountTarget, error) {
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
	case "customLineItems":
		obj := CartDiscountCustomLineItemsTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "lineItems":
		obj := CartDiscountLineItemsTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "pattern":
		obj := CartDiscountPatternTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.TriggerPattern {
			var err error
			obj.TriggerPattern[i], err = mapDiscriminatorPatternComponent(obj.TriggerPattern[i])
			if err != nil {
				return nil, err
			}
		}
		for i := range obj.TargetPattern {
			var err error
			obj.TargetPattern[i], err = mapDiscriminatorPatternComponent(obj.TargetPattern[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "shipping":
		obj := CartDiscountShippingCostTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "totalPrice":
		obj := CartDiscountTotalPriceTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "multiBuyCustomLineItems":
		obj := MultiBuyCustomLineItemsTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "multiBuyLineItems":
		obj := MultiBuyLineItemsTarget{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Discount is applied to [CustomLineItems](ctp:api:type:CustomLineItem) matching the `predicate`.
*
 */
type CartDiscountCustomLineItemsTarget struct {
	// Valid [CustomLineItem target predicate](/../api/projects/predicates#customlineitem-field-identifiers).
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountCustomLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "customLineItems", Alias: (*Alias)(&obj)})
}

/**
*	Discount is applied to [LineItems](ctp:api:type:LineItem) matching the `predicate`.
*
 */
type CartDiscountLineItemsTarget struct {
	// Valid [LineItem target predicate](/../api/projects/predicates#lineitem-field-identifiers).
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lineItems", Alias: (*Alias)(&obj)})
}

/**
*	Pattern targets can be used to model Buy and Get discounts.
*
*	Unlike [CartDiscountLineItemsTarget](#cartdiscountlineitemstarget) and [CartDiscountCustomLineItemsTarget](#cartdiscountcustomlineitemstarget), it does not apply to a (Custom) Line Item as a whole, but to individual units of a (Custom) Line Item. The discounts can apply multiple times on the same cart, but each unit can be discounted only once.
*
 */
type CartDiscountPatternTarget struct {
	// Defines the set of units of (Custom) Line Items in a Cart that trigger a discount application.
	//
	// Based on the availability of matching units, the `triggerPattern` can match multiple times, limiting the number of maximum times the discount will be applied.
	// The units matched in the `triggerPattern` are excluded and not considered for the `targetPattern`.
	//
	// To further limit the discount application, set the `maxOccurrence`.
	TriggerPattern []PatternComponent `json:"triggerPattern"`
	// Defines the set of units of (Custom) Line Items in a Cart on which the Discount is applied.
	//
	// Based on the availability of matching units and the limits from the `triggerPattern` or `maxOccurrence`, the `targetPattern` can match multiple times.
	//
	// This array cannot be empty.
	TargetPattern []PatternComponent `json:"targetPattern"`
	// Maximum number of times the Discount can apply on a Cart.
	//
	// If empty or not set, the Discount will apply indefinitely.
	MaxOccurrence *int `json:"maxOccurrence,omitempty"`
	// Determines which of the matching units of (Custom) Line Items are discounted.
	SelectionMode SelectionMode `json:"selectionMode"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountPatternTarget) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountPatternTarget
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.TriggerPattern {
		var err error
		obj.TriggerPattern[i], err = mapDiscriminatorPatternComponent(obj.TriggerPattern[i])
		if err != nil {
			return err
		}
	}
	for i := range obj.TargetPattern {
		var err error
		obj.TargetPattern[i], err = mapDiscriminatorPatternComponent(obj.TargetPattern[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountPatternTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountPatternTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "pattern", Alias: (*Alias)(&obj)})
}

/**
*	Discount is applied to the shipping costs of the [Cart](ctp:api:type:Cart).
*
 */
type CartDiscountShippingCostTarget struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountShippingCostTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountShippingCostTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "shipping", Alias: (*Alias)(&obj)})
}

/**
*	Discount is applied to the total price of the [Cart](ctp:api:type:Cart).
*	The same percentage of discount applies on the [Cart](ctp:api:type:Cart) or [Order](ctp:api:type:Order) `taxedPrice` and `taxedShippingPrice`.
*
 */
type CartDiscountTotalPriceTarget struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountTotalPriceTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountTotalPriceTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "totalPrice", Alias: (*Alias)(&obj)})
}

type CartDiscountUpdate struct {
	// Expected version of the CartDiscount on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the CartDiscount.
	Actions []CartDiscountUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorCartDiscountUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type CartDiscountUpdateAction interface{}

func mapDiscriminatorCartDiscountUpdateAction(input interface{}) (CartDiscountUpdateAction, error) {
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
	case "addStore":
		obj := CartDiscountAddStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeCartPredicate":
		obj := CartDiscountChangeCartPredicateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeIsActive":
		obj := CartDiscountChangeIsActiveAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := CartDiscountChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeRequiresDiscountCode":
		obj := CartDiscountChangeRequiresDiscountCodeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSortOrder":
		obj := CartDiscountChangeSortOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeStackingMode":
		obj := CartDiscountChangeStackingModeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeTarget":
		obj := CartDiscountChangeTargetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Target != nil {
			var err error
			obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "changeValue":
		obj := CartDiscountChangeValueAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Value != nil {
			var err error
			obj.Value, err = mapDiscriminatorCartDiscountValueDraft(obj.Value)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "removeStore":
		obj := CartDiscountRemoveStoreAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CartDiscountSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CartDiscountSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := CartDiscountSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDiscountGroup":
		obj := CartDiscountSetDiscountGroupAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CartDiscountSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setStores":
		obj := CartDiscountSetStoresAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFrom":
		obj := CartDiscountSetValidFromAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidFromAndUntil":
		obj := CartDiscountSetValidFromAndUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setValidUntil":
		obj := CartDiscountSetValidUntilAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CartDiscountValue interface{}

func mapDiscriminatorCartDiscountValue(input interface{}) (CartDiscountValue, error) {
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
		obj := CartDiscountValueAbsolute{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "fixed":
		obj := CartDiscountValueFixed{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Money {
			var err error
			obj.Money[i], err = mapDiscriminatorTypedMoney(obj.Money[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "giftLineItem":
		obj := CartDiscountValueGiftLineItem{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := CartDiscountValueRelative{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Discounts the [CartDiscountTarget](ctp:api:type:CartDiscountTarget) by an absolute amount (not allowed for [MultiBuyLineItemsTarget](ctp:api:type:MultiBuyLineItemsTarget) and [MultiBuyCustomLineItemsTarget](ctp:api:type:MultiBuyCustomLineItemsTarget)).
*
 */
type CartDiscountValueAbsolute struct {
	// Cent precision money values in different currencies.
	Money []CentPrecisionMoney `json:"money"`
	// Determines how the discount is applied on [CartDiscountLineItemTarget](ctp:api:type:CartDiscountLineItemsTarget) and [CartDiscountCustomLineItemTarget](ctp:api:type:CartDiscountCustomLineItemsTarget).
	ApplicationMode *DiscountApplicationMode `json:"applicationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueAbsolute
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type CartDiscountValueDraft interface{}

func mapDiscriminatorCartDiscountValueDraft(input interface{}) (CartDiscountValueDraft, error) {
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
		obj := CartDiscountValueAbsoluteDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "fixed":
		obj := CartDiscountValueFixedDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Money {
			var err error
			obj.Money[i], err = mapDiscriminatorTypedMoneyDraft(obj.Money[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "giftLineItem":
		obj := CartDiscountValueGiftLineItemDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "relative":
		obj := CartDiscountValueRelativeDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CartDiscountValueAbsoluteDraft struct {
	// Money values in different currencies.
	// An absolute Cart Discount will match a price only if the array contains a value with the same currency. For example, if it contains 10€ and 15$, the matching € price will be decreased by 10€ and the matching $ price will be decreased by 15$.
	//
	// If the array is empty or has multiple values of the same currency, the API returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Money []Money `json:"money"`
	// Determines how the discount applies on [CartDiscountLineItemTarget](ctp:api:type:CartDiscountLineItemsTarget) and [CartDiscountCustomLineItemTarget](ctp:api:type:CartDiscountCustomLineItemsTarget).
	//
	// If not set, the default behavior is `ProportionateDistribution`.
	ApplicationMode *DiscountApplicationMode `json:"applicationMode,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueAbsoluteDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueAbsoluteDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [DiscountedLineItemPrice](ctp:api:type:DiscountedLineItemPrice) of the [CartDiscountLineItemsTarget](ctp:api:type:CartDiscountLineItemsTarget) or [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget) to the value specified in the `money` field, if it is lower than the current Line Item price for the same currency. If the Line Item price is already discounted to a price equal to or lower than the respective price in the `money` field, this Discount is not applied. If the `quantity` of the Line Item eligible for the Discount is greater than `1`, the fixed price discount is only applied to the Line Item portion for which the `money` value is lesser than their current price.
*
 */
type CartDiscountValueFixed struct {
	// Money values in [cent precision](ctp:api:type:CentPrecisionMoney) or [high precision](ctp:api:type:HighPrecisionMoney) of different currencies.
	Money []TypedMoney `json:"money"`
	// Indicates how the discount is applied on [CartDiscountLineItemTarget](ctp:api:type:CartDiscountLineItemsTarget) or [CartDiscountCustomLineItemTarget](ctp:api:type:CartDiscountCustomLineItemsTarget).
	//
	// For [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget), the mode can also be `ProportionateDistribution` or `EvenDistribution`.
	ApplicationMode *DiscountApplicationMode `json:"applicationMode,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountValueFixed) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountValueFixed
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Money {
		var err error
		obj.Money[i], err = mapDiscriminatorTypedMoney(obj.Money[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueFixed) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueFixed
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "fixed", Alias: (*Alias)(&obj)})
}

/**
*	Sets the [DiscountedLineItemPrice](ctp:api:type:DiscountedLineItemPrice) of the [CartDiscountLineItemsTarget](ctp:api:type:CartDiscountLineItemsTarget) or [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget) to the value specified in the `money` field, if it is lower than the current Line Item price for the same currency. If the Line Item price is already discounted to a price equal to or lower than the respective price in the `money` field, this Discount is not applied.
*
 */
type CartDiscountValueFixedDraft struct {
	// Money values provided either in [cent precision](ctp:api:type:Money) or [high precision](ctp:api:type:HighPrecisionMoneyDraft) for different currencies.
	// A fixed Cart Discount will match a price only if the array contains a value with the same currency. For example, if it contains 10€ and 15$, the matching € price will be discounted by 10€ and the matching $ price will be discounted to 15$.
	//
	// If the array is empty or has multiple values of the same currency, the API returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	Money []TypedMoneyDraft `json:"money"`
	// Determines how the discount applies on [CartDiscountLineItemTarget](ctp:api:type:CartDiscountLineItemsTarget) or [CartDiscountCustomLineItemTarget](ctp:api:type:CartDiscountCustomLineItemsTarget).
	//
	// For [CartDiscountPatternTarget](ctp:api:type:CartDiscountPatternTarget), you can also set the mode to `ProportionateDistribution` or `EvenDistribution`.
	ApplicationMode *DiscountApplicationMode `json:"applicationMode,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountValueFixedDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountValueFixedDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Money {
		var err error
		obj.Money[i], err = mapDiscriminatorTypedMoneyDraft(obj.Money[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueFixedDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueFixedDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "fixed", Alias: (*Alias)(&obj)})
}

type CartDiscountValueGiftLineItem struct {
	// Reference to a Product.
	//
	// A Gift Line Item can be present on a Cart even if the referenced Product is unpublished.
	Product ProductReference `json:"product"`
	// [ProductVariant](ctp:api:type:ProductVariant) of the Product.
	VariantId int `json:"variantId"`
	// Channel must have the [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) `InventorySupply`.
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Channel must have the [ChannelRoleEnum](ctp:api:type:ChannelRoleEnum) `ProductDistribution`.
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueGiftLineItem) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

/**
*	Can only be used in a [CartDiscountDraft](ctp:api:type:CartDiscountDraft) with no `target` specified.
*	Hence, this type can not be used in the [Change Value](ctp:api:type:CartDiscountChangeValueAction) update action.
*
 */
type CartDiscountValueGiftLineItemDraft struct {
	// ResourceIdentifier of a Product.
	//
	// A Gift Line Item is added to a Cart even if the referenced Product is unpublished.
	Product ProductResourceIdentifier `json:"product"`
	// [ProductVariant](ctp:api:type:ProductVariant) of the Product.
	VariantId int `json:"variantId"`
	// Channel must have the role `InventorySupply`.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Channel must have the role `ProductDistribution`.
	DistributionChannel *ChannelResourceIdentifier `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueGiftLineItemDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItemDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

/**
*	Discounts the [CartDiscountTarget](ctp:api:type:CartDiscountTarget) relative to its price.
*
 */
type CartDiscountValueRelative struct {
	// Fraction (per ten thousand) the price is reduced by. For example, `1000` will result in a 10% price reduction.
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelative
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type CartDiscountValueRelativeDraft struct {
	// Fraction (per ten thousand) the price is reduced by. For example, `1000` will result in a 10% price reduction.
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountValueRelativeDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelativeDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

/**
*	This mode determines how absolute Discounts are applied on Line Items or Custom Line Items.
*
 */
type DiscountApplicationMode string

const (
	DiscountApplicationModeProportionateDistribution DiscountApplicationMode = "ProportionateDistribution"
	DiscountApplicationModeEvenDistribution          DiscountApplicationMode = "EvenDistribution"
	DiscountApplicationModeIndividualApplication     DiscountApplicationMode = "IndividualApplication"
)

/**
*	This Discount target is similar to `MultiBuyLineItems`, but is applied on Custom Line Items instead of Line Items.
*
 */
type MultiBuyCustomLineItemsTarget struct {
	// Valid [CustomLineItems target predicate](/../api/projects/predicates#customlineitem-field-identifiers). The Discount will be applied to Custom Line Items that are matched by the predicate.
	Predicate string `json:"predicate"`
	// Number of Custom Line Items to be present in order to trigger an application of this Discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Number of Custom Line Items that are discounted per application of this Discount.
	// It must be less than or equal to the `triggerQuantity`.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times this Discount can be applied.
	// Do not set if the Discount should be applied an unlimited number of times.
	MaxOccurrence *int `json:"maxOccurrence,omitempty"`
	// Discounts particular Line Items only according to the SelectionMode.
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MultiBuyCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyCustomLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

type MultiBuyLineItemsTarget struct {
	// Valid [LineItem target predicate](/../api/projects/predicates#lineitem-field-identifiers). The Discount will be applied to Line Items that are matched by the predicate.
	Predicate string `json:"predicate"`
	// Number of Line Items to be present in order to trigger an application of this Discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Number of Line Items that are discounted per application of this Discount.
	// It must be less than or equal to the `triggerQuantity`.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times this Discount can be applied.
	// Do not set if the Discount should be applied an unlimited number of times.
	MaxOccurrence *int `json:"maxOccurrence,omitempty"`
	// Discounts particular Line Items only according to the SelectionMode.
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj MultiBuyLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyLineItems", Alias: (*Alias)(&obj)})
}

/**
*	The pattern component it used to define a set of units based on some criteria. The criteria depends on the type of component used.
*
 */
type PatternComponent interface{}

func mapDiscriminatorPatternComponent(input interface{}) (PatternComponent, error) {
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
	case "CountOnCustomLineItemUnits":
		obj := CountOnCustomLineItemUnits{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "CountOnLineItemUnits":
		obj := CountOnLineItemUnits{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CountOnCustomLineItemUnits struct {
	// Valid [CustomLineItem predicate](/../api/projects/predicates#customlineitem-field-identifiers) that determines the units participating in the Discount.
	Predicate string `json:"predicate"`
	// Minimum number of units of a Custom Line Item that match the predicate.
	MinCount *int `json:"minCount,omitempty"`
	// Maximum number of units of a Custom Line Item that match the predicate.
	// There might be more units matching the predicate, but they will not be participating to the resulting set.
	//
	// The value must be greater than or equal to `minCount`.
	// If not provided, the component will match all units that satisfy the predicate.
	MaxCount *int `json:"maxCount,omitempty"`
	// Number of units of a Custom Line Item to exclude on every application of the Discount.
	//
	// Set only when configuring the `targetPattern`.
	//
	// The units matched first (satisfying the pattern component) will be excluded from the resulting set.
	// The `minCount`and `maxCount` are considered only after the exclusion. Pattern components are matched only if any further units satisfying the pattern component exist.
	// For example, if 5 jeans are required but only 3 should be discounted, the `excludeCount` value must be 2.
	ExcludeCount *int `json:"excludeCount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CountOnCustomLineItemUnits) MarshalJSON() ([]byte, error) {
	type Alias CountOnCustomLineItemUnits
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CountOnCustomLineItemUnits", Alias: (*Alias)(&obj)})
}

type CountOnLineItemUnits struct {
	// Valid [LineItem predicate](/../api/projects/predicates#lineitem-field-identifiers) that determines the units participating in the Discount.
	Predicate string `json:"predicate"`
	// Minimum number of units of a Line Item that match the predicate.
	MinCount *int `json:"minCount,omitempty"`
	// Maximum number of units of a Line Item that match the predicate.
	// There might be more units matching the predicate, but they will not be participating to the resulting set.
	//
	// The value must be greater than or equal to `minCount`.
	// If not provided, the component will match all units that satisfy the predicate.
	MaxCount *int `json:"maxCount,omitempty"`
	// Number of units of a Line Item to exclude on every application of the Discount.
	//
	// Set only when configuring the `targetPattern`.
	//
	// The units matched first (satisfying the pattern component) will be excluded from the resulting set.
	// The `minCount`and `maxCount` are considered only after the exclusion. Pattern components are matched only if any further units satisfying the pattern component exist.
	// For example, if 5 jeans are required but only 3 should be discounted, the `excludeCount` value must be 2.
	ExcludeCount *int `json:"excludeCount,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CountOnLineItemUnits) MarshalJSON() ([]byte, error) {
	type Alias CountOnLineItemUnits
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "CountOnLineItemUnits", Alias: (*Alias)(&obj)})
}

/**
*	Defines which matching items are to be discounted.
*
 */
type SelectionMode string

const (
	SelectionModeCheapest      SelectionMode = "Cheapest"
	SelectionModeMostExpensive SelectionMode = "MostExpensive"
)

/**
*	Describes how the Cart Discount interacts with other Discounts.
*
 */
type StackingMode string

const (
	StackingModeStacking              StackingMode = "Stacking"
	StackingModeStopAfterThisDiscount StackingMode = "StopAfterThisDiscount"
)

/**
*	If a referenced Store does not exist, a [ReferencedResourceNotFound](ctp:api:type:ReferencedResourceNotFoundError) error is returned.
*
*	This action generates a [CartDiscountStoreAdded](ctp:api:type:CartDiscountStoreAddedMessage) Message.
*
 */
type CartDiscountAddStoreAction struct {
	// [Store](ctp:api:type:Store) to add.
	//
	// A failed update can return the following errors:
	//
	// - If the referenced Stores exceed the [limit](/../api/limits#cart-discounts-stores), a [MaxStoreReferencesReached](ctp:api:type:MaxStoreReferencesReachedError) error is returned.
	// - If the referenced Stores exceed the [limit](/../api/limits#cart-discounts) for Cart Discounts that do not require a Discount Code, a [StoreCartDiscountsLimitReached](ctp:api:type:StoreCartDiscountsLimitReachedError) error is returned.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountAddStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountAddStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addStore", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeCartPredicateAction struct {
	// New value to set.
	CartPredicate string `json:"cartPredicate"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartPredicate", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeIsActiveAction struct {
	// New value to set.
	// If set to `true`, the Discount will be applied to the Cart.
	//
	// If the limit for active Cart Discounts is reached, a [MaxCartDiscountsReached](ctp:api:type:MaxCartDiscountsReachedError) error is returned.
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeNameAction struct {
	// New value to set.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeRequiresDiscountCodeAction struct {
	// New value to set.
	// If set to `true`, the Discount can only be used in connection with a [DiscountCode](ctp:api:type:DiscountCode).
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeRequiresDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeRequiresDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeRequiresDiscountCode", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeSortOrderAction struct {
	// New value to set (between `0` and `1`).
	// A Discount with a higher sortOrder is prioritized.
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeStackingModeAction struct {
	// New value to set.
	StackingMode StackingMode `json:"stackingMode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeStackingModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeStackingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeStackingMode", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeTargetAction struct {
	// New value to set.
	Target CartDiscountTarget `json:"target"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountChangeTargetAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeTargetAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		var err error
		obj.Target, err = mapDiscriminatorCartDiscountTarget(obj.Target)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeTargetAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTarget", Alias: (*Alias)(&obj)})
}

/**
*	Changes the [CartDiscountValue](ctp:api:type:CartDiscountValue) for [relative](ctp:api:type:CartDiscountValueRelative), [absolute](ctp:api:type:CartDiscountValueAbsolute) and [fixed price](ctp:api:type:CartDiscountValueFixed) CartDiscounts.
*	Changing to [Gift Line Item](ctp:api:type:CartDiscountValueGiftLineItem) is not supported.
*
 */
type CartDiscountChangeValueAction struct {
	// New value to set.
	// When trying to set a [CartDiscountValueGiftLineItemDraft](ctp:api:type:CartDiscountValueGiftLineItemDraft) an [InvalidInput](ctp:api:type:InvalidInputError) error is returned.
	Value CartDiscountValueDraft `json:"value"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		var err error
		obj.Value, err = mapDiscriminatorCartDiscountValueDraft(obj.Value)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

/**
*	If a referenced Store does not exist, a [ReferencedResourceNotFound](ctp:api:type:ReferencedResourceNotFoundError) error is returned.
*
*	This action generates a [CartDiscountStoreRemoved](ctp:api:type:CartDiscountStoreRemovedMessage) Message.
*
 */
type CartDiscountRemoveStoreAction struct {
	// [Store](ctp:api:type:Store) to remove.
	Store StoreResourceIdentifier `json:"store"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountRemoveStoreAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountRemoveStoreAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStore", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the CartDiscount with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the CartDiscount.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the CartDiscount.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CartDiscountSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CartDiscountSetDiscountGroupAction struct {
	// Reference to a DiscountGroup that the Cart Discount must belong to.
	// If empty, any existing value will be removed.
	DiscountGroup *DiscountGroupResourceIdentifier `json:"discountGroup,omitempty"`
	// New value to set (between `0` and `1`) that determines the order in which the CartDiscounts are applied; a CartDiscount with a higher value is prioritized.
	//
	// Required if `discountGroup` is absent. If `discountGroup` is set, the CartDiscount will use the sort order of the DiscountGroup.
	SortOrder *string `json:"sortOrder,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetDiscountGroupAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetDiscountGroupAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDiscountGroup", Alias: (*Alias)(&obj)})
}

type CartDiscountSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	If a referenced Store does not exist, a [ReferencedResourceNotFound](ctp:api:type:ReferencedResourceNotFoundError) error is returned.
*
*	This action generates a [CartDiscountStoresSet](ctp:api:type:CartDiscountStoresSetMessage) Message.
*
 */
type CartDiscountSetStoresAction struct {
	// [Stores](ctp:api:type:Store) to set.
	// Overrides the current list of Stores.
	// If empty, any existing values will be removed.
	//
	// A failed update can return the following errors:
	//
	// - If the referenced Stores exceed the [limit](/../api/limits#cart-discounts-stores), a [MaxStoreReferencesReached](ctp:api:type:MaxStoreReferencesReachedError) error is returned.
	// - If the referenced Stores exceed the [limit](/../api/limits#cart-discounts) for Cart Discounts that do not require a Discount Code, a [StoreCartDiscountsLimitReached](ctp:api:type:StoreCartDiscountsLimitReachedError) error is returned.
	Stores []StoreResourceIdentifier `json:"stores"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetStoresAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetStoresAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setStores", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAndUntilAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Value to set.
	// If empty, any existing value will be removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidUntilAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
