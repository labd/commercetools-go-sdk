package subscriptions

import "encoding/json"

// SetKey will change the key of the subscription being updated.
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

// SetMessages will set the messages of the subscription being updated.
type SetMessages struct {
	Messages []MessageSubscription `json:"messages"`
}

// MarshalJSON override to add the Action() value
func (ua SetMessages) MarshalJSON() ([]byte, error) {
	type Alias SetMessages

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setMessages",
		Alias:  (*Alias)(&ua),
	})
}

// SetChanges will set the changes of the subscription being updated.
type SetChanges struct {
	Changes []ChangeSubscription `json:"changes"`
}

// MarshalJSON override to add the Action() value
func (ua SetChanges) MarshalJSON() ([]byte, error) {
	type Alias SetChanges

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setChanges",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeDestination Updating the destination is eventually consistent,
// it may take up to a minute before it becomes fully active. During this time,
// messages may be delivered simultaneously to both the old and the
// new destination.
type ChangeDestination struct {
	Destination Destination `json:"destination"`
}

// MarshalJSON override to add the Action() value
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
