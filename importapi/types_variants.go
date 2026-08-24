package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Represents the data used to import a Variant.
*	Once imported, this data is persisted as a [Variant](ctp:api:type:Variant) in the Project.
*
*	This import resource is only available for Projects with `productCatalogModel` set to `Modular`.
*
 */
type VariantImport struct {
	// User-defined unique identifier. If a Variant with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// Maps to `VariantData.sku`.
	Sku string `json:"sku"`
	// Maps to `Variant.product`.
	Product ProductKeyReference `json:"product"`
	// Indicates whether the Variant is published. If `true`, the Variant is immediately available in the current published state.
	// Maps to `VariantDraft.publish`.
	Publish *bool `json:"publish,omitempty"`
	// Maps to `VariantData.images`.
	Images []Image `json:"images"`
	// Maps to `VariantData.attributes`.
	Attributes []Attribute `json:"attributes"`
	// Maps to `VariantData.assets`.
	Assets []Asset `json:"assets"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantImport) UnmarshalJSON(data []byte) error {
	type Alias VariantImport
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Attributes {
		var err error
		obj.Attributes[i], err = mapDiscriminatorAttribute(obj.Attributes[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantImport) MarshalJSON() ([]byte, error) {
	type Alias VariantImport
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

	if raw["images"] == nil {
		delete(raw, "images")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	return json.Marshal(raw)

}
