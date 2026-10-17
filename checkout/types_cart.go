package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Reference to a [Cart](/api/projects/carts).
*
 */
type CartReference struct {
	// Unique identifier of the referenced [Cart](/api/projects/carts).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartReference) MarshalJSON() ([]byte, error) {
	type Alias CartReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

/**
*	Resource identifier to a [Cart](/api/projects/carts). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](ctp:api:type:InvalidJsonInputError) error is returned.
*
 */
type CartResourceIdentifier struct {
	// Unique identifier of the referenced [Cart](/api/projects/carts). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Cart](/api/projects/carts). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CartResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CartResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "cart", Alias: (*Alias)(&obj)})
}

/**
*	Reference to an [Order](/api/projects/orders).
*
 */
type OrderReference struct {
	// Unique identifier of the referenced [Order](/api/projects/orders).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj OrderReference) MarshalJSON() ([]byte, error) {
	type Alias OrderReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "order", Alias: (*Alias)(&obj)})
}
