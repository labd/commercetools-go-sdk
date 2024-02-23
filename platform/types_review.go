package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Review struct {
	// Unique identifier of the Review.
	ID string `json:"id"`
	// Current version of the Review.
	Version int `json:"version"`
	// Date and time (UTC) the Review was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Review was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/general-concepts#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Review.
	Key *string `json:"key,omitempty"`
	// Must be unique among Reviews. For example, if this value is set to Customer `id` + Product `id`, only one Review per Customer and per Product is allowed.
	UniquenessValue *string `json:"uniquenessValue,omitempty"`
	// Language in which the content of the Review is written.
	Locale *string `json:"locale,omitempty"`
	// Name of the author.
	AuthorName *string `json:"authorName,omitempty"`
	// Title of the Review.
	Title *string `json:"title,omitempty"`
	// Content of the Review.
	Text *string `json:"text,omitempty"`
	// Identifies the target of the Review. Can be a [Product](ctp:api:type:Product) or a [Channel](ctp:api:type:Channel), specified as [ProductReference](ctp:api:type:ProductReference) or [ChannelReference](ctp:api:type:ChannelReference), respectively.
	Target interface{} `json:"target,omitempty"`
	// Indicates if this Review is taken into account in the ratings statistics of the target.
	// A Review is per default used in the statistics, unless the Review is in a state that does not have the [role](ctp:api:type:StateRoleEnum) `ReviewIncludedInStatistics`.
	// If the role of a [State](ctp:api:type:State) is modified after the calculation of this field, the calculation is not updated.
	IncludedInStatistics bool `json:"includedInStatistics"`
	// Rating of the Product or Channel.
	Rating *int `json:"rating,omitempty"`
	// State of the Review. Used for approval processes, see [Review approval process](/../tutorials/review-ratings#review-approval-process) for details.
	State *StateReference `json:"state,omitempty"`
	// Customer who created the Review.
	Customer *CustomerReference `json:"customer,omitempty"`
	// Custom Fields of the Review.
	Custom *CustomFields `json:"custom,omitempty"`
}

/**
*	When creating a new Review, at least one of `title`, `text` or `rating` should be set.
*
 */
type ReviewDraft struct {
	// User-defined unique identifier for the Review.
	Key *string `json:"key,omitempty"`
	// If set, this value must be unique among Reviews.
	// For example, if you want to have only one Review per Customer and per Product, you can set the value to Customer `id` + Product `id`.
	UniquenessValue *string `json:"uniquenessValue,omitempty"`
	// Language in which the content of the Review is written.
	Locale *string `json:"locale,omitempty"`
	// Name of the author.
	AuthorName *string `json:"authorName,omitempty"`
	// Title of the Review.
	Title *string `json:"title,omitempty"`
	// Content of the Review.
	Text *string `json:"text,omitempty"`
	// Identifies the target of the Review. Can be a [Product](ctp:api:type:Product) or a [Channel](ctp:api:type:Channel), specified as [ProductResourceIdentifier](ctp:api:type:ProductResourceIdentifier) or [ChannelResourceIdentifier](ctp:api:type:ChannelResourceIdentifier), respectively.
	Target interface{} `json:"target,omitempty"`
	// State of the Review. Used for approval processes, see [Review approval process](/../tutorials/review-ratings#review-approval-process) for details.
	State *StateResourceIdentifier `json:"state,omitempty"`
	// Rating of the targeted Product or Channel.
	// This rating can represent the number of stars, a percentage, or a like (+1)/dislike (-1).
	// A rating is used in the ratings statistics of the targeted object, unless the Review is in a State that does not have the role `ReviewIncludedInStatistics`.
	Rating *int `json:"rating,omitempty"`
	// Customer who created the Review.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	// Custom Fields for the Review.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

type ReviewPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [Reviews](ctp:api:type:Review) matching the query.
	Results []Review `json:"results"`
}

type ReviewRatingStatistics struct {
	// Average rating of one target
	// This number is rounded with 5 decimals.
	AverageRating float64 `json:"averageRating"`
	// Highest rating of one target
	HighestRating float64 `json:"highestRating"`
	// Lowest rating of one target
	LowestRating float64 `json:"lowestRating"`
	// Number of ratings taken into account
	Count int `json:"count"`
	// Full distribution of the ratings.
	// The keys are the different ratings and the values are the count of reviews having this rating.
	// Only the used ratings appear in this object.
	RatingsDistribution interface{} `json:"ratingsDistribution"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Review](ctp:api:type:Review).
*
 */
type ReviewReference struct {
	// Unique identifier of the referenced [Review](ctp:api:type:Review).
	ID string `json:"id"`
	// Contains the representation of the expanded Review. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Reviews.
	Obj *Review `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewReference) MarshalJSON() ([]byte, error) {
	type Alias ReviewReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "review", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Review](ctp:api:type:Review). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type ReviewResourceIdentifier struct {
	// Unique identifier of the referenced [Review](ctp:api:type:Review). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Review](ctp:api:type:Review). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ReviewResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "review", Alias: (*Alias)(&obj)})
}

type ReviewUpdate struct {
	// The expected version of the review on which the changes should be applied. If the expected version does not match the actual version, a 409 Conflict will be returned.
	Version int `json:"version"`
	// The list of update actions to be performed on the review.
	Actions []ReviewUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ReviewUpdate) UnmarshalJSON(data []byte) error {
	type Alias ReviewUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorReviewUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type ReviewUpdateAction interface{}

func mapDiscriminatorReviewUpdateAction(input interface{}) (ReviewUpdateAction, error) {
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
	case "setAuthorName":
		obj := ReviewSetAuthorNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := ReviewSetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := ReviewSetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomer":
		obj := ReviewSetCustomerAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := ReviewSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setLocale":
		obj := ReviewSetLocaleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRating":
		obj := ReviewSetRatingAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTarget":
		obj := ReviewSetTargetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setText":
		obj := ReviewSetTextAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setTitle":
		obj := ReviewSetTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "transitionState":
		obj := ReviewTransitionStateAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type ReviewSetAuthorNameAction struct {
	// Value to set. If empty, any existing value will be removed.
	AuthorName *string `json:"authorName,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetAuthorNameAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetAuthorNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorName", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the Review with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Review.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Review.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomerAction struct {
	// Value to set. If empty, any existing value will be removed.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type ReviewSetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ReviewSetLocaleAction struct {
	// Value to set. If empty, any existing value will be removed.
	Locale *string `json:"locale,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

/**
*	This update action produces the [ReviewRatingSet](ctp:api:type:ReviewRatingSetMessage) Message.
*
 */
type ReviewSetRatingAction struct {
	// Value to set. If empty, any existing value will be removed.
	Rating *int `json:"rating,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetRatingAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetRatingAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRating", Alias: (*Alias)(&obj)})
}

type ReviewSetTargetAction struct {
	// Value to set, specified as [ProductResourceIdentifier](ctp:api:type:ProductResourceIdentifier) or [ChannelResourceIdentifier](ctp:api:type:ChannelResourceIdentifier), respectively. If empty, any existing value will be removed.
	Target interface{} `json:"target"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetTargetAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTarget", Alias: (*Alias)(&obj)})
}

type ReviewSetTextAction struct {
	// Value to set. If empty, any existing value will be removed.
	Text *string `json:"text,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetTextAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setText", Alias: (*Alias)(&obj)})
}

type ReviewSetTitleAction struct {
	// Value to set. If empty, any existing value will be removed.
	Title *string `json:"title,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

/**
*	Transition to a new State. This update action produces the [Review State Transition](ctp:api:type:ReviewStateTransitionMessage) Message.
*
 */
type ReviewTransitionStateAction struct {
	// Value to set. If there is no State yet, the new State must be an initial State. If the existing State has `transitions` set, there must be a direct transition to the new State. If `transitions` is not set, no validation is performed. If the new State does not have the [role](ctp:api:type:StateRoleEnum) `ReviewIncludedInStatistics`, the Review is not taken into account in the ratings statistics of the target.
	State StateResourceIdentifier `json:"state"`
	// Switch validations on or off.
	Force *bool `json:"force,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ReviewTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
