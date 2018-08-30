package project

import "encoding/json"

// ChangeName will change the name of the project being updated.
type ChangeName struct {
	Name string `json:"name"`
}

// MarshalJSON override to add the action value.
func (ua ChangeName) MarshalJSON() ([]byte, error) {
	type Alias ChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeCurrencies will change the currencies of the project being updated.
type ChangeCurrencies struct {
	// A three-digit currency code as per ISO 4217
	Currencies []string `json:"currencies"`
}

// MarshalJSON override to add the action value.
func (ua ChangeCurrencies) MarshalJSON() ([]byte, error) {
	type Alias ChangeCurrencies

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeCurrencies",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeCountries will change the countries of the project being updated.
type ChangeCountries struct {
	// A two-digit country code as per ISO 3166-1 alpha-2
	Countries []string `json:"countries"`
}

// MarshalJSON override to add the action value.
func (ua ChangeCountries) MarshalJSON() ([]byte, error) {
	type Alias ChangeCountries

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeCountries",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeLanguages will change the languages of the project being updated.
type ChangeLanguages struct {
	// IETF language tag
	Languages []string `json:"languages"`
}

// MarshalJSON override to add the action value.
func (ua ChangeLanguages) MarshalJSON() ([]byte, error) {
	type Alias ChangeLanguages

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeLanguages",
		Alias:  (*Alias)(&ua),
	})
}

// ChangeMessagesEnabled will change the state of the enabled messages value for
// the project being updated.
type ChangeMessagesEnabled struct {
	// Setting to true enables creation of Messages for the project.
	MessagesEnabled bool `json:"messagesEnabled"`
}

// MarshalJSON override to add the action value.
func (ua ChangeMessagesEnabled) MarshalJSON() ([]byte, error) {
	type Alias ChangeMessagesEnabled

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeMessagesEnabled",
		Alias:  (*Alias)(&ua),
	})
}

// SetShippingRateInputType will set the shipping rate input type for the
// project being updated.
type SetShippingRateInputType struct {
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType"`
}

// MarshalJSON override to add the action value.
func (ua SetShippingRateInputType) MarshalJSON() ([]byte, error) {
	type Alias SetShippingRateInputType

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setShippingRateInputType",
		Alias:  (*Alias)(&ua),
	})
}
