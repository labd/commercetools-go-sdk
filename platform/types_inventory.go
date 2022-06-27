package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type InventoryEntry struct {
	// Unique identifier of the InventoryEntry.
	ID string `json:"id"`
	// Current version of the InventoryEntry.
	Version int `json:"version"`
	// Date and time (UTC) the InventoryEntry was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the InventoryEntry was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the InventoryEntry.
	Key *string `json:"key,omitempty"`
	// [ProductVariant](ctp:api:type:ProductVariant) `sku` of the InventoryEntry.
	Sku string `json:"sku"`
	// [Channel](ctp:api:type:Channel) that supplies this InventoryEntry.
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Overall amount of stock (`availableQuantity` + reserved).
	QuantityOnStock int `json:"quantityOnStock"`
	// Available amount of stock (`quantityOnStock` - reserved).
	AvailableQuantity int `json:"availableQuantity"`
	// How often the InventoryEntry is restocked (in days).
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Date and time of the next restock.
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
	// Custom Fields of the InventoryEntry.
	Custom *CustomFields `json:"custom,omitempty"`
}

type InventoryEntryDraft struct {
	// [ProductVariant](ctp:api:type:ProductVariant) `sku` of the InventoryEntry.
	Sku string `json:"sku"`
	// User-defined unique identifier for the InventoryEntry.
	Key *string `json:"key,omitempty"`
	// [Channel](ctp:api:type:Channel) that supplies this InventoryEntry.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	// Overall amount of stock.
	QuantityOnStock int `json:"quantityOnStock"`
	// How often the InventoryEntry is restocked (in days).
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Date and time of the next restock.
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
	// Custom Fields of the InventoryEntry.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[Reference](ctp:api:type:Reference) to an [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type InventoryEntryReference struct {
	// Unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry).
	ID string `json:"id"`
	// Contains the representation of the expanded InventoryEntry. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for InventoryEntries.
	Obj *InventoryEntry `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to an [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type InventoryEntryResourceIdentifier struct {
	// Unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

type InventoryEntryUpdate struct {
	// Expected version of the InventoryEntry on which the changes should be applied. If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the InventoryEntry.
	Actions []InventoryEntryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorInventoryEntryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type InventoryEntryUpdateAction interface{}

func mapDiscriminatorInventoryEntryUpdateAction(input interface{}) (InventoryEntryUpdateAction, error) {
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
	case "setKey":
		obj := InventoryEntrySetKeyAction{}
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
	// [Inventory entries](ctp:api:type:InventoryEntry) matching the query.
	Results []InventoryEntry `json:"results"`
}

/**
*	Updates `availableQuantity` based on the new `quantityOnStock` and amount of active reservations.
 */
type InventoryEntryAddQuantityAction struct {
	// Value to add to `quantityOnStock`.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

/**
*	Updates `availableQuantity` based on the new `quantityOnStock` and amount of active reservations.
 */
type InventoryEntryChangeQuantityAction struct {
	// Value to set for `quantityOnStock`.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

/**
*	Updates `availableQuantity` based on the new `quantityOnStock` and amount of active reservations.
 */
type InventoryEntryRemoveQuantityAction struct {
	// Value to remove from `quantityOnStock`.
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the InventoryEntry with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the InventoryEntry.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the InventoryEntry.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetExpectedDeliveryAction struct {
	// Value to set. If empty, any existing value will be removed.
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetRestockableInDaysAction struct {
	// Value to set. If empty, any existing value will be removed.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

/**
*	If an entry with the same `sku` and `supplyChannel` already exists, this action will fail and a [400 Bad Request](/../api/errors#400-bad-request-1) `DuplicateField` error will be returned.
 */
type InventoryEntrySetSupplyChannelAction struct {
	// Value to set. If empty, any existing value will be removed.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}
