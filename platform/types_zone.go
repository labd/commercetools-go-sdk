package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	A geographical location representing a country and optionally a state within this country.  A location can only be assigned to one Zone.
 */
type Location struct {
	// Country code of the geographic location.
	Country string `json:"country"`
	// State within the country.
	State *string `json:"state,omitempty"`
}

type Zone struct {
	// Unique identifier of the Zone.
	ID string `json:"id"`
	// Current version of the Zone.
	Version int `json:"version"`
	// Date and time (UTC) the Zone was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Zone was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Zone.
	Key *string `json:"key,omitempty"`
	// Name of the Zone.
	Name string `json:"name"`
	// Description of the Zone.
	Description *string `json:"description,omitempty"`
	// List of locations that belong to the Zone.
	Locations []Location `json:"locations"`
}

type ZoneDraft struct {
	// User-defined unique identifier for the Zone.
	Key *string `json:"key,omitempty"`
	// Name of the Zone.
	Name string `json:"name"`
	// Description of the Zone.
	Description *string `json:"description,omitempty"`
	// List of locations that belong to the Zone.
	Locations []Location `json:"locations"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneDraft) MarshalJSON() ([]byte, error) {
	type Alias ZoneDraft
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

	if raw["locations"] == nil {
		delete(raw, "locations")
	}

	return json.Marshal(raw)

}

/**
*	[PagedQueryResult](/general-concepts#pagedqueryresult) with `results` containing an array of [Zone](ctp:api:type:Zone).
*
 */
type ZonePagedQueryResponse struct {
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
	// [Zones](ctp:api:type:Zone) matching the query.
	Results []Zone `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Zone](ctp:api:type:Zone).
*
 */
type ZoneReference struct {
	// Unique identifier of the referenced [Zone](ctp:api:type:Zone).
	ID string `json:"id"`
	// Contains the representation of the expanded Zone. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Zones.
	Obj *Zone `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneReference) MarshalJSON() ([]byte, error) {
	type Alias ZoneReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "zone", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Zone](ctp:api:type:Zone). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ZoneResourceIdentifier struct {
	// Unique identifier of the referenced [Zone](ctp:api:type:Zone). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Zone](ctp:api:type:Zone). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ZoneResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "zone", Alias: (*Alias)(&obj)})
}

type ZoneUpdate struct {
	// Expected version of the Zone on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Zone.
	Actions []ZoneUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ZoneUpdate) UnmarshalJSON(data []byte) error {
	type Alias ZoneUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorZoneUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ZoneUpdateAction interface{}

func mapDiscriminatorZoneUpdateAction(input interface{}) (ZoneUpdateAction, error) {
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
	case "addLocation":
		obj := ZoneAddLocationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := ZoneChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeLocation":
		obj := ZoneRemoveLocationAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ZoneSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ZoneSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ZoneAddLocationAction struct {
	// Location to be added to the Zone.
	Location Location `json:"location"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneAddLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneAddLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addLocation", Alias: (*Alias)(&obj)})
}

type ZoneChangeNameAction struct {
	// New name of the Zone.
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ZoneRemoveLocationAction struct {
	// Location to be removed from the Zone.
	Location Location `json:"location"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneRemoveLocationAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneRemoveLocationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeLocation", Alias: (*Alias)(&obj)})
}

type ZoneSetDescriptionAction struct {
	// Description of the Zone.
	Description *string `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type ZoneSetKeyAction struct {
	// If `key` is absent or `null`, the existing key, if any, will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ZoneSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ZoneSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
