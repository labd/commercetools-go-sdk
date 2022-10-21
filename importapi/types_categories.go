package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	The data representation for a Category to be imported that is persisted as a [Category](/../api/projects/categories#category) in the Project.
*
 */
type CategoryImport struct {
	// User-defined unique identifier.
	Key string `json:"key"`
	// Maps to `Category.name`.
	Name LocalizedString `json:"name"`
	// Maps to `Category.slug`.
	// Must match the pattern `[-a-zA-Z0-9_]{2,256}`.
	Slug LocalizedString `json:"slug"`
	// Maps to `Category.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `Category.parent`.
	// The Reference to the parent [Category](/../api/projects/categories#category) with which the Category is associated.
	// If referenced Category does not exist, the `state` of the [ImportOperation](/import-operation#importoperation) will be set to `unresolved` until the necessary Category is created.
	Parent *CategoryKeyReference `json:"parent,omitempty"`
	// Maps to `Category.orderHint`.
	OrderHint *string `json:"orderHint,omitempty"`
	// Maps to `Category.externalId`.
	ExternalId *string `json:"externalId,omitempty"`
	// Maps to `Category.metaTitle`.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Maps to `Category.metaDescription`.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Maps to `Category.metaKeywords`.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	Assets       []Asset          `json:"assets"`
	// The custom fields for this Category.
	Custom *Custom `json:"custom,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryImport) MarshalJSON() ([]byte, error) {
	type Alias CategoryImport
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

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}
