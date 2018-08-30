package project

import (
	"encoding/json"
	"time"

	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/mitchellh/mapstructure"
)

// Project endpoints are used to retrieve certain information from a project.
type Project struct {
	// The current version of the project.
	Version int `json:"version"`

	// The unique key of the project.
	Key string `json:"key"`

	// The name of the project.
	Name string `json:"name"`

	// A two-digit country code as per ISO 3166-1 alpha-2
	Countries []string `json:"countries"`

	// A three-digit currency code as per ISO 4217
	Currencies []string `json:"currencies"`

	// IETF language tag
	Languages []string `json:"languages"`

	CreatedAt time.Time `json:"createdAt"`

	// The time is in the format Year-Month YYYY-MM.
	TrialUntil string `json:"trialUntil,omitempty"`

	Messages              MessagesConfiguration `json:"messages"`
	ShippingRateInputType ShippingRateInputType `json:"shippingRateInputType,omitempty"`
}

// UnmarshalJSON override to map the shipping rate input type to the
// corresponding struct.
func (p *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	if err := json.Unmarshal(data, (*Alias)(p)); err != nil {
		return err
	}
	if p.ShippingRateInputType != nil {
		p.ShippingRateInputType = shippingRateInputTypeMapping(p.ShippingRateInputType)
	}
	return nil
}

// MessagesConfiguration is used to configure the Messages feature for the
// project.
type MessagesConfiguration struct {
	// When true the creation of Messages is enabled.
	Enabled bool `json:"enabled"`
}

// CartValue is used when the shipping rate maps to the sum of the line item
// prices. The value of the cart is used to select a tier. If chosen, it is not
// possible to set a value for the shippingRateInput on the cart. Tiers contain
// the centAmount (a value of 100 in the currency USD corresponds to $ 1.00),
// and start at 1.
type CartValue struct{}

// MarshalJSON override to add the type value.
func (t CartValue) MarshalJSON() ([]byte, error) {
	type Alias CartValue

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "CartValue",
		Alias: (*Alias)(&t),
	})
}

// CartClassification is used when the shipping rate maps to an abstract cart
// categorization expressed through a string, e.g. green, yellow, red or light,
// medium, heavy. Only a key defined inside the values array can be used to
// create a tier, or to set a value for the shippingRateInput on the cart. The
// keys are checked for uniqueness and the request is rejected if keys are not
// unique.
type CartClassification struct {
	Values []commercetools.LocalizedEnumValue `json:"values"`
}

// MarshalJSON override to add the type value.
func (t CartClassification) MarshalJSON() ([]byte, error) {
	type Alias CartClassification

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "CartClassification",
		Alias: (*Alias)(&t),
	})
}

// CartScore is used when the shipping rate maps to an abstract cart
// categorization expressed through an integer, e.g. shipping score or weight
// ranges. The range starts at 0. The default price covers the 0, tiers start at
// 1.
type CartScore struct{}

// MarshalJSON override to add the type value.
func (t CartScore) MarshalJSON() ([]byte, error) {
	type Alias CartScore

	return json.Marshal(struct {
		Type string `json:"type"`
		*Alias
	}{
		Type:  "CartScore",
		Alias: (*Alias)(&t),
	})
}

// ShippingRateInputType can be dynamically selected in three ways. The
// CartValue type uses the sum of all line item prices, whereas
// CartClassification and CartScore use the shipppingRateInput field on the cart
// to select a tier.
type ShippingRateInputType interface{}

func shippingRateInputTypeMapping(input ShippingRateInputType) ShippingRateInputType {
	FieldType := input.(map[string]interface{})["type"]
	switch FieldType {
	case "CartValue":
		newType := CartValue{}
		mapstructure.Decode(input, &newType)
		return newType
	case "CartClassification":
		newType := CartClassification{}
		mapstructure.Decode(input, &newType)
		return newType
	case "CartScore":
		newType := CartScore{}
		mapstructure.Decode(input, &newType)
		return newType
	}
	return nil
}
