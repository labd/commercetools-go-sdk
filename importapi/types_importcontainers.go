package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	The strategy of the retention policy. Used to determine how the ImportContainer should be retained.
 */
type StrategyEnum string

const (
	StrategyEnumTtl StrategyEnum = "ttl"
)

/**
*	The retention policy of the ImportContainer. If not set, the ImportContainer does not expire.
 */
type RetentionPolicy interface{}

func mapDiscriminatorRetentionPolicy(input interface{}) (RetentionPolicy, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["strategy"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'strategy'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "ttl":
		obj := TimeToLiveRetentionPolicy{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type TimeToLiveConfig struct {
	// The time to live of the ImportContainer. Used to generate the `expiresAt` value of the ImportContainer. The lowest accepted value is `1h` and the highest accepted value is `30d`.
	TimeToLive string `json:"timeToLive"`
}

/**
*	Set a time to live retention policy for the ImportContainer.
 */
type TimeToLiveRetentionPolicy struct {
	// The configuration of the time to live retention policy.
	Config TimeToLiveConfig `json:"config"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj TimeToLiveRetentionPolicy) MarshalJSON() ([]byte, error) {
	type Alias TimeToLiveRetentionPolicy
	return json.Marshal(struct {
		Action string `json:"strategy"`
		*Alias
	}{Action: "ttl", Alias: (*Alias)(&obj)})
}

/**
*	Contains the resources to be imported. Unless `resourceType` is specified, the ImportContainer can import all of the supported [ImportResourceTypes](ctp:import:type:ImportResourceType).
*
 */
type ImportContainer struct {
	// User-defined unique identifier of the ImportContainer.
	Key string `json:"key"`
	// The [resource type](ctp:import:type:ImportResourceType) the ImportContainer supports. If not present, the ImportContainer can import all of the supported [ImportResourceTypes](ctp:import:type:ImportResourceType).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
	// Current version of the ImportContainer.
	Version int `json:"version"`
	// The retention policy of the ImportContainer.
	RetentionPolicy RetentionPolicy `json:"retentionPolicy,omitempty"`
	// Date and time (UTC) the ImportContainer was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ImportContainer was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Date and time (UTC) the ImportContainer is automatically deleted. Only present if a `retentionPolicy` is set. ImportContainers without `expiresAt` are permanent until [manually deleted](#delete-importcontainer).
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImportContainer) UnmarshalJSON(data []byte) error {
	type Alias ImportContainer
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.RetentionPolicy != nil {
		var err error
		obj.RetentionPolicy, err = mapDiscriminatorRetentionPolicy(obj.RetentionPolicy)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	The representation sent to the server to create an [ImportContainer](#importcontainer).
*
 */
type ImportContainerDraft struct {
	// User-defined unique identifier of the ImportContainer.
	Key string `json:"key"`
	// The resource type the ImportContainer will accept.
	// If not specified, the ImportContainer can import all of the supported ImportResourceTypes.
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
	// Set a retention policy to automatically delete the ImportContainer after a defined period.
	RetentionPolicy RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ImportContainerDraft) UnmarshalJSON(data []byte) error {
	type Alias ImportContainerDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.RetentionPolicy != nil {
		var err error
		obj.RetentionPolicy, err = mapDiscriminatorRetentionPolicy(obj.RetentionPolicy)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	The representation sent to the server when updating an ImportContainer.
*
 */
type ImportContainerUpdateDraft struct {
	// Current version of the ImportContainer.
	Version int `json:"version"`
	// The [resource type](ctp:import:type:ImportResourceType) to be imported.
	// If not given, the ImportContainer is able to import all of the supported [ImportResourceTypes](ctp:import:type:ImportResourceType).
	ResourceType *ImportResourceType `json:"resourceType,omitempty"`
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ImportContainer](ctp:import:type:ImportContainer).
*
 */
type ImportContainerPagedResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	Total int `json:"total"`
	// [ImportContainers](ctp:import:type:ImportContainer) matching the query.
	Results []ImportContainer `json:"results"`
}
