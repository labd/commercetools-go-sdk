package taxcategories

import "encoding/json"

// ChangeName will change the name of the tax category being updated.
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

// SetKey will set the key of the tax category being updated.
type SetKey struct {
	// If key is absent or null, it is removed if it exists.
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetKey) MarshalJSON() ([]byte, error) {
	type Alias SetKey

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setKey",
		Alias:  (*Alias)(&ua),
	})
}

// SetDescription will set the description of the tax category being updated.
type SetDescription struct {
	Description string `json:"description,omitempty"`
}

// MarshalJSON override to add the action value.
func (ua SetDescription) MarshalJSON() ([]byte, error) {
	type Alias SetDescription

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "setDescription",
		Alias:  (*Alias)(&ua),
	})
}

// AddTaxRate will add a tax rate to the tax category being updated.
type AddTaxRate struct {
	TaxRate TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to add the action value.
func (ua AddTaxRate) MarshalJSON() ([]byte, error) {
	type Alias AddTaxRate

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addTaxRate",
		Alias:  (*Alias)(&ua),
	})
}

// ReplaceTaxRate will replace a tax rate on the tax category being updated.
type ReplaceTaxRate struct {
	TaxRateID string       `json:"taxRateId"`
	TaxRate   TaxRateDraft `json:"taxRate"`
}

// MarshalJSON override to add the action value.
func (ua ReplaceTaxRate) MarshalJSON() ([]byte, error) {
	type Alias ReplaceTaxRate

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "replaceTaxRate",
		Alias:  (*Alias)(&ua),
	})
}

// RemoveTaxRate will remove a tax rate from the tax category being updated.
type RemoveTaxRate struct {
	TaxRateID string `json:"taxRateId"`
}

// MarshalJSON override to add the action value.
func (ua RemoveTaxRate) MarshalJSON() ([]byte, error) {
	type Alias RemoveTaxRate

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "removeTaxRate",
		Alias:  (*Alias)(&ua),
	})
}
