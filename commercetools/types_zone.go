// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Location struct {
	State   string      `json:"state,omitempty"`
	Country CountryCode `json:"country"`
}

type Zone struct {
	Version        int        `json:"version"`
	LastModifiedAt time.Time  `json:"lastModifiedAt"`
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"createdAt"`
	Name           string     `json:"name"`
	Locations      []Location `json:"locations"`
	Key            string     `json:"key,omitempty"`
	Description    string     `json:"description,omitempty"`
}

type ZoneAddLocationAction struct {
	Location *Location `json:"location"`
}

func (obj ZoneAddLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneAddLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocation", Alias: (*Alias)(&obj)})
}

type ZoneChangeNameAction struct {
	Name string `json:"name"`
}

func (obj ZoneChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ZoneDraft struct {
	Name        string     `json:"name"`
	Locations   []Location `json:"locations"`
	Key         string     `json:"key,omitempty"`
	Description string     `json:"description,omitempty"`
}

type ZonePagedQueryResponse struct {
	Total   int    `json:"total,omitempty"`
	Offset  int    `json:"offset"`
	Count   int    `json:"count"`
	Results []Zone `json:"results"`
}

type ZoneReference struct {
	Key string `json:"key,omitempty"`
	ID  string `json:"id,omitempty"`
	Obj *Zone  `json:"obj,omitempty"`
}

func (obj ZoneReference) MarshalJSON() ([]byte, error) {
	type Alias ZoneReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "zone", Alias: (*Alias)(&obj)})
}

type ZoneRemoveLocationAction struct {
	Location *Location `json:"location"`
}

func (obj ZoneRemoveLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneRemoveLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLocation", Alias: (*Alias)(&obj)})
}

type ZoneSetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

func (obj ZoneSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ZoneSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ZoneSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ZoneUpdate struct {
	Version int                `json:"version"`
	Actions []ZoneUpdateAction `json:"actions"`
}

func (obj *ZoneUpdate) UnmarshalJSON(data []byte) error {
	type Alias ZoneUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractZoneUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ZoneUpdateAction interface{}
type AbstractZoneUpdateAction struct{}

func AbstractZoneUpdateActionDiscriminatorMapping(input ZoneUpdateAction) ZoneUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addLocation":
		new := ZoneAddLocationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ZoneChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeLocation":
		new := ZoneRemoveLocationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := ZoneSetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ZoneSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
