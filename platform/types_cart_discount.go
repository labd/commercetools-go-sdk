// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type CartDiscount struct {
	// The unique ID of the cart discount.
	Id string `json:"id"`
	// The current version of the cart discount.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy      `json:"createdBy,omitEmpty"`
	Name      LocalizedString `json:"name"`
	// User-specific unique identifier for a cart discount.
	// Must be unique across a project.
	Key         string                 `json:"key,omitEmpty"`
	Description *LocalizedString       `json:"description,omitEmpty"`
	Value       CartDiscountValueDraft `json:"value"`
	// A valid Cart predicate.
	CartPredicate string `json:"cartPredicate"`
	// Empty when the `value` has type `giftLineItem`, otherwise a CartDiscountTarget is set.
	Target CartDiscountTarget `json:"target,omitEmpty"`
	// The string must contain a number between 0 and 1.
	// All matching cart discounts are applied to a cart in the order defined by this field.
	// A discount with greater sort order is prioritized higher than a discount with lower sort order.
	// The sort order is unambiguous among all cart discounts.
	SortOrder string `json:"sortOrder"`
	// Only active discount can be applied to the cart.
	IsActive   bool      `json:"isActive"`
	ValidFrom  time.Time `json:"validFrom,omitEmpty"`
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
	// States whether the discount can only be used in a connection with a DiscountCode.
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
	// The platform will generate this array from the predicate.
	// It contains the references of all the resources that are addressed in the predicate.
	References []Reference `json:"references"`
	// Specifies whether the application of this discount causes the following discounts to be ignored.
	// Defaults to Stacking.
	StackingMode StackingMode  `json:"stackingMode"`
	Custom       *CustomFields `json:"custom,omitEmpty"`
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

type CartDiscountDraft struct {
	Name LocalizedString `json:"name"`
	// User-specific unique identifier for a cart discount.
	// Must be unique across a project.
	// The field can be reset using the Set Key UpdateAction.
	Key         string                 `json:"key,omitEmpty"`
	Description *LocalizedString       `json:"description,omitEmpty"`
	Value       CartDiscountValueDraft `json:"value"`
	// A valid Cart predicate.
	CartPredicate string `json:"cartPredicate"`
	// Must not be set when the `value` has type `giftLineItem`, otherwise a CartDiscountTarget must be set.
	Target CartDiscountTarget `json:"target,omitEmpty"`
	// The string must contain a number between 0 and 1.
	// A discount with greater sort order is prioritized higher than a discount with lower sort order.
	// The sort order must be unambiguous among all cart discounts.
	SortOrder string `json:"sortOrder"`
	// Only active discount can be applied to the cart.
	// Defaults to `true`.
	IsActive   bool      `json:"isActive,omitEmpty"`
	ValidFrom  time.Time `json:"validFrom,omitEmpty"`
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
	// States whether the discount can only be used in a connection with a DiscountCode.
	// Defaults to `false`.
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
	// Specifies whether the application of this discount causes the following discounts to be ignored.
	// Defaults to Stacking.
	StackingMode StackingMode  `json:"stackingMode,omitEmpty"`
	Custom       *CustomFields `json:"custom,omitEmpty"`
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

type CartDiscountPagedQueryResponse struct {
	Limit   int            `json:"limit"`
	Count   int            `json:"count"`
	Total   int            `json:"total,omitEmpty"`
	Offset  int            `json:"offset"`
	Results []CartDiscount `json:"results"`
}

type CartDiscountReference struct {
	Id  string        `json:"id"`
	Obj *CartDiscount `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart-discount", Alias: (*Alias)(&obj)})
}

type CartDiscountResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "customLineItems":
		new := CartDiscountCustomLineItemsTarget{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "lineItems":
		new := CartDiscountLineItemsTarget{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "shipping":
		new := CartDiscountShippingCostTarget{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "multiBuyCustomLineItems":
		new := MultiBuyCustomLineItemsTarget{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "multiBuyLineItems":
		new := MultiBuyLineItemsTarget{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CartDiscountCustomLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountCustomLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "customLineItems", Alias: (*Alias)(&obj)})
}

type CartDiscountLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "lineItems", Alias: (*Alias)(&obj)})
}

type CartDiscountShippingCostTarget struct {
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountShippingCostTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountShippingCostTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "shipping", Alias: (*Alias)(&obj)})
}

type CartDiscountUpdate struct {
	Version int                        `json:"version"`
	Actions []CartDiscountUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type CartDiscountUpdateAction interface{}

func mapDiscriminatorCartDiscountUpdateAction(input interface{}) (CartDiscountUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "changeCartPredicate":
		new := CartDiscountChangeCartPredicateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeIsActive":
		new := CartDiscountChangeIsActiveAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := CartDiscountChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeRequiresDiscountCode":
		new := CartDiscountChangeRequiresDiscountCodeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSortOrder":
		new := CartDiscountChangeSortOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeStackingMode":
		new := CartDiscountChangeStackingModeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeTarget":
		new := CartDiscountChangeTargetAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeValue":
		new := CartDiscountChangeValueAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CartDiscountSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CartDiscountSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := CartDiscountSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CartDiscountSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFrom":
		new := CartDiscountSetValidFromAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidFromAndUntil":
		new := CartDiscountSetValidFromAndUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setValidUntil":
		new := CartDiscountSetValidUntilAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CartDiscountValue interface{}

func mapDiscriminatorCartDiscountValue(input interface{}) (CartDiscountValue, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "absolute":
		new := CartDiscountValueAbsolute{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "fixed":
		new := CartDiscountValueFixed{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "giftLineItem":
		new := CartDiscountValueGiftLineItem{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "relative":
		new := CartDiscountValueRelative{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CartDiscountValueAbsolute struct {
	Money []TypedMoney `json:"money"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountValueAbsolute) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountValueAbsolute
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

// MarshalJSON override to set the discriminator value
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
			return nil, errors.New("Error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "absolute":
		new := CartDiscountValueAbsoluteDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "fixed":
		new := CartDiscountValueFixedDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "giftLineItem":
		new := CartDiscountValueGiftLineItemDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "relative":
		new := CartDiscountValueRelativeDraft{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CartDiscountValueAbsoluteDraft struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueAbsoluteDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueAbsoluteDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "absolute", Alias: (*Alias)(&obj)})
}

type CartDiscountValueFixed struct {
	Money []TypedMoney `json:"money"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountValueFixed) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountValueFixed
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueFixed) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueFixed
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "fixed", Alias: (*Alias)(&obj)})
}

type CartDiscountValueFixedDraft struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueFixedDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueFixedDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "fixed", Alias: (*Alias)(&obj)})
}

type CartDiscountValueGiftLineItem struct {
	Product   ProductReference `json:"product"`
	VariantId int              `json:"variantId"`
	// The channel must have the role `InventorySupply`
	SupplyChannel ChannelReference `json:"supplyChannel,omitEmpty"`
	// The channel must have the role `ProductDistribution`
	DistributionChannel ChannelReference `json:"distributionChannel,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueGiftLineItem) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItem
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

type CartDiscountValueGiftLineItemDraft struct {
	Product   ProductResourceIdentifier `json:"product"`
	VariantId int                       `json:"variantId"`
	// The channel must have the role `InventorySupply`
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel,omitEmpty"`
	// The channel must have the role `ProductDistribution`
	DistributionChannel ChannelResourceIdentifier `json:"distributionChannel,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueGiftLineItemDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItemDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "giftLineItem", Alias: (*Alias)(&obj)})
}

type CartDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelative
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type CartDiscountValueRelativeDraft struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueRelativeDraft) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelativeDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "relative", Alias: (*Alias)(&obj)})
}

type MultiBuyCustomLineItemsTarget struct {
	// A valid custom line item target predicate. The discount will be applied to custom line items that are
	// matched by the predicate.
	Predicate string `json:"predicate"`
	// Quantity of line items that need to be present in order to trigger an application of this discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of line items that are discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of applications of this discount.
	MaxOccurrence int           `json:"maxOccurrence,omitEmpty"`
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value
func (obj MultiBuyCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyCustomLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

type MultiBuyLineItemsTarget struct {
	// A valid line item target predicate. The discount will be applied to line items that are matched by the predicate.
	Predicate string `json:"predicate"`
	// Quantity of line items that need to be present in order to trigger an application of this discount.
	TriggerQuantity int `json:"triggerQuantity"`
	// Quantity of line items that are discounted per application of this discount.
	DiscountedQuantity int `json:"discountedQuantity"`
	// Maximum number of applications of this discount.
	MaxOccurrence int           `json:"maxOccurrence,omitEmpty"`
	SelectionMode SelectionMode `json:"selectionMode"`
}

// MarshalJSON override to set the discriminator value
func (obj MultiBuyLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyLineItemsTarget
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "multiBuyLineItems", Alias: (*Alias)(&obj)})
}

type SelectionMode string

const (
	SelectionModeCheapest      SelectionMode = "Cheapest"
	SelectionModeMostExpensive SelectionMode = "MostExpensive"
)

type StackingMode string

const (
	StackingModeStacking              StackingMode = "Stacking"
	StackingModeStopAfterThisDiscount StackingMode = "StopAfterThisDiscount"
)

type CartDiscountChangeCartPredicateAction struct {
	// A valid Cart predicate.
	CartPredicate string `json:"cartPredicate"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeCartPredicateAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeCartPredicateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCartPredicate", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeIsActiveAction struct {
	IsActive bool `json:"isActive"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeRequiresDiscountCodeAction struct {
	RequiresDiscountCode bool `json:"requiresDiscountCode"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeRequiresDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeRequiresDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeRequiresDiscountCode", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeSortOrderAction struct {
	// The string must contain a number between 0 and 1.
	// A discount with greater sortOrder is prioritized higher than a discount with lower sortOrder.
	SortOrder string `json:"sortOrder"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeSortOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeSortOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSortOrder", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeStackingModeAction struct {
	StackingMode StackingMode `json:"stackingMode"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeStackingModeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeStackingModeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeStackingMode", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeTargetAction struct {
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

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeTargetAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTarget", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeValueAction struct {
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

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomFieldAction struct {
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an `InvalidOperation` error.
	// If `value` is provided, set the `value` of the field defined by the `name`.
	// The FieldDefinition determines the format for the `value` to be provided.
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *interface{} `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CartDiscountSetDescriptionAction struct {
	// If the `description` parameter is not included, the field will be emptied.
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CartDiscountSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAndUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidFrom time.Time `json:"validFrom,omitEmpty"`
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidUntilAction struct {
	// If absent, the field with the value is removed in case a value was set before.
	ValidUntil time.Time `json:"validUntil,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}
