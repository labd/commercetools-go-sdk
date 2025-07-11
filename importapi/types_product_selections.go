package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Defines how variants are selected for the product.
*
 */
type VariantSelectionType string

const (
	VariantSelectionTypeIncludeOnly      VariantSelectionType = "includeOnly"
	VariantSelectionTypeIncludeAllExcept VariantSelectionType = "includeAllExcept"
)

/**
*	Variant selection specifying how variants are included or excluded.
*
 */
type VariantSelection struct {
	// Type of variant selection.
	Type VariantSelectionType `json:"type"`
	// List of SKUs to include or exclude.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSelection) MarshalJSON() ([]byte, error) {
	type Alias VariantSelection
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

	if raw["skus"] == nil {
		delete(raw, "skus")
	}

	return json.Marshal(raw)

}

/**
*	Defines which Variants of the Product will be excluded from the Product Selection. If not supplied all Variants are deemed to be excluded.
*
 */
type VariantExclusion struct {
	// List of SKUs to be excluded.
	Skus []string `json:"skus"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantExclusion) MarshalJSON() ([]byte, error) {
	type Alias VariantExclusion
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

	if raw["skus"] == nil {
		delete(raw, "skus")
	}

	return json.Marshal(raw)

}

/**
*	An assignment of a product and either variantSelection or variantExclusion (not both).
*
 */
type ProductSelectionAssignment struct {
	// Reference to the Product by key.
	Product ProductKeyReference `json:"product"`
	// Variant selection specifying included SKUs.
	VariantSelection *VariantSelection `json:"variantSelection,omitempty"`
	// Variant exclusion specifying excluded SKUs.
	VariantExclusion *VariantExclusion `json:"variantExclusion,omitempty"`
}

/**
*	Product Selections can have the following modes:
*	- `Individual`: Products must be explicitly assigned.
*	- `IndividualExclusion`: Products are included unless explicitly excluded.
*
 */
type ProductSelectionMode string

const (
	ProductSelectionModeIndividual          ProductSelectionMode = "Individual"
	ProductSelectionModeIndividualExclusion ProductSelectionMode = "IndividualExclusion"
)

/**
*	Represents the data used to import a ProductSelection. Once imported, this data is persisted as a [ProductSelection](ctp:api:type:ProductSelection) in the Project.
*
 */
type ProductSelectionImport struct {
	// User-defined unique identifier. If an [ProductSelection](ctp:api:type:ProductSelection) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `ProductSelection.name`.
	Name LocalizedString `json:"name"`
	// Maps to `ProductSelection.mode`.
	Mode *ProductSelectionMode `json:"mode,omitempty"`
	// Maps to `ProductSelection.custom`.
	Custom *Custom `json:"custom,omitempty"`
	// List of product assignments.
	Assignments []ProductSelectionAssignment `json:"assignments"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductSelectionImport) MarshalJSON() ([]byte, error) {
	type Alias ProductSelectionImport
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

	if raw["assignments"] == nil {
		delete(raw, "assignments")
	}

	return json.Marshal(raw)

}
