package catalog

import (
	"encoding/json"

	"github.com/labd/commercetools-go-sdk/commercetools"
)

type ProductChangeName struct {
	Name commercetools.LocalizedString
}

// MarshalJSON override to add the Action() value
func (pa ProductChangeName) MarshalJSON() ([]byte, error) {
	type Alias ProductChangeName

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "changeName",
		Alias:  (*Alias)(&pa),
	})
}

type ProductAddPrice struct {
	VariantID int
	Price     int
	Staged    bool
}

// MarshalJSON override to add the Action() value
func (pa ProductAddPrice) MarshalJSON() ([]byte, error) {
	type Alias ProductAddPrice

	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{
		Action: "addPrice",
		Alias:  (*Alias)(&pa),
	})
}
