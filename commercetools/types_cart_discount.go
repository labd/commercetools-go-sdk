// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

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

func (obj *CartDiscount) UnmarshalJSON(data []byte) error {
	type Alias CartDiscount
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.References {
		obj.References[i] = AbstractReferenceDiscriminatorMapping(obj.References[i])
	}
	if obj.Target != nil {
		obj.Target = AbstractCartDiscountTargetDiscriminatorMapping(obj.Target)
	}
	if obj.Value != nil {
		obj.Value = AbstractCartDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

type CartDiscountChangeCartPredicateAction struct {
	CartPredicate string `json:"cartPredicate"`
}

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

func (obj CartDiscountChangeIsActiveAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeIsActiveAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeIsActive", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

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

func (obj CartDiscountChangeRequiresDiscountCodeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeRequiresDiscountCodeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeRequiresDiscountCode", Alias: (*Alias)(&obj)})
}

type CartDiscountChangeSortOrderAction struct {
	SortOrder string `json:"sortOrder"`
}

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

func (obj CartDiscountChangeTargetAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeTarget", Alias: (*Alias)(&obj)})
}
func (obj *CartDiscountChangeTargetAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeTargetAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractCartDiscountTargetDiscriminatorMapping(obj.Target)
	}

	return nil
}

type CartDiscountChangeValueAction struct {
	Value CartDiscountValue `json:"value"`
}

func (obj CartDiscountChangeValueAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountChangeValueAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeValue", Alias: (*Alias)(&obj)})
}
func (obj *CartDiscountChangeValueAction) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountChangeValueAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Value != nil {
		obj.Value = AbstractCartDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

type CartDiscountCustomLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

func (obj CartDiscountCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountCustomLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "customLineItems", Alias: (*Alias)(&obj)})
}

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

func (obj *CartDiscountDraft) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractCartDiscountTargetDiscriminatorMapping(obj.Target)
	}
	if obj.Value != nil {
		obj.Value = AbstractCartDiscountValueDiscriminatorMapping(obj.Value)
	}

	return nil
}

type CartDiscountLineItemsTarget struct {
	Predicate string `json:"predicate"`
}

func (obj CartDiscountLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "lineItems", Alias: (*Alias)(&obj)})
}

type CartDiscountPagedQueryResponse struct {
	Total   int            `json:"total,omitempty"`
	Offset  int            `json:"offset"`
	Count   int            `json:"count"`
	Results []CartDiscount `json:"results"`
}

type CartDiscountReference struct {
	Key string        `json:"key,omitempty"`
	ID  string        `json:"id,omitempty"`
	Obj *CartDiscount `json:"obj,omitempty"`
}

func (obj CartDiscountReference) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "cart-discount", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj CartDiscountSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CartDiscountSetCustomTypeAction struct {
	Type   *TypeReference `json:"type,omitempty"`
	Fields interface{}    `json:"fields,omitempty"`
}

func (obj CartDiscountSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CartDiscountSetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

func (obj CartDiscountSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAction struct {
	ValidFrom time.Time `json:"validFrom,omitempty"`
}

func (obj CartDiscountSetValidFromAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFrom", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidFromAndUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
	ValidFrom  time.Time `json:"validFrom,omitempty"`
}

func (obj CartDiscountSetValidFromAndUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidFromAndUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidFromAndUntil", Alias: (*Alias)(&obj)})
}

type CartDiscountSetValidUntilAction struct {
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

func (obj CartDiscountSetValidUntilAction) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountSetValidUntilAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setValidUntil", Alias: (*Alias)(&obj)})
}

type CartDiscountShippingCostTarget struct{}

func (obj CartDiscountShippingCostTarget) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountShippingCostTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "shipping", Alias: (*Alias)(&obj)})
}

type CartDiscountTarget interface{}
type AbstractCartDiscountTarget struct{}

func AbstractCartDiscountTargetDiscriminatorMapping(input CartDiscountTarget) CartDiscountTarget {
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

type CartDiscountUpdate struct {
	Version int                        `json:"version"`
	Actions []CartDiscountUpdateAction `json:"actions"`
}

func (obj *CartDiscountUpdate) UnmarshalJSON(data []byte) error {
	type Alias CartDiscountUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractCartDiscountUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type CartDiscountUpdateAction interface{}
type AbstractCartDiscountUpdateAction struct{}

func AbstractCartDiscountUpdateActionDiscriminatorMapping(input CartDiscountUpdateAction) CartDiscountUpdateAction {
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
			new.Target = AbstractCartDiscountTargetDiscriminatorMapping(new.Target)
		}

		return new
	case "changeValue":
		new := CartDiscountChangeValueAction{}
		mapstructure.Decode(input, &new)
		if new.Value != nil {
			new.Value = AbstractCartDiscountValueDiscriminatorMapping(new.Value)
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

type CartDiscountValue interface{}
type AbstractCartDiscountValue struct{}

func AbstractCartDiscountValueDiscriminatorMapping(input CartDiscountValue) CartDiscountValue {
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

type CartDiscountValueAbsolute struct {
	Money []Money `json:"money"`
}

func (obj CartDiscountValueAbsolute) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueAbsolute
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "absolute", Alias: (*Alias)(&obj)})
}

type CartDiscountValueGiftLineItem struct {
	VariantID           int               `json:"variantId"`
	SupplyChannel       *ChannelReference `json:"supplyChannel,omitempty"`
	Product             *ProductReference `json:"product"`
	DistributionChannel *ChannelReference `json:"distributionChannel,omitempty"`
}

func (obj CartDiscountValueGiftLineItem) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueGiftLineItem
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "giftLineItem", Alias: (*Alias)(&obj)})
}

type CartDiscountValueRelative struct {
	Permyriad int `json:"permyriad"`
}

func (obj CartDiscountValueRelative) MarshalJSON() ([]byte, error) {
	type Alias CartDiscountValueRelative
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "relative", Alias: (*Alias)(&obj)})
}

type MultiBuyCustomLineItemsTarget struct {
	TriggerQuantity    int           `json:"triggerQuantity"`
	SelectionMode      SelectionMode `json:"selectionMode"`
	Predicate          string        `json:"predicate"`
	MaxOccurrence      int           `json:"maxOccurrence,omitempty"`
	DiscountedQuantity int           `json:"discountedQuantity"`
}

func (obj MultiBuyCustomLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyCustomLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "multiBuyCustomLineItems", Alias: (*Alias)(&obj)})
}

type MultiBuyLineItemsTarget struct {
	TriggerQuantity    int           `json:"triggerQuantity"`
	SelectionMode      SelectionMode `json:"selectionMode"`
	Predicate          string        `json:"predicate"`
	MaxOccurrence      int           `json:"maxOccurrence,omitempty"`
	DiscountedQuantity int           `json:"discountedQuantity"`
}

func (obj MultiBuyLineItemsTarget) MarshalJSON() ([]byte, error) {
	type Alias MultiBuyLineItemsTarget
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "multiBuyLineItems", Alias: (*Alias)(&obj)})
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
