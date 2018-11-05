package shippingzones

import "encoding/json"

// ChangeName will change the name of the shipping zone being updated.
type ChangeName struct {
	Name string `json:"name"`
}

// MarshalJSON override to add the action value.
func (ua ChangeName) MarshalJSON() ([]byte, error) {
	type Alias ChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&ua),
	})
}

// SetDescription will set the description of the shipping zone being updated.
type SetDescription struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetDescription) MarshalJSON() ([]byte, error) {
	type Alias SetDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setDescription",
		Alias:  (*Alias)(&ua),
	})
}

// AddLocation will add a location to an existing shipping zone
type AddLocation struct {
	Location Location `json:"location"`
}

// MarshalJSON override to add the action value.
func (ua AddLocation) MarshalJSON() ([]byte, error) {
	type Alias AddLocation

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addLocation",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveLocation will remove a location from an existing shipping zone.
type RemoveLocation struct {
	Location Location `json:"location"`
}

// MarshalJSON override to add the action value.
func (ua RemoveLocation) MarshalJSON() ([]byte, error) {
	type Alias RemoveLocation

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeLocation",
		Alias:  (*Alias)(&ua),
	})
}
