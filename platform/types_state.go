// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type State struct {
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 2019-02-01 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// A unique identifier for the state.
	Key  string        `json:"key"`
	Type StateTypeEnum `json:"type"`
	// A human-readable name of the state.
	Name *LocalizedString `json:"name,omitempty"`
	// A human-readable description of the state.
	Description *LocalizedString `json:"description,omitempty"`
	// A state can be declared as an initial state for any state machine.
	// When a workflow starts, this first state must be an `initial` state.
	Initial bool `json:"initial"`
	// Builtin states are integral parts of the project that cannot be deleted nor the key can be changed.
	BuiltIn bool            `json:"builtIn"`
	Roles   []StateRoleEnum `json:"roles"`
	// Transitions are a way to describe possible transformations of the current state to other states of the same `type` (e.g.: _Initial_ -> _Shipped_).
	// When performing a `transitionState` update action and `transitions` is set, the currently referenced state must have a transition to the new state.
	// If `transitions` is an empty list, it means the current state is a final state and no further transitions are allowed.
	// If `transitions` is not set, the validation is turned off.
	// When performing a `transitionState` update action, any other state of the same `type` can be transitioned to.
	Transitions []StateReference `json:"transitions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj State) MarshalJSON() ([]byte, error) {
	type Alias State
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["roles"] == nil {
		delete(target, "roles")
	}

	if target["transitions"] == nil {
		delete(target, "transitions")
	}

	return json.Marshal(target)
}

type StateDraft struct {
	Key         string                    `json:"key"`
	Type        StateTypeEnum             `json:"type"`
	Name        *LocalizedString          `json:"name,omitempty"`
	Description *LocalizedString          `json:"description,omitempty"`
	Initial     *bool                     `json:"initial,omitempty"`
	Roles       []StateRoleEnum           `json:"roles"`
	Transitions []StateResourceIdentifier `json:"transitions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateDraft) MarshalJSON() ([]byte, error) {
	type Alias StateDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["roles"] == nil {
		delete(target, "roles")
	}

	if target["transitions"] == nil {
		delete(target, "transitions")
	}

	return json.Marshal(target)
}

type StatePagedQueryResponse struct {
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
	Total   *int    `json:"total,omitempty"`
	Offset  int     `json:"offset"`
	Results []State `json:"results"`
}

type StateReference struct {
	// Unique ID of the referenced resource.
	ID  string `json:"id"`
	Obj *State `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateReference) MarshalJSON() ([]byte, error) {
	type Alias StateReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

type StateResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias StateResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

type StateRoleEnum string

const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
	StateRoleEnumReturn                     StateRoleEnum = "Return"
)

type StateTypeEnum string

const (
	StateTypeEnumOrderState    StateTypeEnum = "OrderState"
	StateTypeEnumLineItemState StateTypeEnum = "LineItemState"
	StateTypeEnumProductState  StateTypeEnum = "ProductState"
	StateTypeEnumReviewState   StateTypeEnum = "ReviewState"
	StateTypeEnumPaymentState  StateTypeEnum = "PaymentState"
)

type StateUpdate struct {
	Version int                 `json:"version"`
	Actions []StateUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StateUpdate) UnmarshalJSON(data []byte) error {
	type Alias StateUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type StateUpdateAction interface{}

func mapDiscriminatorStateUpdateAction(input interface{}) (StateUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "addRoles":
		obj := StateAddRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeInitial":
		obj := StateChangeInitialAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeKey":
		obj := StateChangeKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeType":
		obj := StateChangeTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeRoles":
		obj := StateRemoveRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := StateSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := StateSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRoles":
		obj := StateSetRolesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTransitions":
		obj := StateSetTransitionsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type StateAddRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateAddRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateAddRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addRoles", Alias: (*Alias)(&obj)})
}

type StateChangeInitialAction struct {
	Initial bool `json:"initial"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateChangeInitialAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeInitialAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeInitial", Alias: (*Alias)(&obj)})
}

type StateChangeKeyAction struct {
	Key string `json:"key"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateChangeKeyAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeKey", Alias: (*Alias)(&obj)})
}

type StateChangeTypeAction struct {
	Type StateTypeEnum `json:"type"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateChangeTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeType", Alias: (*Alias)(&obj)})
}

type StateRemoveRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

type StateSetDescriptionAction struct {
	Description LocalizedString `json:"description"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type StateSetNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type StateSetRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

type StateSetTransitionsAction struct {
	Transitions []StateResourceIdentifier `json:"transitions"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StateSetTransitionsAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetTransitionsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransitions", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["transitions"] == nil {
		delete(target, "transitions")
	}

	return json.Marshal(target)
}
