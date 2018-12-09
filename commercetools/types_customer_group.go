// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CustomerGroup struct {
	Version        int           `json:"version"`
	LastModifiedAt time.Time     `json:"lastModifiedAt"`
	ID             string        `json:"id"`
	CreatedAt      time.Time     `json:"createdAt"`
	Name           string        `json:"name"`
	Key            string        `json:"key,omitempty"`
	Custom         *CustomFields `json:"custom,omitempty"`
}

type CustomerGroupChangeNameAction struct {
	Name string `json:"name"`
}

func (obj CustomerGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CustomerGroupDraft struct {
	Key       string        `json:"key,omitempty"`
	GroupName string        `json:"groupName"`
	Custom    *CustomFields `json:"custom,omitempty"`
}

type CustomerGroupPagedQueryResponse struct {
	Total   int             `json:"total,omitempty"`
	Offset  int             `json:"offset"`
	Count   int             `json:"count"`
	Results []CustomerGroup `json:"results"`
}

type CustomerGroupReference struct {
	Key string         `json:"key,omitempty"`
	ID  string         `json:"id,omitempty"`
	Obj *CustomerGroup `json:"obj,omitempty"`
}

func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "customer-group", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj CustomerGroupSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj CustomerGroupSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj CustomerGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CustomerGroupUpdate struct {
	Version int                         `json:"version"`
	Actions []CustomerGroupUpdateAction `json:"actions"`
}

func (obj *CustomerGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractCustomerGroupUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type CustomerGroupUpdateAction interface{}
type AbstractCustomerGroupUpdateAction struct{}

func AbstractCustomerGroupUpdateActionDiscriminatorMapping(input CustomerGroupUpdateAction) CustomerGroupUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeName":
		new := CustomerGroupChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := CustomerGroupSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CustomerGroupSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := CustomerGroupSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
