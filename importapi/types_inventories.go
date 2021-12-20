// Generated file, please do not change!!!
package importapi

import (
	"time"
)

/**
*	The data representation for an Inventory to be imported that is persisted as a [Inventory](/../api/projects/inventory#top) in the Project.
*
 */
type InventoryImport struct {
	Key string `json:"key"`
	Sku string `json:"sku"`
	// Maps to `Inventory.quantityOnStock`
	QuantityOnStock float64 `json:"quantityOnStock"`
	// Maps to `Inventory.restockableInDays`
	RestockableInDays *float64   `json:"restockableInDays,omitempty"`
	ExpectedDelivery  *time.Time `json:"expectedDelivery,omitempty"`
	// References a channel by key.
	SupplyChannel *ChannelKeyReference `json:"supplyChannel,omitempty"`
	// Maps to `Inventory.custom`.
	Custom *Custom `json:"custom,omitempty"`
}
