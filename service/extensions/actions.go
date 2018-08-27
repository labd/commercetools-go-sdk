package extensions

import "encoding/json"

// SetKey will set a new key on the type being updated.
type SetKey struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to add the Action() value
func (ua SetKey) MarshalJSON() ([]byte, error) {
	type Alias SetKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setKey",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeTriggers is used to update an existing API Extension with
// a new triggers.
type ChangeTriggers struct {
	Messages []Trigger `json:"messages"`
}

func (ua ChangeTriggers) MarshalJSON() ([]byte, error) {
	type Alias ChangeTriggers

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeTriggers",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeDestination is used to update an existing API Extension with
// a new destination.
type ChangeDestination struct {
	Destination Destination `json:"destination"`
}

func (ua ChangeDestination) MarshalJSON() ([]byte, error) {
	type Alias ChangeDestination

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeDestination",
		Alias:  (*Alias)(&ua),
	})
}
