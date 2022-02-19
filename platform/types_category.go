// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Category struct {
	// The unique ID of the category.
	ID string `json:"id"`
	// The current version of the category.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy      `json:"createdBy,omitempty"`
	Name      LocalizedString `json:"name"`
	// human-readable identifiers usually used as deep-link URL to the related category.
	// Each slug is unique across a project, but a category can have the same slug for different languages.
	Slug        LocalizedString  `json:"slug"`
	Description *LocalizedString `json:"description,omitempty"`
	// Contains the parent path towards the root category.
	Ancestors []CategoryReference `json:"ancestors"`
	// A category that is the parent of this category in the category tree.
	Parent *CategoryReference `json:"parent,omitempty"`
	// An attribute as base for a custom category order in one level.
	OrderHint       string           `json:"orderHint"`
	ExternalId      *string          `json:"externalId,omitempty"`
	MetaTitle       *LocalizedString `json:"metaTitle,omitempty"`
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	MetaKeywords    *LocalizedString `json:"metaKeywords,omitempty"`
	Custom          *CustomFields    `json:"custom,omitempty"`
	// Can be used to store images, icons or movies related to this category.
	Assets []Asset `json:"assets"`
	// User-specific unique identifier for the category.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj Category) MarshalJSON() ([]byte, error) {
	type Alias Category
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["assets"] == nil {
		delete(target, "assets")
	}

	return json.Marshal(target)
}

type CategoryDraft struct {
	Name LocalizedString `json:"name"`
	// human-readable identifier usually used as deep-link URL to the related category.
	// Allowed are alphabetic, numeric, underscore (`_`) and hyphen (`-`) characters.
	// Maximum size is 256.
	// **Must be unique across a project!** The same category can have the same slug for different languages.
	Slug        LocalizedString  `json:"slug"`
	Description *LocalizedString `json:"description,omitempty"`
	// A category that is the parent of this category in the category tree.
	// The parent can be set by its ID or by its key.
	Parent *CategoryResourceIdentifier `json:"parent,omitempty"`
	// An attribute as base for a custom category order in one level.
	// A random value will be assigned by API if not set.
	OrderHint       *string          `json:"orderHint,omitempty"`
	ExternalId      *string          `json:"externalId,omitempty"`
	MetaTitle       *LocalizedString `json:"metaTitle,omitempty"`
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	MetaKeywords    *LocalizedString `json:"metaKeywords,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	Assets []AssetDraft       `json:"assets"`
	// User-defined unique identifier for the category.
	// Keys can only contain alphanumeric characters (`a-Z, 0-9`), underscores and hyphens (`-, _`) and be between 2 and 256 characters.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryDraft) MarshalJSON() ([]byte, error) {
	type Alias CategoryDraft
	data, err := json.Marshal(struct {
		*Alias
	}{Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["assets"] == nil {
		delete(target, "assets")
	}

	return json.Marshal(target)
}

type CategoryPagedQueryResponse struct {
	Limit   int        `json:"limit"`
	Count   int        `json:"count"`
	Total   *int       `json:"total,omitempty"`
	Offset  int        `json:"offset"`
	Results []Category `json:"results"`
}

type CategoryReference struct {
	// Unique ID of the referenced resource.
	ID  string    `json:"id"`
	Obj *Category `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

type CategoryResourceIdentifier struct {
	// Unique ID of the referenced resource. Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// Unique key of the referenced resource. Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias CategoryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

type CategoryUpdate struct {
	Version int                    `json:"version"`
	Actions []CategoryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *CategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias CategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}

	return nil
}

type CategoryUpdateAction interface{}

func mapDiscriminatorCategoryUpdateAction(input interface{}) (CategoryUpdateAction, error) {

	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("Error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("Invalid data")
	}

	switch discriminator {
	case "addAsset":
		obj := CategoryAddAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetName":
		obj := CategoryChangeAssetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetOrder":
		obj := CategoryChangeAssetOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeName":
		obj := CategoryChangeNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeOrderHint":
		obj := CategoryChangeOrderHintAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeParent":
		obj := CategoryChangeParentAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeSlug":
		obj := CategoryChangeSlugAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAsset":
		obj := CategoryRemoveAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomField":
		obj := CategorySetAssetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomType":
		obj := CategorySetAssetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetDescription":
		obj := CategorySetAssetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetKey":
		obj := CategorySetAssetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetSources":
		obj := CategorySetAssetSourcesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetTags":
		obj := CategorySetAssetTagsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := CategorySetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := CategorySetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setDescription":
		obj := CategorySetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExternalId":
		obj := CategorySetExternalIdAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := CategorySetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaDescription":
		obj := CategorySetMetaDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaKeywords":
		obj := CategorySetMetaKeywordsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setMetaTitle":
		obj := CategorySetMetaTitleAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type CategoryAddAssetAction struct {
	Asset AssetDraft `json:"asset"`
	// When specified, the value might be `0` and should be lower than the total of the assets list.
	Position *int `json:"position,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

type CategoryChangeAssetNameAction struct {
	AssetId  *string         `json:"assetId,omitempty"`
	AssetKey *string         `json:"assetKey,omitempty"`
	Name     LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

type CategoryChangeAssetOrderAction struct {
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

type CategoryChangeNameAction struct {
	Name LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

type CategoryChangeOrderHintAction struct {
	OrderHint string `json:"orderHint"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderHint", Alias: (*Alias)(&obj)})
}

type CategoryChangeParentAction struct {
	Parent CategoryResourceIdentifier `json:"parent"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeParentAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeParentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParent", Alias: (*Alias)(&obj)})
}

type CategoryChangeSlugAction struct {
	// Allowed are alphabetic, numeric, underscore (&#95;) and hyphen (&#45;) characters.
	// Maximum size is {{ site.data.api-limits.slugLength }}.
	Slug LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type CategoryRemoveAssetAction struct {
	AssetId  *string `json:"assetId,omitempty"`
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategoryRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomFieldAction struct {
	AssetId  *string `json:"assetId,omitempty"`
	AssetKey *string `json:"assetKey,omitempty"`
	Name     string  `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomTypeAction struct {
	AssetId  *string `json:"assetId,omitempty"`
	AssetKey *string `json:"assetKey,omitempty"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

type CategorySetAssetDescriptionAction struct {
	AssetId     *string          `json:"assetId,omitempty"`
	AssetKey    *string          `json:"assetKey,omitempty"`
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

type CategorySetAssetKeyAction struct {
	AssetId string `json:"assetId"`
	// User-defined identifier for the asset.
	// If left blank or set to `null`, the asset key is unset/removed.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

type CategorySetAssetSourcesAction struct {
	AssetId  *string       `json:"assetId,omitempty"`
	AssetKey *string       `json:"assetKey,omitempty"`
	Sources  []AssetSource `json:"sources"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

type CategorySetAssetTagsAction struct {
	AssetId  *string  `json:"assetId,omitempty"`
	AssetKey *string  `json:"assetKey,omitempty"`
	Tags     []string `json:"tags"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetTagsAction
	data, err := json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
	if err != nil {
		return nil, err
	}

	target := make(map[string]interface{})
	if err := json.Unmarshal(data, &target); err != nil {
		return nil, err
	}

	if target["tags"] == nil {
		delete(target, "tags")
	}

	return json.Marshal(target)
}

type CategorySetCustomFieldAction struct {
	Name string `json:"name"`
	// The value of a Custom Field.
	// The data type of the value depends on the specific [FieldType](/projects/types#fieldtype) given in the `type` field of the [FieldDefinition](/ctp:api:type:FieldDefinition) for a Custom Field.
	// It can be any of the following:
	//
	//  Field type                                                 | Data type                                                                                                                                                                 |
	//  ---------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
	//  [CustomFieldBooleanType](ctp:api:type:CustomFieldBooleanType)                 | Boolean (`true` or `false`)                                                                                                                                                     |
	//  [CustomFieldStringType](ctp:api:type:CustomFieldStringType)                   | String                                                                                                                                                                |
	//  [CustomFieldLocalizedStringType](ctp:api:type:CustomFieldLocalizedStringType) | [LocalizedString](ctp:api:type:LocalizedString)                                                                                                                             |
	//  [CustomFieldEnumType](ctp:api:type:CustomFieldEnumType)                       | String. Must be a `key` of one of the [EnumValues](ctp:api:type:CustomFieldEnumValue) defined in the [EnumType](ctp:api:type:CustomFiedEnumType)                                     |
	//  [CustomFieldLocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType)     | String. Must be a `key` of one of the [LocalizedEnumValues](ctp:api:type:CustomFieldLocalizedEnumValue) defined in the [LocalizedEnumType](ctp:api:type:CustomFieldLocalizedEnumType) |
	//  [CustomFieldNumberType](ctp:api:type:CustomFieldNumberType)                   | Number                                                                                                                                                                |
	//  [CustomFieldMoneyType](ctp:api:type:CustomFieldMoneyType)                     | [CentPrecisionMoney](/../api/types#centprecisionmoney)                                                                                                                                         |
	//  [CustomFieldDateType](ctp:api:type:CustomFieldDateType)                       | [Date](ctp:api:type:Date)                                                                                                                                                   |
	//  [CustomFieldTimeType](ctp:api:type:CustomFieldTimeType)                       | [Time](ctp:api:type:Time)                                                                                                                                                   |
	//  [CustomFieldDateTimeType](ctp:api:type:CustomFieldDateTimeType)               | [DateTime](ctp:api:type:DateTime)                                                                                                                                           |
	//  [CustomFieldReferenceType](ctp:api:type:CustomFieldReferenceType)             | [Reference](/../api/types#reference)                                                                                                                                         |
	//  [CustomFieldSetType](ctp:api:type:CustomFieldSetType)                         | JSON array without duplicates consisting of [CustomFieldValues](ctp:api:type:CustomFieldValue) of a single [FieldType](ctp:api:type:FieldType). For example, a Custom Field of SetType of DateType takes a JSON array of mutually different Dates for its value. The order of items in the array is not fixed. For more examples, see the [example FieldContainer](ctp:api:type:FieldContainer).|
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// A valid JSON object, based on the FieldDefinitions of the Type. Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CategorySetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CategorySetExternalIdAction struct {
	// If not defined, the external ID is unset.
	ExternalId *string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type CategorySetKeyAction struct {
	// User-defined unique identifier for the category.
	// Keys can only contain alphanumeric characters (`a-Z, 0-9`), underscores and hyphens (`-, _`) and be between 2 and 256 characters.
	// If `key` is absent or `null`, this field will be removed if it exists.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CategorySetMetaDescriptionAction struct {
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type CategorySetMetaKeywordsAction struct {
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type CategorySetMetaTitleAction struct {
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj CategorySetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}
