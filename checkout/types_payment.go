package checkout

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Reference to a [Payment](/api/projects/payments).
*
 */
type PaymentReference struct {
	// Unique identifier of the referenced [Payment](/api/projects/payments).
	ID string `json:"id"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentReference) MarshalJSON() ([]byte, error) {
	type Alias PaymentReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}

/**
*	Resource identifier to a [Payment](/api/projects/payments). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](ctp:api:type:InvalidJsonInputError) error is returned.
*
 */
type PaymentResourceIdentifier struct {
	// Unique identifier of the referenced [Payment](/api/projects/payments). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Payment](/api/projects/payments). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj PaymentResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias PaymentResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "payment", Alias: (*Alias)(&obj)})
}
