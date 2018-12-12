// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// SelectionMode is an enum type
type SelectionMode string

// Enum values for SelectionMode
const (
	SelectionModeCheapest      SelectionMode = "Cheapest"
	SelectionModeMostExpensive SelectionMode = "MostExpensive"
)

// StackingMode is an enum type
type StackingMode string

// Enum values for StackingMode
const (
	StackingModeStacking              StackingMode = "Stacking"
	StackingModeStopAfterThisDiscount StackingMode = "StopAfterThisDiscount"
)

// CartDiscountTarget uses type as discriminator attribute
type CartDiscountTarget interface{}

func mapDiscriminatorCartDiscountTarget(input CartDiscountTarget) CartDiscountTarget {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "customLineItems":
		new := CartDiscountCustomLineItemsTarget{}
		mapstructure.Decode(input, &new)
		return new
	case "lineItems":
		new := CartDiscountLineItemsTarget{}
		mapstructure.Decode(input, &new)
		return new
	case "shipping":
		new := CartDiscountShippingCostTarget{}
		mapstructure.Decode(input, &new)
		return new
	case "multiBuyCustomLineItems":
		new := MultiBuyCustomLineItemsTarget{}
		mapstructure.Decode(input, &new)
		return new
	case "multiBuyLineItems":
		new := MultiBuyLineItemsTarget{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// CartDiscountUpdateAction uses action as discriminator attribute
type CartDiscountUpdateAction interface{}

func mapDiscriminatorCartDiscountUpdateAction(input CartDiscountUpdateAction) CartDiscountUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeCartPredicate":
		new := CartDiscountChangeCartPredicateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeIsActive":
		new := CartDiscountChangeIsActiveAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := CartDiscountChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeRequiresDiscountCode":
		new := CartDiscountChangeRequiresDiscountCodeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeSortOrder":
		new := CartDiscountChangeSortOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeStackingMode":
		new := CartDiscountChangeStackingModeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeTarget":
		new := CartDiscountChangeTargetAction{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = mapDiscriminatorCartDiscountTarget(new.Target)
		}

		return new
	case "changeValue":
		new := CartDiscountChangeValueAction{}
		mapstructure.Decode(input, &new)
		if new.Value != nil {
			new.Value = mapDiscriminatorCartDiscountValue(new.Value)
		}

		return new
	case "setCustomField":
		new := CartDiscountSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CartDiscountSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := CartDiscountSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFrom":
		new := CartDiscountSetValidFromAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidFromAndUntil":
		new := CartDiscountSetValidFromAndUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setValidUntil":
		new := CartDiscountSetValidUntilAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// CartDiscountValue uses type as discriminator attribute
type CartDiscountValue interface{}

func mapDiscriminatorCartDiscountValue(input CartDiscountValue) CartDiscountValue {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "absolute":
		new := CartDiscountValueAbsolute{}
		mapstructure.Decode(input, &new)
		return new
	case "giftLineItem":
		new := CartDiscountValueGiftLineItem{}
		mapstructure.Decode(input, &new)
		return new
	case "relative":
		new := CartDiscountValueRelative{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

// CartDiscount is of type Resource
type CartDiscount struct {
	Version              int                `json:"version"`
	LastModifiedAt       time.Time          `json:"lastModifiedAt"`
	ID                   string             `json:"id"`
	CreatedAt            time.Time          `json:"createdAt"`
	Value                CartDiscountValue  `json:"value"`
	ValidUntil           time.Time          `json:"validUntil,omitempty"`
	ValidFrom            time.Time          `json:"validFrom,omitempty"`
	Target               CartDiscountTarget `json:"target,omitempty"`
	StackingMode         StackingMode       `json:"stackingMode"`
	SortOrder            string             `json:"sortOrder"`
	RequiresDiscountCode bool               `json:"requiresDiscountCode"`
	References           []Reference        `json:"references"`
	Name                 *LocalizedString   `json:"name"`
	IsActive             bool               `json:"isActive"`
	Description          *LocalizedString   `json:"description,omitempty"`
	Custom               *CustomFields      `json:"custom,omitempty"`
	CartPredicate        string             `json:"cartPredicate"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscount) UnmarshalJSON(data []byte) error {
	type Alias CartDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		obj.References[i] = mapDiscriminatorReference(obj.References[i])
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorCartDiscountTarget(obj.Target)
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorCartDiscountValue(obj.Value)
	}

	return nil
}

// CartDiscountChangeCartPredicateAction implements the interface CartDiscountUpdateAction
type CartDiscountChangeCartPredicateAction struct {
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

// CartDiscountChangeIsActiveAction implements the interface CartDiscountUpdateAction
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

// CartDiscountChangeNameAction implements the interface CartDiscountUpdateAction
type CartDiscountChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// CartDiscountChangeRequiresDiscountCodeAction implements the interface CartDiscountUpdateAction
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

// CartDiscountChangeSortOrderAction implements the interface CartDiscountUpdateAction
type CartDiscountChangeSortOrderAction struct {
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

// CartDiscountChangeStackingModeAction implements the interface CartDiscountUpdateAction
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

// CartDiscountChangeTargetAction implements the interface CartDiscountUpdateAction
type CartDiscountChangeTargetAction struct {
	Target CartDiscountTarget `json:"target"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeTargetAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTarget", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountChangeTargetAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeTargetAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorCartDiscountTarget(obj.Target)
	}

	return nil
}

// CartDiscountChangeValueAction implements the interface CartDiscountUpdateAction
type CartDiscountChangeValueAction struct {
	Value CartDiscountValue `json:"value"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorCartDiscountValue(obj.Value)
	}

	return nil
}

// CartDiscountCustomLineItemsTarget implements the interface CartDiscountTarget
type CartDiscountCustomLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountCustomLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "customLineItems", Alias: (*Alias)(&obj)})
}

// CartDiscountDraft is a standalone struct
type CartDiscountDraft struct {
	Value                CartDiscountValue  `json:"value"`
	ValidUntil           time.Time          `json:"validUntil,omitempty"`
	ValidFrom            time.Time          `json:"validFrom,omitempty"`
	Target               CartDiscountTarget `json:"target,omitempty"`
	StackingMode         StackingMode       `json:"stackingMode,omitempty"`
	SortOrder            string             `json:"sortOrder"`
	RequiresDiscountCode bool               `json:"requiresDiscountCode"`
	Name                 *LocalizedString   `json:"name"`
	IsActive             bool               `json:"isActive,omitempty"`
	Description          *LocalizedString   `json:"description,omitempty"`
	Custom               *CustomFields      `json:"custom,omitempty"`
	CartPredicate        string             `json:"cartPredicate"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CartDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = mapDiscriminatorCartDiscountTarget(obj.Target)
	}
	if obj.Value != nil {
		obj.Value = mapDiscriminatorCartDiscountValue(obj.Value)
	}

	return nil
}

// CartDiscountLineItemsTarget implements the interface CartDiscountTarget
type CartDiscountLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "lineItems", Alias: (*Alias)(&obj)})
}

// CartDiscountPagedQueryResponse is of type PagedQueryResponse
type CartDiscountPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []CartDiscount `json:"results"`
}

// CartDiscountReference implements the interface Reference
type CartDiscountReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *CartDiscount `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "cart-discount", Alias: (*Alias)(&obj)})
}

// CartDiscountSetCustomFieldAction implements the interface CartDiscountUpdateAction
type CartDiscountSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CartDiscountSetCustomTypeAction implements the interface CartDiscountUpdateAction
type CartDiscountSetCustomTypeAction struct {
	Type   *TypeReference `json:"type,omitempty"`
	Fields interface{}    `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CartDiscountSetDescriptionAction implements the interface CartDiscountUpdateAction
type CartDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// CartDiscountSetValidFromAction implements the interface CartDiscountUpdateAction
type CartDiscountSetValidFromAction struct {
	ValidFrom time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

// CartDiscountSetValidFromAndUntilAction implements the interface CartDiscountUpdateAction
type CartDiscountSetValidFromAndUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
	ValidFrom  time.Time `json:"validFrom,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

// CartDiscountSetValidUntilAction implements the interface CartDiscountUpdateAction
type CartDiscountSetValidUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

// CartDiscountShippingCostTarget implements the interface CartDiscountTarget
type CartDiscountShippingCostTarget struct{}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountShippingCostTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountShippingCostTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "shipping", Alias: (*Alias)(&obj)})
}

// CartDiscountUpdate is of type Update
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
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorCartDiscountUpdateAction(obj.Actions[i])
	}

	return nil
}

// CartDiscountValueAbsolute implements the interface CartDiscountValue
type CartDiscountValueAbsolute struct {
	Money []Money `json:"money"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueAbsolute
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "absolute", Alias: (*Alias)(&obj)})
}

// CartDiscountValueGiftLineItem implements the interface CartDiscountValue
type CartDiscountValueGiftLineItem struct {
	VariantID           int               `json:"variantId"`
	SupplyChannel       *ChannelReference `json:"supplyChannel,omitempty"`
	Product             *ProductReference `json:"product"`
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueGiftLineItem) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItem
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "giftLineItem", Alias: (*Alias)(&obj)})
}

// CartDiscountValueRelative implements the interface CartDiscountValue
type CartDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

// MarshalJSON override to set the discriminator value
func (obj CartDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelative
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "relative", Alias: (*Alias)(&obj)})
}

// MultiBuyCustomLineItemsTarget implements the interface CartDiscountTarget
type MultiBuyCustomLineItemsTarget struct {
	TriggerQuantity    int           `json:"triggerQuantity"`
	SelectionMode      SelectionMode `json:"selectionMode"`
	Predicate          string        `json:"predicate"`
	MaxOccurrence      int           `json:"maxOccurrence,omitempty"`
	DiscountedQuantity int           `json:"discountedQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj MultiBuyCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyCustomLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

// MultiBuyLineItemsTarget implements the interface CartDiscountTarget
type MultiBuyLineItemsTarget struct {
	TriggerQuantity    int           `json:"triggerQuantity"`
	SelectionMode      SelectionMode `json:"selectionMode"`
	Predicate          string        `json:"predicate"`
	MaxOccurrence      int           `json:"maxOccurrence,omitempty"`
	DiscountedQuantity int           `json:"discountedQuantity"`
}

// MarshalJSON override to set the discriminator value
func (obj MultiBuyLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "multiBuyLineItems", Alias: (*Alias)(&obj)})
}
