package importapi

// Generated file, please do not change!!!

import (
	"encoding/json"
)

/**
*	Represents the data of a Product Variant to be tailored for a Store.
*
 */
type ProductVariantTailoringImport struct {
	// The `sku` of the [ProductVariant](ctp:api:type:ProductVariant) to be tailored.
	Sku string `json:"sku"`
	// Maps to `ProductVariantTailoring.images`.
	Images []Image `json:"images"`
	// Maps to `ProductVariantTailoring.assets`.
	Assets []Asset `json:"assets"`
	// Maps to `ProductVariantTailoring.attributes`.
	Attributes []Attribute `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductVariantTailoringImport) UnmarshalJSON(data []byte) error {
	type Alias ProductVariantTailoringImport
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
func (obj ProductVariantTailoringImport) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantTailoringImport
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

	if raw["assets"] == nil {
		delete(raw, "assets")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	return json.Marshal(raw)

}

/**
*	Represents the data used to import a Product Tailoring. Once imported, this data is persisted as a [ProductTailoring](ctp:api:type:ProductTailoring) in the Project.
*
 */
type ProductTailoringImport struct {
	// User-defined unique identifier. If a [ProductTailoring](ctp:api:type:ProductTailoring) with this `key` exists, it is updated with the imported data.
	Key string `json:"key"`
	// The Store to which the ProductTailoring belongs. If the referenced [Store](ctp:api:type:Store) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Store is created.
	Store StoreKeyReference `json:"store"`
	// The Product to which the ProductTailoring belongs. If the referenced [Product](ctp:api:type:Product) does not exist, the `state` of the [ImportOperation](ctp:import:type:ImportOperation) will be set to `unresolved` until the referenced Product is created.
	Product ProductKeyReference `json:"product"`
	// Maps to `ProductTailoring.name`.
	Name *LocalizedString `json:"name,omitempty"`
	// Maps to `ProductTailoring.description`.
	Description *LocalizedString `json:"description,omitempty"`
	// Maps to `ProductTailoring.metaTitle`.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Maps to `ProductTailoring.metaDescription`.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Maps to `ProductTailoring.metaKeywords`.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// Maps to `ProductTailoring.slug`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// Whether to publish the ProductTailoring immediately. Otherwise, the tailored product information is just staged.
	Publish *bool `json:"publish,omitempty"`
	// Tailored Variants of the Product.
	Variants []ProductVariantTailoringImport `json:"variants"`
	// Attributes of the tailored Product.
	Attributes []Attribute `json:"attributes"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductTailoringImport) UnmarshalJSON(data []byte) error {
	type Alias ProductTailoringImport
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
func (obj ProductTailoringImport) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringImport
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

	if raw["variants"] == nil {
		delete(raw, "variants")
	}

	if raw["attributes"] == nil {
		delete(raw, "attributes")
	}

	return json.Marshal(raw)

}
