package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	A single ProductTailoring representation contains the _current_ and the _staged_ representation of its product information tailored per Store.
*
 */
type ProductTailoring struct {
	// Unique identifier of the ProductTailoring.
	ID string `json:"id"`
	// Current version of the ProductTailoring.
	Version int `json:"version"`
	// Date and time (UTC) the ProductTailoring was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the ProductTailoring was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the ProductTailoring.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the ProductTailoring.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// The Store to which the ProductTailoring belongs.
	Store StoreKeyReference `json:"store"`
	// Reference to the Product the ProductTailoring belongs to.
	Product ProductReference `json:"product"`
	// If `true`, the tailored information contained in the `current` [ProductTailoringData](ctp:api:type:ProductTailoringData) is provided when [retrieving the ProductProjection in Store](/../api/projects/product-tailoring#retrieve-product-projection-with-tailored-information).
	// For information not part of the ProductTailoringData, the original information contained in the [ProductData](ctp:api:type:ProductData) is provided.
	// If `false`, only the original information contained in the ProductData is provided.
	Published bool `json:"published"`
	// Current (published) data of the ProductTailoring.
	Current ProductTailoringData `json:"current"`
	// Staged (unpublished) data of the ProductTailoring.
	Staged ProductTailoringData `json:"staged"`
	// `true` if the `staged` data is different from the `current` data.
	HasStagedChanges bool `json:"hasStagedChanges"`
	// Warnings about processing of a request.
	// Appears in response to requests with response status code `202 Accepted`.
	Warnings []WarningObject `json:"warnings"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *ProductTailoring) UnmarshalJSON(data []byte) error {
	type Alias ProductTailoring
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Warnings {
		var err error
		obj.Warnings[i], err = mapDiscriminatorWarningObject(obj.Warnings[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoring) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoring
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

	if raw["warnings"] == nil {
		delete(raw, "warnings")
	}

	return json.Marshal(raw)

}

/**
*	The same rules for `name` and `value` apply as for [Attribute](ctp:api:type:Attribute) in Product Variants.
*
 */
type ProductTailoringAttribute struct {
	// Name of the Attribute.
	Name string `json:"name"`
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	//
	// Tailoring of [Nested Type](ctp:api:type:AttributeNestedType) Attributes is not supported.
	Value interface{} `json:"value"`
}

/**
*	Contains all the tailored information of a Product.
*
 */
type ProductTailoringData struct {
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product that is used by search engines.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product that is used by search engines.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product that are used by search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// Tailored Variants of the Product.
	Variants []ProductVariantTailoring `json:"variants"`
	// Attributes of the tailored Product.
	// If available, these Attributes are selectively merged into the `attributes` of the corresponding [Product](ctp:api:type:Product). If the Product contains an Attribute with the same `name`, then its `value` is overwritten. Otherwise, the Attribute and its `value` are added to the Product.
	Attributes []ProductTailoringAttribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringData) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringData
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

/**
*	Contains the information to be tailored for a Product.
*
 */
type ProductTailoringDraft struct {
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// The Store to which the ProductTailoring belongs.
	Store StoreResourceIdentifier `json:"store"`
	// ResourceIdentifier of the Product the ProductTailoring belongs to.
	Product ProductResourceIdentifier `json:"product"`
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product that is used by search engines.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product that is used by search engines.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product that are used by search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// Set to `true` to [publish](/../api/projects/product-tailoring#stage-and-publish-tailored-product-information) the ProductTailoring immediately.
	// Otherwise, the tailored product information is just staged.
	Publish *bool `json:"publish,omitempty"`
	// Tailored Variants of the Product.
	Variants []ProductVariantTailoringDraft `json:"variants"`
	// Attributes of the tailored Product.
	// If provided, these Attributes are selectively merged into the `attributes` of the corresponding [Product](ctp:api:type:Product). If the Product contains an Attribute with the same `name`, then its `value` is overwritten. Otherwise, the Attribute and its `value` are added to the Product.
	Attributes []ProductTailoringAttribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringDraft
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

/**
*	Contains all the tailored information of a Product for a specific Store.
*
 */
type ProductTailoringInStoreDraft struct {
	// User-defined unique identifier of the ProductTailoring.
	Key *string `json:"key,omitempty"`
	// ResourceIdentifier of the Product the ProductTailoring belongs to.
	Product ProductResourceIdentifier `json:"product"`
	// Tailored name of the Product.
	Name *LocalizedString `json:"name,omitempty"`
	// Tailored description of the Product.
	Description *LocalizedString `json:"description,omitempty"`
	// Tailored title of the Product that is used by search engines.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Tailored description of the Product that is used in search engine.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Tailored keywords related to the Product that are used by search engines.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// User-defined identifier used in a deep-link URL for the ProductTailoring.
	// Matches the pattern `[a-zA-Z0-9_\\-]{2,256}`.
	Slug *LocalizedString `json:"slug,omitempty"`
	// If `true`, the ProductTailoring is published immediately.
	Publish *bool `json:"publish,omitempty"`
	// Tailored Variants of the Product.
	Variants []ProductVariantTailoringDraft `json:"variants"`
	// Attributes of the tailored Product.
	// If provided, these Attributes are selectively merged into the `attributes` of the corresponding [Product](ctp:api:type:Product). If the Product contains an Attribute with the same `name`, then its `value` is overwritten. Otherwise, the Attribute and its `value` are added to the Product.
	Attributes []ProductTailoringAttribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringInStoreDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringInStoreDraft
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

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// Present only when the `withTotal` query parameter is set to `true`.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [ProductTailoring](ctp:api:type:ProductTailoring) list matching the query.
	Results []ProductTailoring `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringReference struct {
	// Unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring).
	ID string `json:"id"`
	// Contains the representation of the expanded ProductTailoring. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for ProductTailoring.
	Obj *ProductTailoring `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringReference) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-tailoring", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [ProductTailoring](ctp:api:type:ProductTailoring).
*
 */
type ProductTailoringResourceIdentifier struct {
	// Unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [ProductTailoring](ctp:api:type:ProductTailoring). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "product-tailoring", Alias: (*Alias)(&obj)})
}

type ProductTailoringUpdateAction interface{}

func mapDiscriminatorProductTailoringUpdateAction(input interface{}) (ProductTailoringUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addAsset":
		obj := ProductTailoringAddAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addExternalImage":
		obj := ProductTailoringAddExternalImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addVariant":
		obj := ProductTailoringAddVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetName":
		obj := ProductTailoringChangeAssetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetOrder":
		obj := ProductTailoringChangeAssetOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "moveImageToPosition":
		obj := ProductTailoringMoveImageToPositionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "publish":
		obj := ProductTailoringPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAsset":
		obj := ProductTailoringRemoveAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeImage":
		obj := ProductTailoringRemoveImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeVariant":
		obj := ProductTailoringRemoveVariantAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomField":
		obj := ProductTailoringSetAssetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomType":
		obj := ProductTailoringSetAssetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetDescription":
		obj := ProductTailoringSetAssetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetKey":
		obj := ProductTailoringSetAssetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetSources":
		obj := ProductTailoringSetAssetSourcesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetTags":
		obj := ProductTailoringSetAssetTagsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttribute":
		obj := ProductTailoringSetAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttributeInAllVariants":
		obj := ProductTailoringSetAttributeInAllVariantsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := ProductTailoringSetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImages":
		obj := ProductTailoringSetExternalImagesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImageLabel":
		obj := ProductTailoringSetImageLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaAttributes":
		obj := ProductTailoringSetMetaAttributesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaDescription":
		obj := ProductTailoringSetMetaDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaKeywords":
		obj := ProductTailoringSetMetaKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaTitle":
		obj := ProductTailoringSetMetaTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setName":
		obj := ProductTailoringSetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setProductAttribute":
		obj := ProductTailoringSetProductAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSlug":
		obj := ProductTailoringSetSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unpublish":
		obj := ProductTailoringUnpublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	The tailoring of a [ProductVariant](ctp:api:type:ProductVariant).
*
 */
type ProductVariantTailoring struct {
	// The `id` of the tailored [ProductVariant](ctp:api:type:ProductVariant).
	ID int `json:"id"`
	// Images of the tailored Product Variant.
	// If present, these images will override the images of the corresponding [ProductVariant](ctp:api:type:ProductVariant) in total.
	Images []Image `json:"images"`
	// Media assets of the tailored Product Variant.
	// If present, these assets will override the assets of the corresponding [ProductVariant](ctp:api:type:ProductVariant) in total.
	Assets []Asset `json:"assets"`
	// Attributes of the tailored Product Variant.
	// If present, these Attributes are selectively merged into the `attributes` of the corresponding [ProductVariant](ctp:api:type:ProductVariant):
	//
	// - If the ProductVariant contains an Attribute with the same `name`, its `value` is overwritten,
	// - otherwise the Attribute and its value are added to the ProductVariant.
	Attributes []ProductTailoringAttribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantTailoring) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantTailoring
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
*	Either `id` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductVariantTailoringDraft struct {
	// The `id` of the [ProductVariant](ctp:api:type:ProductVariant) to be tailored.
	ID *int `json:"id,omitempty"`
	// The `sku` of the [ProductVariant](ctp:api:type:ProductVariant) to be tailored.
	Sku *string `json:"sku,omitempty"`
	// Images of the tailored Product Variant.
	Images []Image `json:"images"`
	// Media assets of the tailored Product Variant.
	Assets []Asset `json:"assets"`
	// Attributes of the tailored Product Variant according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	// If provided, these Attributes are selectively merged into the `attributes` of the corresponding [ProductVariant](ctp:api:type:ProductVariant):
	//
	// - If the ProductVariant contains an Attribute with the same `name`, its `value` is overwritten,
	// - otherwise the Attribute and its value are added to the ProductVariant.
	Attributes []ProductTailoringAttribute `json:"attributes"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductVariantTailoringDraft) MarshalJSON() ([]byte, error) {
	type Alias ProductVariantTailoringDraft
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
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringAddAssetAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged `assets` are updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// Value to append.
	Asset AssetDraft `json:"asset"`
	// Position in `assets` where the Asset should be put. When specified, the value must be between `0` and the total number of Assets minus `1`.
	Position *int `json:"position,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists. Produces the [ProductTailoringImageAdded](/projects/messages/product-catalog-messages#product-tailoring-image-added) Message.
*
 */
type ProductTailoringAddExternalImageAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Value to add to `images`.
	Image Image `json:"image"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

/**
*	Either `id` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	Produces the [ProductVariantTailoringAdded](ctp:api:type:ProductVariantTailoringAddedMessage) Message.
*
 */
type ProductTailoringAddVariantAction struct {
	// The `id` of the tailored ProductVariant to update.
	ID *int `json:"id,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Images for the Product Variant Tailoring.
	Images []Image `json:"images"`
	// Media assets for the Product Variant Tailoring.
	Assets []AssetDraft `json:"assets"`
	// Attributes for the Product Variant Tailoring.
	Attributes []ProductTailoringAttribute `json:"attributes"`
	// If `true` the new Product Variant Tailoring is only staged. If `false` the new Product Variant Tailoring is both current and staged.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringAddVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringAddVariantAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addVariant", Alias: (*Alias)(&obj)})
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
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringChangeAssetNameAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringChangeAssetOrderAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged `assets` is updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// All existing Asset `id`s of the ProductVariantTailoring in the desired new order.
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringMoveImageToPositionAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The URL of the image to update.
	ImageUrl string `json:"imageUrl"`
	// Position in `images` where the image should be moved. Must be between `0` and the total number of images minus `1`.
	Position int `json:"position"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

/**
*	Publishes the `staged` data of the ProductTailoring to `current`. Sets `hasStagedChanges` to `false`.
*	Generates the [ProductTailoringPublished](ctp:api:type:ProductTailoringPublishedMessage) Message.
*
 */
type ProductTailoringPublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringPublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The Asset to remove must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringRemoveAssetAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is removed. If `false`, both the current and staged Asset is removed.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to remove.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to remove.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringRemoveImageAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The URL of the image to remove.
	ImageUrl string `json:"imageUrl"`
	// If `true`, only the staged image is removed. If `false`, both the current and staged image is removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

/**
*	Either `id` or `sku` is required.
*	Produces the [ProductVariantTailoringDeleted](ctp:api:type:ProductVariantTailoringRemovedMessage) Message.
*
 */
type ProductTailoringRemoveVariantAction struct {
	// The `id` of the ProductVariant to remove from the Tailoring.
	ID *int `json:"id,omitempty"`
	// The `sku` of the ProductVariant to remove from the Tailoring.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Product Variant Tailoring is removed. If `false`, both the current and staged Product Variant Tailoring is removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringRemoveVariantAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringRemoveVariantAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeVariant", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringSetAssetCustomFieldAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Removing a field that does not exist returns an [InvalidOperation](ctp:api:type:InvalidOperationError) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringSetAssetCustomTypeAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Asset with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Asset.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Asset.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringSetAssetDescriptionAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringSetAssetKeyAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId string `json:"assetId"`
	// Value to set. If empty, any existing value will be removed.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The [Asset](ctp:api:type:Asset) to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringSetAssetSourcesAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false` both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set.
	Sources []AssetSource `json:"sources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*	The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type ProductTailoringSetAssetTagsAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// If `true`, only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAssetTagsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	if raw["tags"] == nil {
		delete(raw, "tags")
	}

	return json.Marshal(raw)

}

/**
*	Either `variantId` or `sku` is required.
*
 */
type ProductTailoringSetAttributeAction struct {
	// The `id` of the ProductVariant to update.
	// Required if `sku` is absent.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the ProductVariant to update.
	// Required if `variantId` is absent.
	Sku *string `json:"sku,omitempty"`
	// The name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attribute. If empty, any existing value will be removed.
	//
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	//
	// Tailoring of [Nested Type](ctp:api:type:AttributeNestedType) Attributes is **not supported**.
	Value interface{} `json:"value,omitempty"`
	// If `true`, only the staged Attribute is set. If `false`, both current and staged Attribute is set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

/**
*	Adds, removes, or changes a tailored Attribute in all Product Variants of a Product at the same time.
*	This action is useful for setting tailored values for Attributes with the [Constraint](ctp:api:type:AttributeConstraintEnum) `SameForAll`.
 */
type ProductTailoringSetAttributeInAllVariantsAction struct {
	// The name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attributes. If empty, any existing value will be removed.
	//
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) objects,
	//   or the complete objects as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	//
	// Tailoring of [Nested Type](ctp:api:type:AttributeNestedType) Attributes is **not supported**.
	Value interface{} `json:"value,omitempty"`
	// If `true`, only the staged Attributes are set. If `false`, both the current and staged Attributes are set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetAttributeInAllVariantsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetAttributeInAllVariantsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributeInAllVariants", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringDescriptionSet](ctp:api:type:ProductTailoringDescriptionSetMessage) Message.
*
 */
type ProductTailoringSetDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	Description *LocalizedString `json:"description,omitempty"`
	// If `true`, only the staged `description` is updated. If `false`, both the current and staged `description` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists. Produces the [ProductTailoringImagesSet](/projects/messages/product-catalog-messages#product-tailoring-images-set) Message.
*
 */
type ProductTailoringSetExternalImagesAction struct {
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// Images of the tailored ProductVariant.
	//
	// Don't provide this field if you want to remove all images from the tailored Product Variant.
	// Set to `[]` (empty) if you want to hide all images of the original ProductVariant on the tailored ProductVariant.
	Images []Image `json:"images"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetExternalImagesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetExternalImagesAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImages", Alias: (*Alias)(&obj)})
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

	return json.Marshal(raw)

}

/**
*	Either `variantId` or `sku` is required to reference a [ProductVariant](ctp:api:type:ProductVariant) that exists.
*
 */
type ProductTailoringSetImageLabelAction struct {
	// The `sku` of the tailored ProductVariant to update.
	Sku *string `json:"sku,omitempty"`
	// The `id` of the tailored ProductVariant to update.
	VariantId *int `json:"variantId,omitempty"`
	// The URL of the image to set the label.
	ImageUrl string `json:"imageUrl"`
	// Value to set. If empty, any existing value will be removed.
	Label *string `json:"label,omitempty"`
	// If `true`, only the staged image is updated. If `false`, both the current and staged image is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

/**
*	Updates all meta attributes at the same time.
 */
type ProductTailoringSetMetaAttributesAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Value to set. If empty, any existing value will be removed.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// If `true`, only the staged attributes are updated. If `false`, both the current and staged attributes are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaAttributesAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaAttributesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaAttributes", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaDescriptionAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// If `true`, only the staged `metaDescription` is updated. If `false`, both the current and staged `metaDescription` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaKeywordsAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// If `true`, only the staged `metaKeywords` is updated. If `false`, both the current and staged `metaKeywords` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetMetaTitleAction struct {
	// Value to set. If empty, any existing value will be removed.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// If `true`, only the staged `metaTitle` is updated. If `false`, both the current and staged `metaTitle` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringNameSet](ctp:api:type:ProductTailoringNameSetMessage) Message.
*
 */
type ProductTailoringSetNameAction struct {
	// Value to set. If empty, any existing value will be removed.
	Name *LocalizedString `json:"name,omitempty"`
	// If `true`, only the staged `name` is updated. If `false`, both the current and staged `name` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetNameAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setName", Alias: (*Alias)(&obj)})
}

type ProductTailoringSetProductAttributeAction struct {
	// Name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attribute. If empty, then any existing value will be removed.
	//
	// [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType),
	//   use either the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) object or the complete object as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	//
	// Tailoring of [Nested Type](ctp:api:type:AttributeNestedType) Attributes is **not supported**.
	Value interface{} `json:"value,omitempty"`
	// If `true`, then only the staged Attribute is set. If `false`, then both the current and staged Attributes are set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetProductAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetProductAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setProductAttribute", Alias: (*Alias)(&obj)})
}

/**
*	Generates the [ProductTailoringSlugSet](ctp:api:type:ProductTailoringSlugSetMessage) Message.
*
 */
type ProductTailoringSetSlugAction struct {
	// Value to set. If empty, any existing value will be removed.
	Slug *LocalizedString `json:"slug,omitempty"`
	// If `true`, only the staged `slug` is updated. If `false`, both the current and staged `slug` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringSetSlugAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringSetSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSlug", Alias: (*Alias)(&obj)})
}

/**
*	Unpublishes the `current` data of the ProductTailoring. Sets the `published` field to `false`.
*	Generates the [ProductTailoringUnpublished](ctp:api:type:ProductTailoringUnpublishedMessage) Message.
*
 */
type ProductTailoringUnpublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj ProductTailoringUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias ProductTailoringUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}
