package taxcategories

import "time"

// TaxCategory defines how products are to be taxed in different countries.
type TaxCategory struct {
	// The unique ID of the category.
	ID string `json:"id"`

	// User-specific unique identifier for the category.
	Key string `json:"key,omitempty"`

	// The current version of the category.
	Version int `json:"version"`

	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	Name           string    `json:"name"`
	Description    string    `json:"description,omitempty"`

	// The tax rates have unique IDs in the rates list.
	Rates []TaxRate `json:"rates"`
}

// TaxCategoryDraft is sent with Create TaxCategory requests.
type TaxCategoryDraft struct {
	Name        string         `json:"name"`
	Key         string         `json:"key,omitempty"`
	Description string         `json:"description,omitempty"`
	Rates       []TaxRateDraft `json:"rates"`
}

type TaxRate struct {
	// The id is always set if the tax rate is part of a TaxCategory. The
	// external tax rates in a Cart do not contain an id.
	ID string `json:"id,omitempty"`

	Name string `json:"name"`

	// Number Percentage in the range of [0..1]. The sum of the amounts of all
	// subRates, if there are any.
	Amount float64 `json:"amount"`

	IncludedInPrice bool `json:"includedInPrice"`

	// A two-digit country code as per ISO 3166-1 alpha-2.
	Country string `json:"country"`

	// The state in the country
	State string `json:"state,omitempty"`

	// For countries (e.g. the US) where the total tax is a combination of
	// multiple taxes (e.g. state and local taxes).
	SubRates []SubRate `json:"subRates"`
}

// SubRate is used to calculate the taxPortions field in a cart or order. It is
// useful if the total tax of a country is a combination of multiple taxes (e.g.
// state and local taxes).
type SubRate struct {
	Name string `json:"name"`

	// Number Percentage in the range of [0..1].
	Amount float64 `json:"amount"`
}

type TaxRateDraft struct {
	Name            string  `json:"name"`
	Amount          float64 `json:"amount,omitempty"`
	IncludedInPrice bool    `json:"includedInPrice"`

	// A two-digit country code as per ISO 3166-1 alpha-2.
	Country string `json:"country"`

	// The state in the country
	State string `json:"state,omitempty"`

	// For countries (e.g. the US) where the total tax is a combination of
	// multiple taxes (e.g. state and local taxes).
	SubRates []SubRate `json:"subrates,omitempty"`
}
