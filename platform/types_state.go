// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type State struct {
	Id             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// A unique identifier for the state.
	Key  string        `json:"key"`
	Type StateTypeEnum `json:"type"`
	// A human-readable name of the state.
	Name *LocalizedString `json:"name,omitEmpty"`
	// A human-readable description of the state.
	Description *LocalizedString `json:"description,omitEmpty"`
	// A state can be declared as an initial state for any state machine.
	// When a workflow starts, this first state must be an `initial` state.
	Initial bool `json:"initial"`
	// Builtin states are integral parts of the project that cannot be deleted nor the key can be changed.
	BuiltIn bool            `json:"builtIn"`
	Roles   []StateRoleEnum `json:"roles,omitEmpty"`
	// Transitions are a way to describe possible transformations of the current state to other states of the same `type` (e.g.: _Initial_ -> _Shipped_).
	// When performing a `transitionState` update action and `transitions` is set, the currently referenced state must have a transition to the new state.
	// If `transitions` is an empty list, it means the current state is a final state and no further transitions are allowed.
	// If `transitions` is not set, the validation is turned off.
	// When performing a `transitionState` update action, any other state of the same `type` can be transitioned to.
	Transitions []StateReference `json:"transitions,omitEmpty"`
}

type StateDraft struct {
	Key         string                    `json:"key"`
	Type        StateTypeEnum             `json:"type"`
	Name        *LocalizedString          `json:"name,omitEmpty"`
	Description *LocalizedString          `json:"description,omitEmpty"`
	Initial     bool                      `json:"initial,omitEmpty"`
	Roles       []StateRoleEnum           `json:"roles,omitEmpty"`
	Transitions []StateResourceIdentifier `json:"transitions,omitEmpty"`
}

type StatePagedQueryResponse struct {
	Limit   int     `json:"limit"`
	Count   int     `json:"count"`
	Total   int     `json:"total,omitEmpty"`
	Offset  int     `json:"offset"`
	Results []State `json:"results"`
}

type StateReference struct {
	Id  string `json:"id"`
	Obj *State `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj StateReference) MarshalJSON() ([]byte, error) {
	type Alias StateReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "state", Alias: (*Alias)(&obj)})
}

type StateResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
		new := StateAddRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeInitial":
		new := StateChangeInitialAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeKey":
		new := StateChangeKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeType":
		new := StateChangeTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeRoles":
		new := StateRemoveRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := StateSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setName":
		new := StateSetNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRoles":
		new := StateSetRolesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTransitions":
		new := StateSetTransitionsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type StateAddRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj StateSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

type StateSetTransitionsAction struct {
	Transitions []StateResourceIdentifier `json:"transitions,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj StateSetTransitionsAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetTransitionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransitions", Alias: (*Alias)(&obj)})
}
