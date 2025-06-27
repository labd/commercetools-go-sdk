package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	Supported interval units for [Recurrence Policies](ctp:api:type:RecurrencePolicy) which are used in [Recurring Orders](ctp:api:type:RecurringOrder).
*
 */
type IntervalUnit string

const (
	IntervalUnitDays   IntervalUnit = "Days"
	IntervalUnitWeeks  IntervalUnit = "Weeks"
	IntervalUnitMonths IntervalUnit = "Months"
)

type RecurrencePolicy struct {
	// Unique identifier of the RecurrencePolicy.
	ID string `json:"id"`
	// Current version of the RecurrencePolicy.
	Version int `json:"version"`
	// Date and time (UTC) the RecurrencePolicy was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the RecurrencePolicy was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// User-defined unique identifier of the RecurrencePolicy.
	Key string `json:"key"`
	// Name of the RecurrencePolicy.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the RecurrencePolicy.
	Description *LocalizedString `json:"description,omitempty"`
	// Schedule of the RecurrencePolicy.
	Schedule RecurrencePolicySchedule `json:"schedule"`
	// IDs and references that created the RecurrencePolicy.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// IDs and references that last modified the RecurrencePolicy.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurrencePolicy) UnmarshalJSON(data []byte) error {
	type Alias RecurrencePolicy
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Schedule != nil {
		var err error
		obj.Schedule, err = mapDiscriminatorRecurrencePolicySchedule(obj.Schedule)
		if err != nil {
			return err
		}
	}

	return nil
}

type RecurrencePolicyDraft struct {
	// User-defined unique identifier for the RecurrencePolicy.
	Key string `json:"key"`
	// Name of the RecurrencePolicy.
	Name *LocalizedString `json:"name,omitempty"`
	// Description of the RecurrencePolicy.
	Description *LocalizedString `json:"description,omitempty"`
	// Schedule where the recurrence is defined.
	Schedule RecurrencePolicyScheduleDraft `json:"schedule"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurrencePolicyDraft) UnmarshalJSON(data []byte) error {
	type Alias RecurrencePolicyDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Schedule != nil {
		var err error
		obj.Schedule, err = mapDiscriminatorRecurrencePolicyScheduleDraft(obj.Schedule)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [RecurrencePolicy](ctp:api:type:RecurrencePolicy).
*
 */
type RecurrencePolicyPagedQueryResponse struct {
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
	// [RecurrencePolicies](ctp:api:type:RecurrencePolicy) matching the query.
	Results []RecurrencePolicy `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [RecurrencePolicy](ctp:api:type:RecurrencePolicy).
*
 */
type RecurrencePolicyReference struct {
	// Unique identifier of the referenced [RecurrencePolicy](ctp:api:type:RecurrencePolicy).
	ID string `json:"id"`
	// Contains the representation of the expanded RecurrencePolicy.
	// Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for RecurrencePolicies.
	Obj *RecurrencePolicy `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicyReference) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicyReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurrence-policy", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [RecurrencePolicy](ctp:api:type:RecurrencePolicy). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type RecurrencePolicyResourceIdentifier struct {
	// Unique identifier of the referenced [RecurrencePolicy](ctp:api:type:RecurrencePolicy). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [RecurrencePolicy](ctp:api:type:RecurrencePolicy). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicyResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicyResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "recurrence-policy", Alias: (*Alias)(&obj)})
}

type RecurrencePolicySchedule interface{}

func mapDiscriminatorRecurrencePolicySchedule(input interface{}) (RecurrencePolicySchedule, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "dayOfMonth":
		obj := DayOfMonthSchedule{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "standard":
		obj := StandardSchedule{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Schedule of a [Recurring Order](ctp:api:type:RecurringOrder) that occurs on a specific day of each month.
*	This type is returned in the [RecurrencePolicy](ctp:api:type:RecurrencePolicy) for an active [Recurring Order](ctp:api:type:RecurringOrder).
*
 */
type DayOfMonthSchedule struct {
	// The day of the month when the [Recurring Order](ctp:api:type:RecurringOrder) is created.
	// If the value is greater than the number of days in a given month, the order is created on the last day of the month.
	Day int `json:"day"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj DayOfMonthSchedule) MarshalJSON() ([]byte, error) {
	type Alias DayOfMonthSchedule
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "dayOfMonth", Alias: (*Alias)(&obj)})
}

type RecurrencePolicyScheduleDraft interface{}

func mapDiscriminatorRecurrencePolicyScheduleDraft(input interface{}) (RecurrencePolicyScheduleDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "standard":
		obj := StandardScheduleDraft{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Determines the schedule for a [Recurring Order](ctp:api:type:RecurringOrder) to occur on a specific day of each month.
*
*	- Orders will be created even if the specified day is a weekend or holiday.
*	- To place orders on different dates within the same month (for example, the 1st and 15th), create separate [Recurring Orders](ctp:api:type:RecurringOrder)—each with its own schedule.
*
 */
type DayOfMonthScheduleDraft interface{}

func mapDiscriminatorDayOfMonthScheduleDraft(input interface{}) (DayOfMonthScheduleDraft, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["type"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'type'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {

	}
	return nil, nil
}

type RecurrencePolicyUpdate struct {
	// Expected version of the RecurrencePolicy on which the changes should be applied.
	// If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the RecurrencePolicy.
	Actions []RecurrencePolicyUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurrencePolicyUpdate) UnmarshalJSON(data []byte) error {
	type Alias RecurrencePolicyUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorRecurrencePolicyUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type RecurrencePolicyUpdateAction interface{}

func mapDiscriminatorRecurrencePolicyUpdateAction(input interface{}) (RecurrencePolicyUpdateAction, error) {
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
	case "setDescription":
		obj := RecurrencePolicySetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := RecurrencePolicySetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := RecurrencePolicySetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSchedule":
		obj := RecurrencePolicySetScheduleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		if obj.Schedule != nil {
			var err error
			obj.Schedule, err = mapDiscriminatorRecurrencePolicyScheduleDraft(obj.Schedule)
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Schedule of a [RecurringOrder](ctp:api:type:RecurringOrder) that occurs at a fixed interval—for example, every two weeks or every month.
*
 */
type StandardSchedule struct {
	// Number of intervals between orders.
	Value int `json:"value"`
	// Interval of this schedule.
	IntervalUnit IntervalUnit `json:"intervalUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandardSchedule) MarshalJSON() ([]byte, error) {
	type Alias StandardSchedule
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "standard", Alias: (*Alias)(&obj)})
}

/**
*	Determines the schedule for a [Recurring Order](ctp:api:type:RecurringOrder) to occur at a fixed interval—for example, every two weeks or every month.
*
 */
type StandardScheduleDraft struct {
	// Number of intervals between orders.
	Value int `json:"value"`
	// Interval for this schedule.
	IntervalUnit IntervalUnit `json:"intervalUnit"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj StandardScheduleDraft) MarshalJSON() ([]byte, error) {
	type Alias StandardScheduleDraft
	return json.Marshal(struct {
		Action string `json:"type"`
		*Alias
	}{Action: "standard", Alias: (*Alias)(&obj)})
}

type RecurrencePolicySetDescriptionAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type RecurrencePolicySetKeyAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type RecurrencePolicySetNameAction struct {
	// Value to set.
	// If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicySetNameAction) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicySetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type RecurrencePolicySetScheduleAction struct {
	// Schedule where the recurrence is defined.
	Schedule RecurrencePolicyScheduleDraft `json:"schedule"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *RecurrencePolicySetScheduleAction) UnmarshalJSON(data []byte) error {
	type Alias RecurrencePolicySetScheduleAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Schedule != nil {
		var err error
		obj.Schedule, err = mapDiscriminatorRecurrencePolicyScheduleDraft(obj.Schedule)
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj RecurrencePolicySetScheduleAction) MarshalJSON() ([]byte, error) {
	type Alias RecurrencePolicySetScheduleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSchedule", Alias: (*Alias)(&obj)})
}
