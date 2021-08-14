// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type CustomerGroup struct {
	// The unique ID of the customer group.
	Id string `json:"id"`
	// The current version of the customer group.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// User-specific unique identifier for the customer group.
	Key    string        `json:"key,omitEmpty"`
	Name   string        `json:"name"`
	Custom *CustomFields `json:"custom,omitEmpty"`
}

type CustomerGroupDraft struct {
	// User-specific unique identifier for the customer group.
	Key       string        `json:"key,omitEmpty"`
	GroupName string        `json:"groupName"`
	Custom    *CustomFields `json:"custom,omitEmpty"`
}

type CustomerGroupPagedQueryResponse struct {
	Limit   int             `json:"limit"`
	Count   int             `json:"count"`
	Total   int             `json:"total,omitEmpty"`
	Offset  int             `json:"offset"`
	Results []CustomerGroup `json:"results"`
}

type CustomerGroupReference struct {
	Id  string         `json:"id"`
	Obj *CustomerGroup `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupReference) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

type CustomerGroupResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "customer-group", Alias: (*Alias)(&obj)})
}

type CustomerGroupUpdate struct {
	Version int                         `json:"version"`
	Actions []CustomerGroupUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CustomerGroupUpdate) UnmarshalJSON(data []byte) error {
	type Alias CustomerGroupUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type CustomerGroupUpdateAction interface{}

func mapDiscriminatorCustomerGroupUpdateAction(input interface{}) (CustomerGroupUpdateAction, error) {

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
	case "changeName":
		new := CustomerGroupChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CustomerGroupSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CustomerGroupSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CustomerGroupSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CustomerGroupChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CustomerGroupSetKeyAction struct {
	// User-specific unique identifier for the customer group.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CustomerGroupSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CustomerGroupSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
