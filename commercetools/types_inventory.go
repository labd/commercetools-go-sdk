// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type InventoryAddQuantityAction struct {
	Quantity int `json:"quantity"`
}

func (obj InventoryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

type InventoryChangeQuantityAction struct {
	Quantity int `json:"quantity"`
}

func (obj InventoryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntry struct {
	Version           int               `json:"version"`
	LastModifiedAt    time.Time         `json:"lastModifiedAt"`
	ID                string            `json:"id"`
	CreatedAt         time.Time         `json:"createdAt"`
	SupplyChannel     *ChannelReference `json:"supplyChannel,omitempty"`
	Sku               string            `json:"sku"`
	RestockableInDays int               `json:"restockableInDays,omitempty"`
	QuantityOnStock   int               `json:"quantityOnStock"`
	ExpectedDelivery  time.Time         `json:"expectedDelivery,omitempty"`
	Custom            *CustomFields     `json:"custom,omitempty"`
	AvailableQuantity int               `json:"availableQuantity"`
}

type InventoryEntryDraft struct {
	SupplyChannel     *ChannelReference  `json:"supplyChannel,omitempty"`
	Sku               string             `json:"sku"`
	RestockableInDays int                `json:"restockableInDays,omitempty"`
	QuantityOnStock   int                `json:"quantityOnStock"`
	ExpectedDelivery  time.Time          `json:"expectedDelivery,omitempty"`
	Custom            *CustomFieldsDraft `json:"custom,omitempty"`
}

type InventoryEntryReference struct {
	Key string          `json:"key,omitempty"`
	ID  string          `json:"id,omitempty"`
	Obj *InventoryEntry `json:"obj,omitempty"`
}

func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "inventory-entry", Alias: (*Alias)(&obj)})
}

type InventoryPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Offset  int              `json:"offset"`
	Count   int              `json:"count"`
	Results []InventoryEntry `json:"results"`
}

type InventoryRemoveQuantityAction struct {
	Quantity int `json:"quantity"`
}

func (obj InventoryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

type InventorySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj InventorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type InventorySetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj InventorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type InventorySetExpectedDeliveryAction struct {
	ExpectedDelivery time.Time `json:"expectedDelivery,omitempty"`
}

func (obj InventorySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

type InventorySetRestockableInDaysAction struct {
	RestockableInDays int `json:"restockableInDays,omitempty"`
}

func (obj InventorySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

type InventorySetSupplyChannelAction struct {
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
}

func (obj InventorySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}

type InventoryUpdate struct {
	Version int                     `json:"version"`
	Actions []InventoryUpdateAction `json:"actions"`
}

func (obj *InventoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractInventoryUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type InventoryUpdateAction interface{}
type AbstractInventoryUpdateAction struct{}

func AbstractInventoryUpdateActionDiscriminatorMapping(input InventoryUpdateAction) InventoryUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addQuantity":
		new := InventoryAddQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeQuantity":
		new := InventoryChangeQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeQuantity":
		new := InventoryRemoveQuantityAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := InventorySetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := InventorySetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setExpectedDelivery":
		new := InventorySetExpectedDeliveryAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setRestockableInDays":
		new := InventorySetRestockableInDaysAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setSupplyChannel":
		new := InventorySetSupplyChannelAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
