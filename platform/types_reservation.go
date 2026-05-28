package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"time"
)

/**
*	[Reference](ctp:api:type:Reference) to a [Reservation](ctp:api:type:Reservation).
*
 */
type ReservationReference struct {
	// Unique identifier of the referenced [Reservation](ctp:api:type:Reservation).
	ID string `json:"id"`
	// Contains the representation of the expanded Reservation.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Reservations.
	Obj *Reservation `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReservationReference) MarshalJSON() ([]byte, error) {
	type Alias ReservationReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "reservation", Alias: (*Alias)(&obj)})
}

/**
*	A Reservation represents a quantity of inventory that is reserved.
*
 */
type Reservation struct {
	// Unique identifier of the Reservation.
	ID string `json:"id"`
	// Current version of the Reservation.
	Version int `json:"version"`
	// Date and time (UTC) the Reservation was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Reservation was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the Reservation will no longer be valid. It can take a few minutes after expiry
	// for the `state` to be updated to `Expired` and the stock to be released.
	ExpiresAt time.Time `json:"expiresAt"`
	// IDs and references that last modified the Reservation.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Reservation.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Reference to the related [InventoryEntry](ctp:api:type:InventoryEntry) used to manage the inventory.
	InventoryEntry InventoryEntryReference `json:"inventoryEntry"`
	// The quantity of inventory that is reserved.
	Quantity int `json:"quantity"`
	// Reference to the owner of the Reservation.
	Owner Reference `json:"owner"`
	// Unique identifier of the LineItem this Reservation is associated with.
	LineItemId string `json:"lineItemId"`
	// Current state of the Reservation.
	State ReservationState `json:"state"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Reservation) UnmarshalJSON(data []byte) error {
	type Alias Reservation
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Owner != nil {
		var err error
		obj.Owner, err = mapDiscriminatorReference(obj.Owner)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Status of the [Reservation](ctp:api:type:Reservation).
*
 */
type ReservationState string

const (
	ReservationStateCreated   ReservationState = "Created"
	ReservationStateUpdated   ReservationState = "Updated"
	ReservationStateReleased  ReservationState = "Released"
	ReservationStateExpired   ReservationState = "Expired"
	ReservationStateCommitted ReservationState = "Committed"
)
