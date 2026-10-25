package integrationlayer

// Generated file, please do not change!!!

import (
	"encoding/json"
)

type CheckoutMode string

const (
	CheckoutModeCOMPLETE     CheckoutMode = "COMPLETE"
	CheckoutModePAYMENT_ONLY CheckoutMode = "PAYMENT_ONLY"
)

type CheckoutRule struct {
	// ISO 3166-1 alpha-2 codes; absent matches any.
	Countries []string `json:"countries"`
	// Store keys; absent matches any.
	Stores         []string     `json:"stores"`
	ApplicationKey string       `json:"applicationKey"`
	Mode           CheckoutMode `json:"mode"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CheckoutRule) MarshalJSON() ([]byte, error) {
	type Alias CheckoutRule
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["countries"] == nil {
		delete(raw, "countries")
	}

	if raw["stores"] == nil {
		delete(raw, "stores")
	}

	return json.Marshal(raw)

}

type CheckoutConfig struct {
	// In order; the first match wins, and a catch-all rule must be last.
	Rules []CheckoutRule `json:"rules"`
}

type CheckoutConfigReplaceResult struct {
	Checkout CheckoutConfig `json:"checkout"`
	Version  int            `json:"version"`
}
