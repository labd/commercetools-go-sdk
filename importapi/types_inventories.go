package importapi

// Generated file, please do not change!!!

import (
	"time"
)

/**
*	The data representation for an Inventory to be imported that is persisted as a [Inventory](/../api/projects/inventory#top) in the Project.
*
 */
type InventoryImport struct {
	// User-defined unique identifier. If an [InventoryEntry](/../api/projects/inventory#inventoryentry) with this `key` exists, it will be updated with the imported data.
	Key string `json:"key"`
	// Maps to `Inventory.sku`
	Sku string `json:"sku"`
	// Maps to `Inventory.quantityOnStock`
	QuantityOnStock int `json:"quantityOnStock"`
	// Maps to `Inventory.restockableInDays`
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// Maps to `Inventory.expectedDelivery`
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
	// Maps to `Inventory.supplyChannel`
	SupplyChannel *ChannelKeyReference `json:"supplyChannel,omitempty"`
	// Maps to `Inventory.custom`.
	Custom *Custom `json:"custom,omitempty"`
}
