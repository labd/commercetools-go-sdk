// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type SubRate struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TaxCategory struct {
	// The unique ID of the category.
	Id string `json:"id"`
	// The current version of the category.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy   *CreatedBy `json:"createdBy,omitEmpty"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitEmpty"`
	// The tax rates have unique IDs in the rates list
	Rates []TaxRate `json:"rates"`
	// User-specific unique identifier for the category.
	Key string `json:"key,omitEmpty"`
}

type TaxCategoryDraft struct {
	Name        string         `json:"name"`
	Description string         `json:"description,omitEmpty"`
	Rates       []TaxRateDraft `json:"rates"`
	Key         string         `json:"key,omitEmpty"`
}

type TaxCategoryPagedQueryResponse struct {
	Limit   int           `json:"limit"`
	Count   int           `json:"count"`
	Total   int           `json:"total,omitEmpty"`
	Offset  int           `json:"offset"`
	Results []TaxCategory `json:"results"`
}

type TaxCategoryReference struct {
	Id  string       `json:"id"`
	Obj *TaxCategory `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

type TaxCategoryResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "tax-category", Alias: (*Alias)(&obj)})
}

type TaxCategoryUpdate struct {
	Version int                       `json:"version"`
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type TaxCategoryUpdateAction interface{}

func mapDiscriminatorTaxCategoryUpdateAction(input interface{}) (TaxCategoryUpdateAction, error) {

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
	case "addTaxRate":
		new := TaxCategoryAddTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := TaxCategoryChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeTaxRate":
		new := TaxCategoryRemoveTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "replaceTaxRate":
		new := TaxCategoryReplaceTaxRateAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := TaxCategorySetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := TaxCategorySetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type TaxRate struct {
	// The ID is always set if the tax rate is part of a TaxCategory.
	// The external tax rates in a
	// Cart do not contain an `id`.
	Id   string `json:"id,omitEmpty"`
	Name string `json:"name"`
	// Percentage in the range of [0..1].
	// The sum of the amounts of all `subRates`, if there are any.
	Amount          float64 `json:"amount"`
	IncludedInPrice bool    `json:"includedInPrice"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// The state in the country
	State string `json:"state,omitEmpty"`
	// For countries (e.g.
	// the US) where the total tax is a combination of multiple taxes (e.g.
	// state and local taxes).
	SubRates []SubRate `json:"subRates,omitEmpty"`
}

type TaxRateDraft struct {
	Name string `json:"name"`
	// Percentage in the range of [0..1].
	// Must be supplied if no `subRates` are specified.
	// If `subRates` are specified
	// then the `amount` can be omitted or it must be the sum of the amounts of all `subRates`.
	Amount          float64 `json:"amount,omitEmpty"`
	IncludedInPrice bool    `json:"includedInPrice"`
	// A two-digit country code as per [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	Country string `json:"country"`
	// The state in the country
	State string `json:"state,omitEmpty"`
	// For countries (e.g.
	// the US) where the total tax is a combination of multiple taxes (e.g.
	// state and local taxes).
	SubRates []SubRate `json:"subRates,omitEmpty"`
}

type TaxCategoryAddTaxRateAction struct {
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryAddTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryAddTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TaxCategoryRemoveTaxRateAction struct {
	TaxRateId string `json:"taxRateId"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryReplaceTaxRateAction struct {
	TaxRateId string       `json:"taxRateId"`
	TaxRate   TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategorySetDescriptionAction struct {
	Description string `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type TaxCategorySetKeyAction struct {
	// If `key` is absent or `null`, it is removed if it exists.
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}
