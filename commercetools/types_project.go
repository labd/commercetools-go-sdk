// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// ProjectUpdateAction uses action as discriminator attribute
type ProjectUpdateAction interface{}

func mapDiscriminatorProjectUpdateAction(input interface{}) ProjectUpdateAction {
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
			new.ShippingRateInputType = mapDiscriminatorShippingRateInputType(new.ShippingRateInputType)
		}

		return new
	}
	return nil
}

// ShippingRateInputType uses type as discriminator attribute
type ShippingRateInputType interface{}

func mapDiscriminatorShippingRateInputType(input interface{}) ShippingRateInputType {
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

// CartClassificationType implements the interface ShippingRateInputType
type CartClassificationType struct {
	Values []CustomFieldLocalizedEnumValue `json:"values"`
}

// MarshalJSON override to set the discriminator value
func (obj CartClassificationType) MarshalJSON() ([]byte, error) {
	type Alias CartClassificationType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartClassification", Alias: (*Alias)(&obj)})
}

// CartScoreType implements the interface ShippingRateInputType
type CartScoreType struct{}

// MarshalJSON override to set the discriminator value
func (obj CartScoreType) MarshalJSON() ([]byte, error) {
	type Alias CartScoreType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartScore", Alias: (*Alias)(&obj)})
}

// CartValueType implements the interface ShippingRateInputType
type CartValueType struct{}

// MarshalJSON override to set the discriminator value
func (obj CartValueType) MarshalJSON() ([]byte, error) {
	type Alias CartValueType
	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{Type: "CartValue", Alias: (*Alias)(&obj)})
}

// Project is a standalone struct
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

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		obj.ShippingRateInputType = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
	}

	return nil
}

// ProjectChangeCountriesAction implements the interface ProjectUpdateAction
type ProjectChangeCountriesAction struct {
	Countries []CountryCode `json:"countries"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeCountriesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCountriesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCountries", Alias: (*Alias)(&obj)})
}

// ProjectChangeCurrenciesAction implements the interface ProjectUpdateAction
type ProjectChangeCurrenciesAction struct {
	Currencies []CurrencyCode `json:"currencies"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeCurrenciesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeCurrenciesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeCurrencies", Alias: (*Alias)(&obj)})
}

// ProjectChangeLanguagesAction implements the interface ProjectUpdateAction
type ProjectChangeLanguagesAction struct {
	Languages []Locale `json:"languages"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeLanguagesAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeLanguagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeLanguages", Alias: (*Alias)(&obj)})
}

// ProjectChangeMessagesConfigurationAction implements the interface ProjectUpdateAction
type ProjectChangeMessagesConfigurationAction struct {
	MessagesConfiguration *MessageConfigurationDraft `json:"messagesConfiguration"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeMessagesConfigurationAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesConfigurationAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesConfiguration", Alias: (*Alias)(&obj)})
}

// ProjectChangeMessagesEnabledAction implements the interface ProjectUpdateAction
type ProjectChangeMessagesEnabledAction struct {
	MessagesEnabled bool `json:"messagesEnabled"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeMessagesEnabledAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeMessagesEnabledAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeMessagesEnabled", Alias: (*Alias)(&obj)})
}

// ProjectChangeNameAction implements the interface ProjectUpdateAction
type ProjectChangeNameAction struct {
	Name string `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// ProjectSetShippingRateInputTypeAction implements the interface ProjectUpdateAction
type ProjectSetShippingRateInputTypeAction struct {
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj ProjectSetShippingRateInputTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProjectSetShippingRateInputTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setShippingRateInputType", Alias: (*Alias)(&obj)})
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectSetShippingRateInputTypeAction) UnmarshalJSON(data []byte) error {
	type Alias ProjectSetShippingRateInputTypeAction
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	if obj.ShippingRateInputType != nil {
		obj.ShippingRateInputType = mapDiscriminatorShippingRateInputType(obj.ShippingRateInputType)
	}

	return nil
}

// ProjectUpdate is of type Update
type ProjectUpdate struct {
	Version int                   `json:"version"`
	Actions []ProjectUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProjectUpdate) UnmarshalJSON(data []byte) error {
	type Alias ProjectUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorProjectUpdateAction(obj.Actions[i])
	}

	return nil
}
