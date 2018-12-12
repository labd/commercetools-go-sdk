// Automatically generated, do not edit

package commercetools

import (
	"encoding/json"
	"time"

	mapstructure "github.com/mitchellh/mapstructure"
)

// CategoryUpdateAction uses action as discriminator attribute
type CategoryUpdateAction interface{}

func mapDiscriminatorCategoryUpdateAction(input CategoryUpdateAction) CategoryUpdateAction {
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
		new := CategorySetExternalIDAction{}
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

// Category is of type Resource
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

// CategoryAddAssetAction implements the interface CategoryUpdateAction
type CategoryAddAssetAction struct {
	Position float64     `json:"position,omitempty"`
	Asset    *AssetDraft `json:"asset"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

// CategoryChangeAssetNameAction implements the interface CategoryUpdateAction
type CategoryChangeAssetNameAction struct {
	Name     *LocalizedString `json:"name"`
	AssetKey string           `json:"assetKey,omitempty"`
	AssetID  string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

// CategoryChangeAssetOrderAction implements the interface CategoryUpdateAction
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

// CategoryChangeNameAction implements the interface CategoryUpdateAction
type CategoryChangeNameAction struct {
	Name *LocalizedString `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeNameAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeName", Alias: (*Alias)(&obj)})
}

// CategoryChangeOrderHintAction implements the interface CategoryUpdateAction
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

// CategoryChangeParentAction implements the interface CategoryUpdateAction
type CategoryChangeParentAction struct {
	Parent *CategoryReference `json:"parent"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeParentAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeParentAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeParent", Alias: (*Alias)(&obj)})
}

// CategoryChangeSlugAction implements the interface CategoryUpdateAction
type CategoryChangeSlugAction struct {
	Slug *LocalizedString `json:"slug"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryChangeSlugAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryChangeSlugAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeSlug", Alias: (*Alias)(&obj)})
}

// CategoryDraft is a standalone struct
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

// CategoryPagedQueryResponse is of type PagedQueryResponse
type CategoryPagedQueryResponse struct {
	Total   int        `json:"total,omitempty"`
	Offset  int        `json:"offset"`
	Count   int        `json:"count"`
	Results []Category `json:"results"`
}

// CategoryReference implements the interface Reference
type CategoryReference struct {
	Key string    `json:"key,omitempty"`
	ID  string    `json:"id,omitempty"`
	Obj *Category `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryReference) MarshalJSON() ([]byte, error) {
	type Alias CategoryReference
	return json.Marshal(struct {
		TypeID string `json:"typeId"`
		*Alias
	}{TypeID: "category", Alias: (*Alias)(&obj)})
}

// CategoryRemoveAssetAction implements the interface CategoryUpdateAction
type CategoryRemoveAssetAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategoryRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias CategoryRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

// CategorySetAssetCustomFieldAction implements the interface CategoryUpdateAction
type CategorySetAssetCustomFieldAction struct {
	Value    interface{} `json:"value,omitempty"`
	Name     string      `json:"name"`
	AssetKey string      `json:"assetKey,omitempty"`
	AssetID  string      `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

// CategorySetAssetCustomTypeAction implements the interface CategoryUpdateAction
type CategorySetAssetCustomTypeAction struct {
	Type     *TypeReference `json:"type,omitempty"`
	Fields   interface{}    `json:"fields,omitempty"`
	AssetKey string         `json:"assetKey,omitempty"`
	AssetID  string         `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

// CategorySetAssetDescriptionAction implements the interface CategoryUpdateAction
type CategorySetAssetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
	AssetKey    string           `json:"assetKey,omitempty"`
	AssetID     string           `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

// CategorySetAssetKeyAction implements the interface CategoryUpdateAction
type CategorySetAssetKeyAction struct {
	AssetKey string `json:"assetKey,omitempty"`
	AssetID  string `json:"assetId"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

// CategorySetAssetSourcesAction implements the interface CategoryUpdateAction
type CategorySetAssetSourcesAction struct {
	Sources  []AssetSource `json:"sources"`
	AssetKey string        `json:"assetKey,omitempty"`
	AssetID  string        `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

// CategorySetAssetTagsAction implements the interface CategoryUpdateAction
type CategorySetAssetTagsAction struct {
	Tags     []string `json:"tags,omitempty"`
	AssetKey string   `json:"assetKey,omitempty"`
	AssetID  string   `json:"assetId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetAssetTagsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetTags", Alias: (*Alias)(&obj)})
}

// CategorySetCustomFieldAction implements the interface CategoryUpdateAction
type CategorySetCustomFieldAction struct {
	Value interface{} `json:"value,omitempty"`
	Name  string      `json:"name"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

// CategorySetCustomTypeAction implements the interface CategoryUpdateAction
type CategorySetCustomTypeAction struct {
	Type   *TypeReference  `json:"type,omitempty"`
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

// CategorySetDescriptionAction implements the interface CategoryUpdateAction
type CategorySetDescriptionAction struct {
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setDescription", Alias: (*Alias)(&obj)})
}

// CategorySetExternalIDAction implements the interface CategoryUpdateAction
type CategorySetExternalIDAction struct {
	ExternalID string `json:"externalId,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetExternalIDAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetExternalIDAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExternalId", Alias: (*Alias)(&obj)})
}

// CategorySetKeyAction implements the interface CategoryUpdateAction
type CategorySetKeyAction struct {
	Key string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

// CategorySetMetaDescriptionAction implements the interface CategoryUpdateAction
type CategorySetMetaDescriptionAction struct {
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaDescription", Alias: (*Alias)(&obj)})
}

// CategorySetMetaKeywordsAction implements the interface CategoryUpdateAction
type CategorySetMetaKeywordsAction struct {
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaKeywordsAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaKeywordsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaKeywords", Alias: (*Alias)(&obj)})
}

// CategorySetMetaTitleAction implements the interface CategoryUpdateAction
type CategorySetMetaTitleAction struct {
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
}

// MarshalJSON override to set the discriminator value
func (obj CategorySetMetaTitleAction) MarshalJSON() ([]byte, error) {
	type Alias CategorySetMetaTitleAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setMetaTitle", Alias: (*Alias)(&obj)})
}

// CategoryUpdate is of type Update
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
	for i := range obj.Actions {
		obj.Actions[i] = mapDiscriminatorCategoryUpdateAction(obj.Actions[i])
	}

	return nil
}
