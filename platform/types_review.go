// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Review struct {
	// The unique ID of the review.
	Id string `json:"id"`
	// The current version of the review.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy `json:"createdBy,omitEmpty"`
	// User-specific unique identifier for the review.
	Key             string `json:"key,omitEmpty"`
	UniquenessValue string `json:"uniquenessValue,omitEmpty"`
	Locale          string `json:"locale,omitEmpty"`
	AuthorName      string `json:"authorName,omitEmpty"`
	Title           string `json:"title,omitEmpty"`
	Text            string `json:"text,omitEmpty"`
	// Identifies the target of the review.
	// Can be a Product or a Channel
	Target Reference `json:"target,omitEmpty"`
	// Indicates if this review is taken into account in the ratings statistics of the target.
	// A review is per default used in the statistics, unless the review is in a state that does not have the role `ReviewIncludedInStatistics`.
	// If the role of a State is modified after the calculation of this field, the calculation is not updated.
	IncludedInStatistics bool `json:"includedInStatistics"`
	// Number between -100 and 100 included.
	Rating int            `json:"rating,omitEmpty"`
	State  StateReference `json:"state,omitEmpty"`
	// The customer who created the review.
	Customer CustomerReference `json:"customer,omitEmpty"`
	Custom   *CustomFields     `json:"custom,omitEmpty"`
}

type ReviewDraft struct {
	// User-specific unique identifier for the review.
	Key string `json:"key,omitEmpty"`
	// If set, this value must be unique among reviews.
	// For example, if you want to have only one review per customer and per product, you can set the value to `customer's id` and `product's id`.
	UniquenessValue string `json:"uniquenessValue,omitEmpty"`
	Locale          string `json:"locale,omitEmpty"`
	AuthorName      string `json:"authorName,omitEmpty"`
	Title           string `json:"title,omitEmpty"`
	Text            string `json:"text,omitEmpty"`
	// Identifies the target of the review.
	// Can be a Product or a Channel
	Target ResourceIdentifier      `json:"target,omitEmpty"`
	State  StateResourceIdentifier `json:"state,omitEmpty"`
	// Number between -100 and 100 included.
	// Rating of the targeted object, like a product.
	// This rating can represent the number of stars, or a percentage, or a like (+1)/dislike (-1)
	// A rating is used in the ratings statistics of the targeted object, unless the review is in a state that does not have the role `ReviewIncludedInStatistics`.
	Rating int `json:"rating,omitEmpty"`
	// The customer who created the review.
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
	Custom   *CustomFieldsDraft         `json:"custom,omitEmpty"`
}

type ReviewPagedQueryResponse struct {
	Limit   int      `json:"limit"`
	Count   int      `json:"count"`
	Total   int      `json:"total,omitEmpty"`
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
	Id  string  `json:"id"`
	Obj *Review `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewReference) MarshalJSON() ([]byte, error) {
	type Alias ReviewReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "review", Alias: (*Alias)(&obj)})
}

type ReviewResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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

	return nil
}

type ReviewUpdateAction interface{}

func mapDiscriminatorReviewUpdateAction(input interface{}) (ReviewUpdateAction, error) {

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
	case "setAuthorName":
		new := ReviewSetAuthorNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := ReviewSetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := ReviewSetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomer":
		new := ReviewSetCustomerAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := ReviewSetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setLocale":
		new := ReviewSetLocaleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setRating":
		new := ReviewSetRatingAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTarget":
		new := ReviewSetTargetAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setText":
		new := ReviewSetTextAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setTitle":
		new := ReviewSetTitleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "transitionState":
		new := ReviewTransitionStateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type ReviewSetAuthorNameAction struct {
	// If `authorName` is absent or `null`, this field will be removed if it exists.
	AuthorName string `json:"authorName,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetAuthorNameAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetAuthorNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorName", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomTypeAction struct {
	// If absent, the custom type and any existing custom fields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type.
	// Sets the CustomFields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	Customer CustomerResourceIdentifier `json:"customer,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type ReviewSetKeyAction struct {
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ReviewSetLocaleAction struct {
	// If `locale` is absent or `null`, this field will be removed if it exists.
	Locale string `json:"locale,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	Rating int `json:"rating,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	Target ResourceIdentifier `json:"target"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTargetAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTarget", Alias: (*Alias)(&obj)})
}

type ReviewSetTextAction struct {
	// If `text` is absent or `null`, this field will be removed if it exists.
	Text string `json:"text,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTextAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setText", Alias: (*Alias)(&obj)})
}

type ReviewSetTitleAction struct {
	// If `title` is absent or `null`, this field will be removed if it exists.
	Title string `json:"title,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type ReviewTransitionStateAction struct {
	State StateResourceIdentifier `json:"state"`
	Force bool                    `json:"force,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj ReviewTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}
