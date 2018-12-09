// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Review struct {
	Version              int                `json:"version"`
	LastModifiedAt       time.Time          `json:"lastModifiedAt"`
	ID                   string             `json:"id"`
	CreatedAt            time.Time          `json:"createdAt"`
	UniquenessValue      string             `json:"uniquenessValue,omitempty"`
	Title                string             `json:"title,omitempty"`
	Text                 string             `json:"text,omitempty"`
	Target               Reference          `json:"target,omitempty"`
	State                *StateReference    `json:"state,omitempty"`
	Rating               float64            `json:"rating,omitempty"`
	Locale               string             `json:"locale,omitempty"`
	Key                  string             `json:"key,omitempty"`
	IncludedInStatistics bool               `json:"includedInStatistics"`
	Customer             *CustomerReference `json:"customer,omitempty"`
	Custom               *CustomFields      `json:"custom,omitempty"`
	AuthorName           string             `json:"authorName,omitempty"`
}

func (obj *Review) UnmarshalJSON(data []byte) error {
	type Alias Review
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewDraft struct {
	UniquenessValue string              `json:"uniquenessValue,omitempty"`
	Title           string              `json:"title,omitempty"`
	Text            string              `json:"text,omitempty"`
	Target          Reference           `json:"target,omitempty"`
	State           *ResourceIdentifier `json:"state,omitempty"`
	Rating          float64             `json:"rating,omitempty"`
	Locale          string              `json:"locale,omitempty"`
	Key             string              `json:"key,omitempty"`
	Customer        *CustomerReference  `json:"customer,omitempty"`
	Custom          *CustomFieldsDraft  `json:"custom,omitempty"`
	AuthorName      string              `json:"authorName,omitempty"`
}

func (obj *ReviewDraft) UnmarshalJSON(data []byte) error {
	type Alias ReviewDraft
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewPagedQueryResponse struct {
	Total   int      `json:"total,omitempty"`
	Offset  int      `json:"offset"`
	Count   int      `json:"count"`
	Results []Review `json:"results"`
}

type ReviewRatingStatistics struct {
	RatingsDistribution interface{} `json:"ratingsDistribution"`
	LowestRating        float64     `json:"lowestRating"`
	HighestRating       float64     `json:"highestRating"`
	Count               int         `json:"count"`
	AverageRating       float64     `json:"averageRating"`
}

type ReviewReference struct {
	Key string  `json:"key,omitempty"`
	ID  string  `json:"id,omitempty"`
	Obj *Review `json:"obj,omitempty"`
}

func (obj ReviewReference) MarshalJSON() ([]byte, error) {
	type Alias ReviewReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "review", Alias: (*Alias)(&obj)})
}

type ReviewSetAuthorNameAction struct {
	AuthorName string `json:"authorName,omitempty"`
}

func (obj ReviewSetAuthorNameAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetAuthorNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAuthorName", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj ReviewSetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

func (obj ReviewSetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type ReviewSetCustomerAction struct {
	Customer *CustomerReference `json:"customer,omitempty"`
}

func (obj ReviewSetCustomerAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetCustomerAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomer", Alias: (*Alias)(&obj)})
}

type ReviewSetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj ReviewSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type ReviewSetLocaleAction struct {
	Locale string `json:"locale,omitempty"`
}

func (obj ReviewSetLocaleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetLocaleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setLocale", Alias: (*Alias)(&obj)})
}

type ReviewSetRatingAction struct {
	Rating float64 `json:"rating,omitempty"`
}

func (obj ReviewSetRatingAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetRatingAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRating", Alias: (*Alias)(&obj)})
}

type ReviewSetTargetAction struct {
	Target Reference `json:"target,omitempty"`
}

func (obj ReviewSetTargetAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTargetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTarget", Alias: (*Alias)(&obj)})
}
func (obj *ReviewSetTargetAction) UnmarshalJSON(data []byte) error {
	type Alias ReviewSetTargetAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.Target != nil {
		obj.Target = AbstractReferenceDiscriminatorMapping(obj.Target)
	}

	return nil
}

type ReviewSetTextAction struct {
	Text string `json:"text,omitempty"`
}

func (obj ReviewSetTextAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTextAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setText", Alias: (*Alias)(&obj)})
}

type ReviewSetTitleAction struct {
	Title string `json:"title,omitempty"`
}

func (obj ReviewSetTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewSetTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setTitle", Alias: (*Alias)(&obj)})
}

type ReviewTransitionStateAction struct {
	State *StateReference `json:"state"`
	Force bool            `json:"force,omitempty"`
}

func (obj ReviewTransitionStateAction) MarshalJSON() ([]byte, error) {
	type Alias ReviewTransitionStateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "transitionState", Alias: (*Alias)(&obj)})
}

type ReviewUpdate struct {
	Version int                  `json:"version"`
	Actions []ReviewUpdateAction `json:"actions"`
}

func (obj *ReviewUpdate) UnmarshalJSON(data []byte) error {
	type Alias ReviewUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractReviewUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ReviewUpdateAction interface{}
type AbstractReviewUpdateAction struct{}

func AbstractReviewUpdateActionDiscriminatorMapping(input ReviewUpdateAction) ReviewUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "setAuthorName":
		new := ReviewSetAuthorNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := ReviewSetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := ReviewSetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomer":
		new := ReviewSetCustomerAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := ReviewSetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setLocale":
		new := ReviewSetLocaleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setRating":
		new := ReviewSetRatingAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTarget":
		new := ReviewSetTargetAction{}
		mapstructure.Decode(input, &new)
		if new.Target != nil {
			new.Target = AbstractReferenceDiscriminatorMapping(new.Target)
		}

		return new
	case "setText":
		new := ReviewSetTextAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setTitle":
		new := ReviewSetTitleAction{}
		mapstructure.Decode(input, &new)
		return new
	case "transitionState":
		new := ReviewTransitionStateAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
