// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type State struct {
	Version        int              `json:"version"`
	LastModifiedAt time.Time        `json:"lastModifiedAt"`
	ID             string           `json:"id"`
	CreatedAt      time.Time        `json:"createdAt"`
	Type           StateTypeEnum    `json:"type"`
	Transitions    []StateReference `json:"transitions,omitempty"`
	Roles          []StateRoleEnum  `json:"roles,omitempty"`
	Name           *LocalizedString `json:"name,omitempty"`
	Key            string           `json:"key"`
	Initial        bool             `json:"initial"`
	Description    *LocalizedString `json:"description,omitempty"`
	BuiltIn        bool             `json:"builtIn"`
}

type StateAddRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

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

func (obj StateChangeTypeAction) MarshalJSON() ([]byte, error) {
	type Alias StateChangeTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeType", Alias: (*Alias)(&obj)})
}

type StateDraft struct {
	Type        StateTypeEnum    `json:"type"`
	Transitions []StateReference `json:"transitions,omitempty"`
	Roles       []StateRoleEnum  `json:"roles,omitempty"`
	Name        *LocalizedString `json:"name,omitempty"`
	Key         string           `json:"key"`
	Initial     bool             `json:"initial,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
}

type StatePagedQueryResponse struct {
	Total   int     `json:"total,omitempty"`
	Offset  int     `json:"offset"`
	Count   int     `json:"count"`
	Results []State `json:"results"`
}

type StateReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *State `json:"obj,omitempty"`
}

func (obj StateReference) MarshalJSON() ([]byte, error) {
	type Alias StateReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "state", Alias: (*Alias)(&obj)})
}

type StateRemoveRolesAction struct {
	Roles []StateRoleEnum `json:"roles"`
}

func (obj StateRemoveRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateRemoveRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeRoles", Alias: (*Alias)(&obj)})
}

type StateRoleEnum string

const (
	StateRoleEnumReviewIncludedInStatistics StateRoleEnum = "ReviewIncludedInStatistics"
)

type StateSetDescriptionAction struct {
	Description *LocalizedString `json:"description"`
}

func (obj StateSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type StateSetNameAction struct {
	Name *LocalizedString `json:"name"`
}

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

func (obj StateSetRolesAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetRolesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRoles", Alias: (*Alias)(&obj)})
}

type StateSetTransitionsAction struct {
	Transitions []StateReference `json:"transitions,omitempty"`
}

func (obj StateSetTransitionsAction) MarshalJSON() ([]byte, error) {
	type Alias StateSetTransitionsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTransitions", Alias: (*Alias)(&obj)})
}

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

func (obj *StateUpdate) UnmarshalJSON(data []byte) error {
	type Alias StateUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractStateUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type StateUpdateAction interface{}
type AbstractStateUpdateAction struct{}

func AbstractStateUpdateActionDiscriminatorMapping(input StateUpdateAction) StateUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addRoles":
		new := StateAddRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeInitial":
		new := StateChangeInitialAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeKey":
		new := StateChangeKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeType":
		new := StateChangeTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeRoles":
		new := StateRemoveRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := StateSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setName":
		new := StateSetNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setRoles":
		new := StateSetRolesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTransitions":
		new := StateSetTransitionsAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
