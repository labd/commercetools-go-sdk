// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"errors"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// InventoryUpdateAction uses action as discriminator attribute
type InventoryUpdateAction interface{}

func mapDiscriminatorInventoryUpdateAction(input interface{}) (InventoryUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Invalid discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}
	switch discriminator {
	case "addQuantity":
		new := InventoryAddQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeQuantity":
		new := InventoryChangeQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeQuantity":
		new := InventoryRemoveQuantityAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := InventorySetCustomFieldAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := InventorySetCustomTypeAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExpectedDelivery":
		new := InventorySetExpectedDeliveryAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRestockableInDays":
		new := InventorySetRestockableInDaysAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSupplyChannel":
		new := InventorySetSupplyChannelAction{}
		err := mapstructure.Decode(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

// InventoryAddQuantityAction implements the interface InventoryUpdateAction
type InventoryAddQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

// InventoryChangeQuantityAction implements the interface InventoryUpdateAction
type InventoryChangeQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

// InventoryEntry is of type Resource
type InventoryEntry struct {
	Version           int               `json:"version"`
	LastModifiedAt    time.Time         `json:"lastModifiedAt"`
	ID                string            `json:"id"`
	CreatedAt         time.Time         `json:"createdAt"`
	SupplyChannel     *ChannelReference `json:"supplyChannel,omitempty"`
	SKU               string            `json:"sku"`
	RestockableInDays int               `json:"restockableInDays,omitempty"`
	QuantityOnStock   int               `json:"quantityOnStock"`
	ExpectedDelivery  time.Time         `json:"expectedDelivery,omitempty"`
	Custom            *CustomFields     `json:"custom,omitempty"`
	AvailableQuantity int               `json:"availableQuantity"`
}

// InventoryEntryDraft is a standalone struct
type InventoryEntryDraft struct {
	SupplyChannel     *ChannelReference  `json:"supplyChannel,omitempty"`
	SKU               string             `json:"sku"`
	RestockableInDays int                `json:"restockableInDays,omitempty"`
	QuantityOnStock   int                `json:"quantityOnStock"`
	ExpectedDelivery  time.Time          `json:"expectedDelivery,omitempty"`
	Custom            *CustomFieldsDraft `json:"custom,omitempty"`
}

// InventoryEntryReference implements the interface Reference
type InventoryEntryReference struct {
	Key string          `json:"key,omitempty"`
	ID  string          `json:"id,omitempty"`
	Obj *InventoryEntry `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "inventory-entry", Alias: (*Alias)(&obj)})
}

// InventoryPagedQueryResponse is of type PagedQueryResponse
type InventoryPagedQueryResponse struct {
	Total   int              `json:"total,omitempty"`
	Offset  int              `json:"offset"`
	Count   int              `json:"count"`
	Results []InventoryEntry `json:"results"`
}

// InventoryRemoveQuantityAction implements the interface InventoryUpdateAction
type InventoryRemoveQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

// InventorySetCustomFieldAction implements the interface InventoryUpdateAction
type InventorySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj InventorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// InventorySetCustomTypeAction implements the interface InventoryUpdateAction
type InventorySetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// InventorySetExpectedDeliveryAction implements the interface InventoryUpdateAction
type InventorySetExpectedDeliveryAction struct {
	ExpectedDelivery time.Time `json:"expectedDelivery,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventorySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

// InventorySetRestockableInDaysAction implements the interface InventoryUpdateAction
type InventorySetRestockableInDaysAction struct {
	RestockableInDays int `json:"restockableInDays,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventorySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

// InventorySetSupplyChannelAction implements the interface InventoryUpdateAction
type InventorySetSupplyChannelAction struct {
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventorySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventorySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}

// InventoryUpdate is of type Update
type InventoryUpdate struct {
	Version int                     `json:"version"`
	Actions []InventoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorInventoryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
