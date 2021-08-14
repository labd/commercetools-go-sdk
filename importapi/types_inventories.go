// Generated file, please do not change!!!
package importapi

import (
	"encoding/json"
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
	RestockableInDays float64   `json:"restockableInDays,omitEmpty"`
	ExpectedDelivery  time.Time `json:"expectedDelivery,omitEmpty"`
	// References a channel by its key.
	SupplyChannel ChannelKeyReference `json:"supplyChannel,omitEmpty"`
	// Maps to `Inventory.custom`.
	Custom *Custom `json:"custom,omitEmpty"`
}
