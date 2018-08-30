package statemachines

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// ChangeKey will change the key of the state being updated.
type ChangeKey struct {
	Key string `json:"key"`
}

// MarshalJSON override to add the action value.
func (ua ChangeKey) MarshalJSON() ([]byte, error) {
	type Alias ChangeKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeKey",
		Alias:  (*Alias)(&ua),
	})
}

// SetName will set the name of the state being updated.
type SetName struct {
	Name commercetools.LocalizedString `json:"name"`
}

// MarshalJSON override to add the action value.
func (ua SetName) MarshalJSON() ([]byte, error) {
	type Alias SetName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setName",
		Alias:  (*Alias)(&ua),
	})
}

// SetDescription will set the description of the state being updated.
type SetDescription struct {
	Description commercetools.LocalizedString `json:"description"`
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

// ChangeType will change the type of the state being updated.
type ChangeType struct {
	Type StateType `json:"type"`
}

// MarshalJSON override to add the action value.
func (ua ChangeType) MarshalJSON() ([]byte, error) {
	type Alias ChangeType

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeType",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeInitial will change the initial state of the state being updated.
type ChangeInitial struct {
	Initial bool `json:"initial"`
}

// MarshalJSON override to add the action value.
func (ua ChangeInitial) MarshalJSON() ([]byte, error) {
	type Alias ChangeInitial

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeInitial",
		Alias:  (*Alias)(&ua),
	})
}

// SetTransitions will set the transitions of the state being updated.
// Transitions are a way to describe possible transformations of the current
// state to other states of the same type (e.g.: Initial -> Shipped). When
// performing a transitionState update action and transitions is set, the
// currently referenced state must have a transition to the new state. If
// transitions is an empty list, it means the current state is a final state and
// no further transitions are allowed. If transitions is not set, the validation
// is turned off. When performing a transitionState update action, any other
// state of the same type can be transitioned to.
type SetTransitions struct {
	Transitions []commercetools.Reference `json:"transitions,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetTransitions) MarshalJSON() ([]byte, error) {
	type Alias SetTransitions

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setTransitions",
		Alias:  (*Alias)(&ua),
	})
}

// SetRoles will set the roles of the state being updated.
type SetRoles struct {
	Roles []StateRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua SetRoles) MarshalJSON() ([]byte, error) {
	type Alias SetRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setRoles",
		Alias:  (*Alias)(&ua),
	})
}

// AddRoles will add roles to the state being updated.
type AddRoles struct {
	Roles []StateRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua AddRoles) MarshalJSON() ([]byte, error) {
	type Alias AddRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addRoles",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveRoles will remove roles from the state being updated.
type RemoveRoles struct {
	Roles []StateRole `json:"roles"`
}

// MarshalJSON override to add the action value.
func (ua RemoveRoles) MarshalJSON() ([]byte, error) {
	type Alias RemoveRoles

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeRoles",
		Alias:  (*Alias)(&ua),
	})
}
