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
	// Date and time (UTC) for the CartDiscount was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) for the CartDiscount was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the CartDiscount.
	Name LocalizedString `json:"name"`
	// User-defined unique identifier of the CartDiscount.
	Key *string `json:"key,omitempty"`
	// Description of the CartDiscount.
	Description *LocalizedString `json:"description,omitempty"`
	// Effect of the CartDiscount.
	Value CartDiscountValue `json:"value"`
	// Valid [Cart Predicate](/../api/projects/predicates#cart-predicates).
	CartPredicate string `json:"cartPredicate"`
	// Sets a [CartDiscountTarget](ctp:api:type:CartDiscountTarget). Empty if `value` has type `giftLineItem`.
	Target CartDiscountTarget `json:"target,omitempty"`
	// Value between `0` and `1`.
	// All matching CartDiscounts are applied to a Cart in the order defined by this field.
	// A Discount with a higher sortOrder is prioritized.
	// The sort order is unambiguous among all CartDiscounts.
	SortOrder string `json:"sortOrder"`
	// Indicates if the CartDiscount is active and can be applied to the Cart.
	IsActive bool `json:"isActive"`
	// Date and time (UTC) from which the Discount is effective.
	ValidFrom *time.Time `json:"validFrom,omitempty"`
	// Date and time (UTC) until which the Discount is effective.
	ValidUntil *time.Time `json:"validUntil,omitempty"`
	// Indicates if the Discount can be used in connection with a [DiscountCode](ctp:api:type:DiscountCode).
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
	// References of all resources that are addressed in the predicate.
	// The API generates this array from the predicate.
	References []Reference `json:"references"`
	// Indicates whether the application of the CartDiscount causes other discounts to be ignored.
	StackingMode StackingMode `json:"stackingMode"`
	// Custom Fields of the CartDiscount.
	Custom *CustomFields `json:"custom,omitempty"`
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
	// Effect of the CartDiscount.
	// For a target, relative or absolute discount values, or a fixed item price value can be specified. If no target is specified, a gift line item can be added to the cart.
	Value CartDiscountValueDraft `json:"value"`
	// Valid [Cart Predicate](/../api/projects/predicates#cart-predicates).
	CartPredicate string `json:"cartPredicate"`
	// Must not be set when the `value` has type `giftLineItem`, otherwise a [CartDiscountTarget](ctp:api:type:CartDiscountTarget) must be set.
	Target CartDiscountTarget `json:"target,omitempty"`
	// Value between `0` and `1`.
	// A Discount with a higher sortOrder is prioritized.
	// The sort order must be unambiguous among all CartDiscounts.
	SortOrder string `json:"sortOrder"`
	// Only active Discounts can be applied to the Cart.
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
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [CartDiscount](ctp:api:type:CartDiscount).
*
 */
type CartDiscountResourceIdentifier struct {
	// Unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [CartDiscount](ctp:api:type:CartDiscount). Either `id` or `key` is required.
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
	case "shipping":
		obj := CartDiscountShippingCostTarget{}
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

type CartDiscountUpdate struct {
	// Expected version of the CartDiscount on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
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
	case "setKey":
		obj := CartDiscountSetKeyAction{}
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
	// An absolute Cart Discount will only match a price if this array contains a value with the same currency. If it contains 10€ and 15$, the matching € price will be decreased by 10€ and the matching $ price will be decreased by 15$.
	Money []Money `json:"money"`
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
*	Sets the [DiscountedLineItemPrice](ctp:api:type:DiscountedLineItemPrice) of the [CartDiscountLineItemsTarget](ctp:api:type:CartDiscountLineItemsTarget) or [CartDiscountCustomLineItemsTarget](ctp:api:type:CartDiscountCustomLineItemsTarget) to the value specified in the `money` field, if it is lower than the current Line Item price for the same currency. If the Line Item price is already discounted to a price equal to or lower than the respective price in the `money` field, this Discount is not applied.
*
 */
type CartDiscountValueFixed struct {
	// Cent precision money values in different currencies.
	Money []CentPrecisionMoney `json:"money"`
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
	// Money values in different currencies.
	// A fixed Cart Discount will only match a price if this array contains a value with the same currency. If it contains 10€ and 15$, the matching € price will be discounted by 10€ and the matching $ price will be discounted to 15$.
	Money []Money `json:"money"`
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

type CartDiscountValueGiftLineItemDraft struct {
	// ResourceIdentifier of a Product.
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
*	This Discount target is similar to `MultiBuyLineItems`, but is applied on Custom Line Items instead of Line Items.
*
 */
type MultiBuyCustomLineItemsTarget struct {
	// Valid [CustomLineItems target predicate](/../api/projects/predicates#customlineitem-field-identifiers). The Discount will be applied to Custom Line Items that are matched by the predicate.
	Predicate string `json:"predicate"`
	// Number of Custom Line Items to be present in order to trigger an application of this Discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Number of Custom Line Items that are discounted per application of this Discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times this Discount can be applied.
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
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of times this Discount can be applied.
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

type CartDiscountChangeValueAction struct {
	// New value to set.
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

type CartDiscountSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
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
