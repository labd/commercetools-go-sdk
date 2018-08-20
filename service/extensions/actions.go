package extensions

import "encoding/json"

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
