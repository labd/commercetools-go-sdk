package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type State struct {
	// Unique identifier of the State.
	ID string `json:"id"`
	// Current version of the State.
	Version int `json:"version"`
	// Date and time (UTC) the State was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the State was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the State.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the State.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the State.
	Key string `json:"key"`
	// Indicates to which resource or object types the State is assigned to.
	Type StateTypeEnum `json:"type"`
	// Name of the State.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the State.
	Description *LocalizedString `json:"description,omitempty"`
	// `true` for an initial State, the first State in a workflow.
	Initial bool `json:"initial"`
	// `true` for States that are an integral part of the [Project](ctp:api:type:Project). Those States cannot be deleted and their `key` cannot be changed.
	BuiltIn bool `json:"builtIn"`
	// Roles the State can fulfill for [Reviews](ctp:api:type:Review) and [Line Items](ctp:api:type:LineItem).
	Roles []StateRoleEnum `json:"roles"`
	// - list of States of the same `type` that the current State can be transitioned to. For example, when the current State is the _Initial_ State of [StateType](ctp:api:type:StateTypeEnum) `OrderState` and this list contains the reference to the _Shipped_ `OrderState`, the transition _Initial_ -> _Shipped_ is allowed.
	// - if empty, no transitions are allowed from the current State, defining the current State as final for this workflow.
	// - if not set, the validation is turned off and the current State can be transitioned to any other State of the same `type` as the current State.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["roles"] == nil {
		delete(raw, "roles")
	}

	if raw["transitions"] == nil {
		delete(raw, "transitions")
	}

	return json.Marshal(raw)

}

type StateDraft struct {
	// User-defined unique identifier for the State.
	Key string `json:"key"`
	// Specify to which resource or object type the State is assigned to.
	Type StateTypeEnum `json:"type"`
	// Name of the State.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the State.
	Description *LocalizedString `json:"description,omitempty"`
	// Set to `false` if the State is not the first step in a workflow.
	Initial *bool `json:"initial,omitempty"`
	// If suitable, assign predifined roles the State can fulfill in case the State's `type` is `LineItemState` or `ReviewState`.
	Roles []StateRoleEnum `json:"roles"`
	// Define the list of States of the same `type` to which the current State can be transitioned to.
	//
	// - If, for example, the current State is the _Initial_ State of [StateType](ctp:api:type:StateTypeEnum) `OrderState` and you want to allow the transition _Initial_ -> _Shipped_, then add the [StateResourceIdentifier](ctp:api:type:StateResourceIdentifier) to the _Shipped_ `OrderState` to this list.
	// - Set to empty list for not allowing any transition from the current State and defining it as final State for a workflow.
	// - Do not set this field at all to turn off validation and allowing transitions to any other State of the same `type` as the current State.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["roles"] == nil {
		delete(raw, "roles")
	}

	if raw["transitions"] == nil {
		delete(raw, "transitions")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with `results` containing an array of [State](ctp:api:type:State).
*
 */
type StatePagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [States](ctp:api:type:State) matching the query.
	Results []State `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [State](ctp:api:type:State).
*
 */
type StateReference struct {
	// Unique identifier of the referenced [State](ctp:api:type:State).
	ID string `json:"id"`
	// Contains the representation of the expanded State. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for States.
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

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [State](ctp:api:type:State). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type StateResourceIdentifier struct {
	// Unique identifier of the referenced [State](ctp:api:type:State). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [State](ctp:api:type:State). Required if `id` is absent.
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

/**
*	For some resource types, a State can fulfill the following predefined roles:
*
 */
type StateRoleEnum string

const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
	StateRoleEnumReturn                     StateRoleEnum = "Return"
)

/**
*	Resource or object type the State can be assigned to.
*
 */
type StateTypeEnum string

const (
	StateTypeEnumOrderState        StateTypeEnum = "OrderState"
	StateTypeEnumLineItemState     StateTypeEnum = "LineItemState"
	StateTypeEnumProductState      StateTypeEnum = "ProductState"
	StateTypeEnumReviewState       StateTypeEnum = "ReviewState"
	StateTypeEnumPaymentState      StateTypeEnum = "PaymentState"
	StateTypeEnumQuoteRequestState StateTypeEnum = "QuoteRequestState"
	StateTypeEnumStagedQuoteState  StateTypeEnum = "StagedQuoteState"
	StateTypeEnumQuoteState        StateTypeEnum = "QuoteState"
)

type StateUpdate struct {
	// Expected version of the State on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the State.
	Actions []StateUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *StateUpdate) UnmarshalJSON(data []byte) error {
	type Alias StateUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorStateUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type StateUpdateAction interface{}

func mapDiscriminatorStateUpdateAction(input interface{}) (StateUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
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
	// Value to append to the array.
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
	// Set to `true` for defining the State as initial State in a state machine and making it the first step in a workflow.
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
	// New value to set.
	// Must not be empty.
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
	// Resource or object types the State shall be assigned to.
	// Must not be empty.
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
	// Roles to remove from the State.
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
	// Value to set.
	// If empty, any existing value will be removed.
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
	// Value to set.
	// If empty, any existing value will be removed.
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
	// Value to set.
	// If empty, any existing value will be removed.
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
	// Value to set.
	// If empty, any existing value will be removed.
	//
	// Possible transformations of the current State to other States of the same `type` (for example, _Initial_ -> _Shipped_).
	// When performing a `transitionState` update action and `transitions` is set, the currently referenced State must have a transition to the new State.
	//
	// If `transitions` is an empty list, it means the current State is a final State and no further transitions are allowed.
	// If `transitions` is not set, the validation is turned off.
	//
	// When performing a `transitionState` update action, any other State of the same `type` can be transitioned to.
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

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["transitions"] == nil {
		delete(raw, "transitions")
	}

	return json.Marshal(raw)

}
