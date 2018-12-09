// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type SubRate struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TaxCategory struct {
	Version        int       `json:"version"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	Rates          []TaxRate `json:"rates"`
	Name           string    `json:"name"`
	Key            string    `json:"key,omitempty"`
	Description    string    `json:"description,omitempty"`
}

type TaxCategoryAddTaxRateAction struct {
	TaxRate *TaxRateDraft `json:"taxRate"`
}

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

func (obj TaxCategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type TaxCategoryDraft struct {
	Rates       []TaxRateDraft `json:"rates"`
	Name        string         `json:"name"`
	Key         string         `json:"key,omitempty"`
	Description string         `json:"description,omitempty"`
}

type TaxCategoryPagedQueryResponse struct {
	Total   int           `json:"total,omitempty"`
	Offset  int           `json:"offset"`
	Count   int           `json:"count"`
	Results []TaxCategory `json:"results"`
}

type TaxCategoryReference struct {
	Key string       `json:"key,omitempty"`
	ID  string       `json:"id,omitempty"`
	Obj *TaxCategory `json:"obj,omitempty"`
}

func (obj TaxCategoryReference) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "tax-category", Alias: (*Alias)(&obj)})
}

type TaxCategoryRemoveTaxRateAction struct {
	TaxRateID string `json:"taxRateId"`
}

func (obj TaxCategoryRemoveTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryRemoveTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategoryReplaceTaxRateAction struct {
	TaxRateID string        `json:"taxRateId"`
	TaxRate   *TaxRateDraft `json:"taxRate"`
}

func (obj TaxCategoryReplaceTaxRateAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategoryReplaceTaxRateAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "replaceTaxRate", Alias: (*Alias)(&obj)})
}

type TaxCategorySetDescriptionAction struct {
	Description string `json:"description,omitempty"`
}

func (obj TaxCategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type TaxCategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

func (obj TaxCategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias TaxCategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type TaxCategoryUpdate struct {
	Version int                       `json:"version"`
	Actions []TaxCategoryUpdateAction `json:"actions"`
}

func (obj *TaxCategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias TaxCategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractTaxCategoryUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type TaxCategoryUpdateAction interface{}
type AbstractTaxCategoryUpdateAction struct{}

func AbstractTaxCategoryUpdateActionDiscriminatorMapping(input TaxCategoryUpdateAction) TaxCategoryUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addTaxRate":
		new := TaxCategoryAddTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := TaxCategoryChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeTaxRate":
		new := TaxCategoryRemoveTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "replaceTaxRate":
		new := TaxCategoryReplaceTaxRateAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := TaxCategorySetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := TaxCategorySetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}

type TaxRate struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	ID              string      `json:"id,omitempty"`
	Country         CountryCode `json:"country"`
	Amount          float64     `json:"amount"`
}

type TaxRateDraft struct {
	SubRates        []SubRate   `json:"subRates,omitempty"`
	State           string      `json:"state,omitempty"`
	Name            string      `json:"name"`
	IncludedInPrice bool        `json:"includedInPrice"`
	Country         CountryCode `json:"country"`
	Amount          float64     `json:"amount,omitempty"`
}
