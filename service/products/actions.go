package products

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

// ChangeName will change the name of the product being updated.
type ChangeName struct {
	Name commercetools.LocalizedString
}

// MarshalJSON override to add the Action() value
func (pa ChangeName) MarshalJSON() ([]byte, error) {
	type Alias ChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&pa),
	})
}

// AddPrice will add a price to the product being updated.
type AddPrice struct {
	VariantID int
	Price     int
	Staged    bool
}

// MarshalJSON override to add the Action() value
func (pa AddPrice) MarshalJSON() ([]byte, error) {
	type Alias AddPrice

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addPrice",
		Alias:  (*Alias)(&pa),
	})
}
