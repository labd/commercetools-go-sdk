package shippingzones

import "time"

// ShippingZone allows you to specific shipping costs per zone
type ShippingZone struct {
	// The unique ID of the category.
	ID string `json:"id"`

	// User-specific unique identifier for the category.
	Key string `json:"key,omitempty"`

	// The current version of the category.
	Version int `json:"version"`

	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`

	Locations []Location `json:"locations"`
}

// ShippingZoneDraft is sent with Create ShippingZone requests.
type ShippingZoneDraft struct {
	Name        string     `json:"name"`
	Key         string     `json:"key,omitempty"`
	Description string     `json:"description,omitempty"`
	Locations   []Location `json:"locations"`
}

type Location struct {
	Country string `json:"country"`
	State   string `json:"state,omitempty"`
}
