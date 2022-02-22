// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Review struct {
	// The unique ID of the review.
	ID string `json:"id"`
	// The current version of the review.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-specific unique identifier for the review.
	Key             *string `json:"key,omitempty"`
	UniquenessValue *string `json:"uniquenessValue,omitempty"`
	Locale          *string `json:"locale,omitempty"`
	AuthorName      *string `json:"authorName,omitempty"`
	Title           *string `json:"title,omitempty"`
	Text            *string `json:"text,omitempty"`
	// Identifies the target of the review.
	// Can be a Product or a Channel
	Target interface{} `json:"target,omitempty"`
	// Indicates if this review is taken into account in the ratings statistics of the target.
	// A review is per default used in the statistics, unless the review is in a state that does not have the role `ReviewIncludedInStatistics`.
	// If the role of a State is modified after the calculation of this field, the calculation is not updated.
	IncludedInStatistics bool `json:"includedInStatistics"`
	// Number between -100 and 100 included.
	Rating *int            `json:"rating,omitempty"`
	State  *StateReference `json:"state,omitempty"`
	// The customer who created the review.
	Customer *CustomerReference `json:"customer,omitempty"`
	Custom   *CustomFields      `json:"custom,omitempty"`
}

type ReviewDraft struct {
	// User-specific unique identifier for the review.
	Key *string `json:"key,omitempty"`
	// If set, this value must be unique among reviews.
	// For example, if you want to have only one review per customer and per product, you can set the value to `customer's id` and `product's id`.
	UniquenessValue *string `json:"uniquenessValue,omitempty"`
	Locale          *string `json:"locale,omitempty"`
	AuthorName      *string `json:"authorName,omitempty"`
	Title           *string `json:"title,omitempty"`
	Text            *string `json:"text,omitempty"`
	// Identifies the target of the review.
	// Can be a Product or a Channel
	Target interface{}              `json:"target,omitempty"`
	State  *StateResourceIdentifier `json:"state,omitempty"`
	// Number between -100 and 100 included.
	// Rating of the targeted object, like a product.
	// This rating can represent the number of stars, or a percentage, or a like (+1)/dislike (-1)
	// A rating is used in the ratings statistics of the targeted object, unless the review is in a state that does not have the role `ReviewIncludedInStatistics`.
	Rating *int `json:"rating,omitempty"`
	// The customer who created the review.
	Customer *CustomerResourceIdentifier `json:"customer,omitempty"`
	Custom   *CustomFieldsDraft          `json:"custom,omitempty"`
}

type ReviewPagedQueryResponse struct {
	Limit   int      `json:"limit"`
	Count   int      `json:"count"`
	Total   *int     `json:"total,omitempty"`
	Offset  int      `json:"offset"`
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
	// The full distribution of the ratings.
	// The keys are the different ratings and the values are the count of reviews having this rating.
	// Only the used ratings appear in this object.
	RatingsDistribution interface{} `json:"ratingsDistribution"`
}

type ReviewReference struct {
	// Unique ID of the referenced resource.
	ID  string  `json:"id"`
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

type ReviewResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
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
	Version int                  `json:"version"`
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
	// If `authorName` is absent or `null`, this field will be removed if it exists.
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
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
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
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the CustomFields to this value.
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
	// The customer who created the review.
	// If `customer` is absent or `null`, this field will be removed if it exists.
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
	// If `key` is absent or `null`, this field will be removed if it exists.
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
	// If `locale` is absent or `null`, this field will be removed if it exists.
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

type ReviewSetRatingAction struct {
	// Number between -100 and 100 included.
	// If `rating` is absent or `null`, this field will be removed if it exists.
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
	// Identifies the target of the review.
	// Can be a Product or a Channel.
	// If `target` is absent or `null`, this field will be removed if it exists.
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
	// If `text` is absent or `null`, this field will be removed if it exists.
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
	// If `title` is absent or `null`, this field will be removed if it exists.
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

type ReviewTransitionStateAction struct {
	State StateResourceIdentifier `json:"state"`
	Force *bool                   `json:"force,omitempty"`
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
