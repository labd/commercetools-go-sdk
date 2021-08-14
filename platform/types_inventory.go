// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type InventoryEntry struct {
	// The unique ID of the inventory entry.
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	Sku       string     `json:"sku"`
	// Connection to a particular supplier.
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Overall amount of stock.
	// (available + reserved)
	QuantityOnStock int `json:"quantityOnStock"`
	// Available amount of stock.
	// (available means: `quantityOnStock` - reserved quantity)
	AvailableQuantity int `json:"availableQuantity"`
	// The time period in days, that tells how often this inventory entry is restocked.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// The date and time of the next restock.
	ExpectedDelivery *time.Time    `json:"expectedDelivery,omitempty"`
	Custom           *CustomFields `json:"custom,omitempty"`
}

type InventoryEntryDraft struct {
	Sku               string                     `json:"sku"`
	SupplyChannel     *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	QuantityOnStock   int                        `json:"quantityOnStock"`
	RestockableInDays *int                       `json:"restockableInDays,omitempty"`
	ExpectedDelivery  *time.Time                 `json:"expectedDelivery,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type InventoryEntryReference struct {
	// Unique ID of the referenced resource.
	ID  string          `json:"id"`
	Obj *InventoryEntry `json:"obj,omitempty"`
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
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
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
		obj := InventoryEntryAddQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeQuantity":
		obj := InventoryEntryChangeQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeQuantity":
		obj := InventoryEntryRemoveQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := InventoryEntrySetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := InventoryEntrySetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExpectedDelivery":
		obj := InventoryEntrySetExpectedDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRestockableInDays":
		obj := InventoryEntrySetRestockableInDaysAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSupplyChannel":
		obj := InventoryEntrySetSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type InventoryPagedQueryResponse struct {
	Limit   int              `json:"limit"`
	Count   int              `json:"count"`
	Total   *int             `json:"total,omitempty"`
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
	Value interface{} `json:"value,omitempty"`
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
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitempty"`
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
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
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
	RestockableInDays *int `json:"restockableInDays,omitempty"`
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
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj InventoryEntrySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}
