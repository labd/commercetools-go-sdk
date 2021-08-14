// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Location struct {
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	State   string `json:"state,omitEmpty"`
}

type Zone struct {
	// The unique ID of the zone.
	Id string `json:"id"`
	// The current version of the zone.
	Version        int             `json:"version"`
	CreatedAt      time.Time       `json:"createdAt"`
	LastModifiedAt time.Time       `json:"lastModifiedAt"`
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	CreatedBy      *CreatedBy      `json:"createdBy,omitEmpty"`
	// User-specific unique identifier for a zone.
	// Must be unique across a project.
	// The field can be reset using the Set Key UpdateAction.
	Key         string     `json:"key,omitEmpty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitEmpty"`
	Locations   []Location `json:"locations"`
}

type ZoneDraft struct {
	// User-specific unique identifier for a zone.
	// Must be unique across a project.
	// The field can be reset using the Set Key UpdateAction.
	Key         string     `json:"key,omitEmpty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitEmpty"`
	Locations   []Location `json:"locations"`
}

type ZonePagedQueryResponse struct {
	Limit   int    `json:"limit"`
	Count   int    `json:"count"`
	Total   int    `json:"total,omitEmpty"`
	Offset  int    `json:"offset"`
	Results []Zone `json:"results"`
}

type ZoneReference struct {
	Id  string `json:"id"`
	Obj *Zone  `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneReference) MarshalJSON() ([]byte, error) {
	type Alias ZoneReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "zone", Alias: (*Alias)(&obj)})
}

type ZoneResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ZoneResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "zone", Alias: (*Alias)(&obj)})
}

type ZoneUpdate struct {
	Version int                `json:"version"`
	Actions []ZoneUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ZoneUpdate) UnmarshalJSON(data []byte) error {
	type Alias ZoneUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type ZoneUpdateAction interface{}

func mapDiscriminatorZoneUpdateAction(input interface{}) (ZoneUpdateAction, error) {

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
	case "addLocation":
		new := ZoneAddLocationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := ZoneChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeLocation":
		new := ZoneRemoveLocationAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := ZoneSetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ZoneSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ZoneAddLocationAction struct {
	Location Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj ZoneChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ZoneRemoveLocationAction struct {
	Location Location `json:"location"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneRemoveLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneRemoveLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLocation", Alias: (*Alias)(&obj)})
}

type ZoneSetDescriptionAction struct {
	Description string `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ZoneSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ZoneSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
