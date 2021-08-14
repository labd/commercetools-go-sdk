// Generated file, please do not change!!!
package platform

import (
	"encoding/json"
	"errors"
	"time"
)

type Category struct {
	// The unique ID of the category.
	Id string `json:"id"`
	// The current version of the category.
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1/02/2019 except for events not tracked.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitEmpty"`
	// Present on resources created after 1/02/2019 except for events not tracked.
	CreatedBy *CreatedBy      `json:"createdBy,omitEmpty"`
	Name      LocalizedString `json:"name"`
	// human-readable identifiers usually used as deep-link URL to the related category.
	// Each slug is unique across a project, but a category can have the same slug for different languages.
	Slug        LocalizedString  `json:"slug"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// Contains the parent path towards the root category.
	Ancestors []CategoryReference `json:"ancestors"`
	// A category that is the parent of this category in the category tree.
	Parent CategoryReference `json:"parent,omitEmpty"`
	// An attribute as base for a custom category order in one level.
	OrderHint       string           `json:"orderHint"`
	ExternalId      string           `json:"externalId,omitEmpty"`
	MetaTitle       *LocalizedString `json:"metaTitle,omitEmpty"`
	MetaDescription *LocalizedString `json:"metaDescription,omitEmpty"`
	MetaKeywords    *LocalizedString `json:"metaKeywords,omitEmpty"`
	Custom          *CustomFields    `json:"custom,omitEmpty"`
	// Can be used to store images, icons or movies related to this category.
	Assets []Asset `json:"assets,omitEmpty"`
	// User-specific unique identifier for the category.
	Key string `json:"key,omitEmpty"`
}

type CategoryDraft struct {
	Name LocalizedString `json:"name"`
	// human-readable identifier usually used as deep-link URL to the related category.
	// Allowed are alphabetic, numeric, underscore (`_`) and hyphen (`-`) characters.
	// Maximum size is 256.
	// **Must be unique across a project!** The same category can have the same slug for different languages.
	Slug        LocalizedString  `json:"slug"`
	Description *LocalizedString `json:"description,omitEmpty"`
	// A category that is the parent of this category in the category tree.
	// The parent can be set by its ID or by its key.
	Parent CategoryResourceIdentifier `json:"parent,omitEmpty"`
	// An attribute as base for a custom category order in one level.
	// A random value will be assigned by API if not set.
	OrderHint       string           `json:"orderHint,omitEmpty"`
	ExternalId      string           `json:"externalId,omitEmpty"`
	MetaTitle       *LocalizedString `json:"metaTitle,omitEmpty"`
	MetaDescription *LocalizedString `json:"metaDescription,omitEmpty"`
	MetaKeywords    *LocalizedString `json:"metaKeywords,omitEmpty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitEmpty"`
	Assets []AssetDraft       `json:"assets,omitEmpty"`
	// User-defined unique identifier for the category.
	// Keys can only contain alphanumeric characters (`a-Z, 0-9`), underscores and hyphens (`-, _`) and be between 2 and 256 characters.
	Key string `json:"key,omitEmpty"`
}

type CategoryPagedQueryResponse struct {
	Limit   int        `json:"limit"`
	Count   int        `json:"count"`
	Total   int        `json:"total,omitEmpty"`
	Offset  int        `json:"offset"`
	Results []Category `json:"results"`
}

type CategoryReference struct {
	Id  string    `json:"id"`
	Obj *Category `json:"obj,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "category", Alias: (*Alias)(&obj)})
}

type CategoryResourceIdentifier struct {
	Id  string `json:"id,omitEmpty"`
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
		new := CategoryAddAssetAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetName":
		new := CategoryChangeAssetNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeAssetOrder":
		new := CategoryChangeAssetOrderAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeName":
		new := CategoryChangeNameAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeOrderHint":
		new := CategoryChangeOrderHintAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeParent":
		new := CategoryChangeParentAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "changeSlug":
		new := CategoryChangeSlugAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "removeAsset":
		new := CategoryRemoveAssetAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomField":
		new := CategorySetAssetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetCustomType":
		new := CategorySetAssetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetDescription":
		new := CategorySetAssetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetKey":
		new := CategorySetAssetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetSources":
		new := CategorySetAssetSourcesAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setAssetTags":
		new := CategorySetAssetTagsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomField":
		new := CategorySetCustomFieldAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setCustomType":
		new := CategorySetCustomTypeAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setDescription":
		new := CategorySetDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setExternalId":
		new := CategorySetExternalIdAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setKey":
		new := CategorySetKeyAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaDescription":
		new := CategorySetMetaDescriptionAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaKeywords":
		new := CategorySetMetaKeywordsAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	case "setMetaTitle":
		new := CategorySetMetaTitleAction{}
		err := decodeStruct(input, &new)
		if err != nil {
			return nil, err
		}
		return new, nil
	}
	return nil, nil
}

type CategoryAddAssetAction struct {
	Asset AssetDraft `json:"asset"`
	// When specified, the value might be `0` and should be lower than the total of the assets list.
	Position int `json:"position,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

type CategoryChangeAssetNameAction struct {
	AssetId  string          `json:"assetId,omitEmpty"`
	AssetKey string          `json:"assetKey,omitEmpty"`
	Name     LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
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

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type CategoryRemoveAssetAction struct {
	AssetId  string `json:"assetId,omitEmpty"`
	AssetKey string `json:"assetKey,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomFieldAction struct {
	AssetId  string      `json:"assetId,omitEmpty"`
	AssetKey string      `json:"assetKey,omitEmpty"`
	Name     string      `json:"name"`
	Value    interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomTypeAction struct {
	AssetId  string `json:"assetId,omitEmpty"`
	AssetKey string `json:"assetKey,omitEmpty"`
	// If set, the custom type is set to this new value.
	// If absent, the custom type and any existing custom fields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// If set, the custom fields are set to this new value.
	Fields *interface{} `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

type CategorySetAssetDescriptionAction struct {
	AssetId     string           `json:"assetId,omitEmpty"`
	AssetKey    string           `json:"assetKey,omitEmpty"`
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	AssetKey string `json:"assetKey,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

type CategorySetAssetSourcesAction struct {
	AssetId  string        `json:"assetId,omitEmpty"`
	AssetKey string        `json:"assetKey,omitEmpty"`
	Sources  []AssetSource `json:"sources"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

type CategorySetAssetTagsAction struct {
	AssetId  string   `json:"assetId,omitEmpty"`
	AssetKey string   `json:"assetKey,omitEmpty"`
	Tags     []string `json:"tags,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

type CategorySetCustomFieldAction struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetCustomTypeAction struct {
	// If absent, the custom type and any existing CustomFields are removed.
	Type TypeResourceIdentifier `json:"type,omitEmpty"`
	// A valid JSON object, based on the FieldDefinitions of the Type. Sets the custom fields to this value.
	Fields *FieldContainer `json:"fields,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type CategorySetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CategorySetExternalIdAction struct {
	// If not defined, the external ID is unset.
	ExternalId string `json:"externalId,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
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
	Key string `json:"key,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type CategorySetMetaDescriptionAction struct {
	MetaDescription *LocalizedString `json:"metaDescription,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

type CategorySetMetaKeywordsAction struct {
	MetaKeywords *LocalizedString `json:"metaKeywords,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

type CategorySetMetaTitleAction struct {
	MetaTitle *LocalizedString `json:"metaTitle,omitEmpty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}
