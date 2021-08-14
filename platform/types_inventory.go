// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type InventoryEntry struct {
	// The unique ID of the inventory entry.
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	Sku       string     `json:"sku"`
	// Optional connection to a particular supplier.
	SupplyChannel ChannelReference `json:"supplyChannel,omitEmpty"`
	// Overall amount of stock.
	// (available + reserved)
	QuantityOnStock int `json:"quantityOnStock"`
	// Available amount of stock.
	// (available means: `quantityOnStock` - reserved quantity)
	AvailableQuantity int `json:"availableQuantity"`
	// The time period in days, that tells how often this inventory entry is restocked.
	RestockableInDays int `json:"restockableInDays,omitEmpty"`
	// The date and time of the next restock.
	ExpectedDelivery time.Time     `json:"expectedDelivery,omitEmpty"`
	Custom           *CustomFields `json:"custom,omitEmpty"`
}

type InventoryEntryDraft struct {
	Sku               string                    `json:"sku"`
	SupplyChannel     ChannelResourceIdentifier `json:"supplyChannel,omitEmpty"`
	QuantityOnStock   int                       `json:"quantityOnStock"`
	RestockableInDays int                       `json:"restockableInDays,omitEmpty"`
	ExpectedDelivery  time.Time                 `json:"expectedDelivery,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
}

type InventoryEntryReference struct {
	Id  string          `json:"id"`
	Obj *InventoryEntry `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

type InventoryEntryResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

type InventoryEntryUpdate struct {
	Version int                          `json:"version"`
	Actions []InventoryEntryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type InventoryEntryUpdateAction interface{}

func mapDiscriminatorInventoryEntryUpdateAction(input interface{}) (InventoryEntryUpdateAction, error) {

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
	case "addQuantity":
		new := InventoryEntryAddQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeQuantity":
		new := InventoryEntryChangeQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeQuantity":
		new := InventoryEntryRemoveQuantityAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := InventoryEntrySetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := InventoryEntrySetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExpectedDelivery":
		new := InventoryEntrySetExpectedDeliveryAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRestockableInDays":
		new := InventoryEntrySetRestockableInDaysAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setSupplyChannel":
		new := InventoryEntrySetSupplyChannelAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type InventoryPagedQueryResponse struct {
	Limit   int              `json:"limit"`
	Count   int              `json:"count"`
	Total   int              `json:"total,omitEmpty"`
	Offset  int              `json:"offset"`
	Results []InventoryEntry `json:"results"`
}

type InventoryEntryAddQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntryChangeQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntryRemoveQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetExpectedDeliveryAction struct {
	ExpectedDelivery time.Time `json:"expectedDelivery,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetRestockableInDaysAction struct {
	RestockableInDays int `json:"restockableInDays,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetSupplyChannelAction struct {
	// If absent, the supply channel is removed.
	// This action will fail if an entry with the combination of sku and supplyChannel already exists.
	SupplyChannel ChannelResourceIdentifier `json:"supplyChannel,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}
