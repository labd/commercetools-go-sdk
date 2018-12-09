// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type CartClassificationType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

func (obj CartClassificationType) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

type CartScoreType struct{}

func (obj CartScoreType) MarshalJSON() ([]byte, error) {
	type Alias CartScoreType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

type CartValueType struct{}

func (obj CartValueType) MarshalJSON() ([]byte, error) {
	type Alias CartValueType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

type Project struct {
	Version               int                   `json:"version"`
	TrialUntil            string                `json:"trialUntil,omitempty"`
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
	Name                  string                `json:"name"`
	Messages              *MessageConfiguration `json:"messages"`
	Languages             []Locale              `json:"languages"`
	Key                   string                `json:"key"`
	Currencies            []CurrencyCode        `json:"currencies"`
	CreatedAt             time.Time             `json:"createdAt"`
	Countries             []CountryCode         `json:"countries"`
}

func (obj *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		obj.ShippingRateInputType = AbstractShippingRateInputTypeDiscriminatorMapping(obj.ShippingRateInputType)
	}

	return nil
}

type ProjectChangeCountriesAction struct {
	Countries []CountryCode `json:"countries"`
}

func (obj ProjectChangeCountriesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCountriesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCountries", Alias: (*Alias)(&obj)})
}

type ProjectChangeCurrenciesAction struct {
	Currencies []CurrencyCode `json:"currencies"`
}

func (obj ProjectChangeCurrenciesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCurrenciesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCurrencies", Alias: (*Alias)(&obj)})
}

type ProjectChangeLanguagesAction struct {
	Languages []Locale `json:"languages"`
}

func (obj ProjectChangeLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLanguages", Alias: (*Alias)(&obj)})
}

type ProjectChangeMessagesConfigurationAction struct {
	MessagesConfiguration *MessageConfigurationDraft `json:"messagesConfiguration"`
}

func (obj ProjectChangeMessagesConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesConfiguration", Alias: (*Alias)(&obj)})
}

type ProjectChangeMessagesEnabledAction struct {
	MessagesEnabled bool `json:"messagesEnabled"`
}

func (obj ProjectChangeMessagesEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesEnabled", Alias: (*Alias)(&obj)})
}

type ProjectChangeNameAction struct {
	Name string `json:"name"`
}

func (obj ProjectChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type ProjectSetShippingRateInputTypeAction struct {
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
}

func (obj ProjectSetShippingRateInputTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetShippingRateInputTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInputType", Alias: (*Alias)(&obj)})
}
func (obj *ProjectSetShippingRateInputTypeAction) UnmarshalJSON(data []byte) error {
	type Alias ProjectSetShippingRateInputTypeAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		obj.ShippingRateInputType = AbstractShippingRateInputTypeDiscriminatorMapping(obj.ShippingRateInputType)
	}

	return nil
}

type ProjectUpdate struct {
	Version int                   `json:"version"`
	Actions []ProjectUpdateAction `json:"actions"`
}

func (obj *ProjectUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProjectUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractProjectUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type ProjectUpdateAction interface{}
type AbstractProjectUpdateAction struct{}

func AbstractProjectUpdateActionDiscriminatorMapping(input ProjectUpdateAction) ProjectUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "changeCountries":
		new := ProjectChangeCountriesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeCurrencies":
		new := ProjectChangeCurrenciesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeLanguages":
		new := ProjectChangeLanguagesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeMessagesConfiguration":
		new := ProjectChangeMessagesConfigurationAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeMessagesEnabled":
		new := ProjectChangeMessagesEnabledAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := ProjectChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setShippingRateInputType":
		new := ProjectSetShippingRateInputTypeAction{}
		mapstructure.Decode(input, &new)
		if new.ShippingRateInputType != nil {
			new.ShippingRateInputType = AbstractShippingRateInputTypeDiscriminatorMapping(new.ShippingRateInputType)
		}

		return new
	}
	return nil
}

type ShippingRateInputType interface{}
type AbstractShippingRateInputType struct{}

func AbstractShippingRateInputTypeDiscriminatorMapping(input ShippingRateInputType) ShippingRateInputType {
	discriminator := input.(map[string]interface{})["type"]
	switch discriminator {
	case "CartClassification":
		new := CartClassificationType{}
		mapstructure.Decode(input, &new)
		return new
	case "CartScore":
		new := CartScoreType{}
		mapstructure.Decode(input, &new)
		return new
	case "CartValue":
		new := CartValueType{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
