// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

type Category struct {
	Version         int                 `json:"version"`
	LastModifiedAt  time.Time           `json:"lastModifiedAt"`
	ID              string              `json:"id"`
	CreatedAt       time.Time           `json:"createdAt"`
	Slug            *LocalizedString    `json:"slug"`
	Parent          *CategoryReference  `json:"parent,omitempty"`
	OrderHint       string              `json:"orderHint"`
	Name            *LocalizedString    `json:"name"`
	MetaTitle       *LocalizedString    `json:"metaTitle,omitempty"`
	MetaKeywords    *LocalizedString    `json:"metaKeywords,omitempty"`
	MetaDescription *LocalizedString    `json:"metaDescription,omitempty"`
	Key             string              `json:"key,omitempty"`
	ExternalID      string              `json:"externalId,omitempty"`
	Description     *LocalizedString    `json:"description,omitempty"`
	Custom          *CustomFields       `json:"custom,omitempty"`
	Assets          []Asset             `json:"assets,omitempty"`
	Ancestors       []CategoryReference `json:"ancestors"`
}

type CategoryAddAssetAction struct {
	Position float64     `json:"position,omitempty"`
	Asset    *AssetDraft `json:"asset"`
}

func (obj CategoryAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

type CategoryChangeAssetNameAction struct {
	Name     *LocalizedString `json:"name"`
	AssetKey string           `json:"assetKey,omitempty"`
	AssetID  string           `json:"assetId,omitempty"`
}

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

func (obj CategoryChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

type CategoryChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

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

func (obj CategoryChangeOrderHintAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeOrderHintAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeOrderHint", Alias: (*Alias)(&obj)})
}

type CategoryChangeParentAction struct {
	Parent *CategoryReference `json:"parent"`
}

func (obj CategoryChangeParentAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeParentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParent", Alias: (*Alias)(&obj)})
}

type CategoryChangeSlugAction struct {
	Slug *LocalizedString `json:"slug"`
}

func (obj CategoryChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

type CategoryDraft struct {
	Slug            *LocalizedString   `json:"slug"`
	Parent          *CategoryReference `json:"parent,omitempty"`
	OrderHint       string             `json:"orderHint,omitempty"`
	Name            *LocalizedString   `json:"name"`
	MetaTitle       *LocalizedString   `json:"metaTitle,omitempty"`
	MetaKeywords    *LocalizedString   `json:"metaKeywords,omitempty"`
	MetaDescription *LocalizedString   `json:"metaDescription,omitempty"`
	Key             string             `json:"key,omitempty"`
	ExternalID      string             `json:"externalId,omitempty"`
	Description     *LocalizedString   `json:"description,omitempty"`
	Custom          *CustomFieldsDraft `json:"custom,omitempty"`
	Assets          []AssetDraft       `json:"assets,omitempty"`
}

type CategoryPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
	Results []Category `json:"results"`
}

type CategoryReference struct {
	Key string    `json:"key,omitempty"`
	ID  string    `json:"id,omitempty"`
	Obj *Category `json:"obj,omitempty"`
}

func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "category", Alias: (*Alias)(&obj)})
}

type CategoryRemoveAssetAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId,omitempty"`
}

func (obj CategoryRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomFieldAction struct {
	Value    interface{} `json:"value,omitempty"`
	Name     string      `json:"name"`
	AssetKey string      `json:"assetKey,omitempty"`
	AssetID  string      `json:"assetId,omitempty"`
}

func (obj CategorySetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetAssetCustomTypeAction struct {
	Type     *TypeReference `json:"type,omitempty"`
	Fields   interface{}    `json:"fields,omitempty"`
	AssetKey string         `json:"assetKey,omitempty"`
	AssetID  string         `json:"assetId,omitempty"`
}

func (obj CategorySetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

type CategorySetAssetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
	AssetKey    string           `json:"assetKey,omitempty"`
	AssetID     string           `json:"assetId,omitempty"`
}

func (obj CategorySetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

type CategorySetAssetKeyAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId"`
}

func (obj CategorySetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

type CategorySetAssetSourcesAction struct {
	Sources  []AssetSource `json:"sources"`
	AssetKey string        `json:"assetKey,omitempty"`
	AssetID  string        `json:"assetId,omitempty"`
}

func (obj CategorySetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

type CategorySetAssetTagsAction struct {
	Tags     []string `json:"tags,omitempty"`
	AssetKey string   `json:"assetKey,omitempty"`
	AssetID  string   `json:"assetId,omitempty"`
}

func (obj CategorySetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

type CategorySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

func (obj CategorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type CategorySetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

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

func (obj CategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

type CategorySetExternalIdAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

func (obj CategorySetExternalIdAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetExternalIdAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

type CategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

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

func (obj CategorySetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

type CategoryUpdate struct {
	Version int                    `json:"version"`
	Actions []CategoryUpdateAction `json:"actions"`
}

func (obj *CategoryUpdate) UnmarshalJSON(data []byte) error {
	type Alias CategoryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		obj.Actions[i] = AbstractCategoryUpdateActionDiscriminatorMapping(obj.Actions[i])
	}

	return nil
}

type CategoryUpdateAction interface{}
type AbstractCategoryUpdateAction struct{}

func AbstractCategoryUpdateActionDiscriminatorMapping(input CategoryUpdateAction) CategoryUpdateAction {
	discriminator := input.(map[string]interface{})["action"]
	switch discriminator {
	case "addAsset":
		new := CategoryAddAssetAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAssetName":
		new := CategoryChangeAssetNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeAssetOrder":
		new := CategoryChangeAssetOrderAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeName":
		new := CategoryChangeNameAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeOrderHint":
		new := CategoryChangeOrderHintAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeParent":
		new := CategoryChangeParentAction{}
		mapstructure.Decode(input, &new)
		return new
	case "changeSlug":
		new := CategoryChangeSlugAction{}
		mapstructure.Decode(input, &new)
		return new
	case "removeAsset":
		new := CategoryRemoveAssetAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetCustomField":
		new := CategorySetAssetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetCustomType":
		new := CategorySetAssetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetDescription":
		new := CategorySetAssetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetKey":
		new := CategorySetAssetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetSources":
		new := CategorySetAssetSourcesAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setAssetTags":
		new := CategorySetAssetTagsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomField":
		new := CategorySetCustomFieldAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setCustomType":
		new := CategorySetCustomTypeAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setDescription":
		new := CategorySetDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setExternalId":
		new := CategorySetExternalIdAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setKey":
		new := CategorySetKeyAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaDescription":
		new := CategorySetMetaDescriptionAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaKeywords":
		new := CategorySetMetaKeywordsAction{}
		mapstructure.Decode(input, &new)
		return new
	case "setMetaTitle":
		new := CategorySetMetaTitleAction{}
		mapstructure.Decode(input, &new)
		return new
	}
	return nil
}
