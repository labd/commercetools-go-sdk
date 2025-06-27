package importapi

// Generated file, please do not change!!!

import (
	"time"
)

/**
*	Represents the data used to import an InventoryEntry. Once imported, this data is persisted as a [InventoryEntry](ctp:api:type:InventoryEntry) in the Project.
*
 */
type InventoryImport struct {
	// User-defined unique identifier. If an [InventoryEntry](ctp:api:type:InventoryEntry) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `InventoryEntry.sku`
	Sku string `json:"sku"`
	// Maps to `InventoryEntry.quantityOnStock`
	QuantityOnStock int `json:"quantityOnStock"`
	// Maps to `InventoryEntry.restockableInDays`
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Maps to `InventoryEntry.expectedDelivery`
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
	// Maps to `InventoryEntry.supplyChannel`. If the referenced [Channel](ctp:api:type:Channel) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Channel is created.
	SupplyChannel *ChannelKeyReference `json:"supplyChannel,omitempty"`
	// Maps to `InventoryEntry.custom`.
	Custom *Custom `json:"custom,omitempty"`
}
