package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type Category struct {
	// Unique identifier of the Category.
	ID string `json:"id"`
	// Current version of the Category.
	Version int `json:"version"`
	// Date and time (UTC) the Category was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Category was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources updated after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/../api/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// Name of the Category.
	Name LocalizedString `json:"name"`
	// User-defined identifier used as a deep-link URL to the related Category per [Locale](ctp:api:type:Locale).
	// A Category can have the same slug for different Locales, but they are unique across the [Project](ctp:api:type:Project).
	// Valid slugs match the pattern `^[A-Za-z0-9_-]{2,256}+$`.
	// For [good performance](/../api/predicates/query#performance-considerations), indexes are provided for the first 15 `languages` set in a Project.
	Slug LocalizedString `json:"slug"`
	// Description of the Category.
	Description *LocalizedString `json:"description,omitempty"`
	// Contains the parent path towards the root Category.
	Ancestors []CategoryReference `json:"ancestors"`
	// Parent Category of this Category.
	Parent *CategoryReference `json:"parent,omitempty"`
	// Decimal value between 0 and 1 used to order Categories that are on the same level in the Category tree.
	OrderHint string `json:"orderHint"`
	// Additional identifier for external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Name of the Category used by external search engines for improved search engine performance.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Description of the Category used by external search engines for improved search engine performance.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Keywords related to the Category for improved search engine performance.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// Custom Fields for the Category.
	Custom *CustomFields `json:"custom,omitempty"`
	// Media related to the Category.
	Assets []Asset `json:"assets"`
	// User-defined unique identifier of the Category.
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
	// Name of the Category.
	Name LocalizedString `json:"name"`
	// User-defined identifier used as a deep-link URL to the related Category.
	// A Category can have the same slug for different [Locales](ctp:api:type:Locale), but it must be unique across the [Project](ctp:api:type:Project).
	// Valid slugs must match the pattern `^[A-Za-z0-9_-]{2,256}+$`.
	Slug LocalizedString `json:"slug"`
	// Description of the Category.
	Description *LocalizedString `json:"description,omitempty"`
	// Parent Category of the Category.
	// The parent can be set by its `id` or `key`.
	Parent *CategoryResourceIdentifier `json:"parent,omitempty"`
	// Decimal value between 0 and 1 used to order Categories that are on the same level in the Category tree.
	// If not set, a random value will be assigned.
	OrderHint *string `json:"orderHint,omitempty"`
	// Additional identifier for external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
	ExternalId *string `json:"externalId,omitempty"`
	// Name of the Category used by external search engines for improved search engine performance.
	MetaTitle *LocalizedString `json:"metaTitle,omitempty"`
	// Description of the Category used by external search engines for improved search engine performance.
	MetaDescription *LocalizedString `json:"metaDescription,omitempty"`
	// Keywords related to the Category for improved search engine performance.
	MetaKeywords *LocalizedString `json:"metaKeywords,omitempty"`
	// Custom Fields for the Category.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
	// Media related to the Category.
	Assets []AssetDraft `json:"assets"`
	// User-defined unique identifier for the Category.
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

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Category](ctp:api:type:Category).
*
 */
type CategoryPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Category](ctp:api:type:Category) matching the query.
	Results []Category `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Category](ctp:api:type:Category).
*
 */
type CategoryReference struct {
	// Unique identifier of the referenced [Category](ctp:api:type:Category).
	ID string `json:"id"`
	// Contains the representation of the expanded Category. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Categories.
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

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Category](ctp:api:type:Category).
*
 */
type CategoryResourceIdentifier struct {
	// Unique identifier of the referenced [Category](ctp:api:type:Channel). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Category](ctp:api:type:Category). Either `id` or `key` is required.
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
	// Expected version of the Category on which the changes should be applied.
	// If the expected version does not match the actual version, a [409 Conflict](/../api/errors#409-conflict) will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Category.
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
		var err error
		obj.Actions[i], err = mapDiscriminatorCategoryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type CategoryUpdateAction interface{}

func mapDiscriminatorCategoryUpdateAction(input interface{}) (CategoryUpdateAction, error) {
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
	// Value to append.
	Asset AssetDraft `json:"asset"`
	// Position in the array at which the Asset should be put. When specified, the value must be between `0` and the total number of Assets minus `1`.
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// New value to set. Must not be empty.
	Name LocalizedString `json:"name"`
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

/**
*	This update action changes the order of the `assets` array. The new order is defined by listing the `id`s of the Assets.
*
 */
type CategoryChangeAssetOrderAction struct {
	// New value to set. Must contain all Asset `id`s.
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
	// New value to set. Must not be empty.
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
	// New value to set. Must be a decimal value between 0 and 1.
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
	// New value to set as parent.
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

/**
*	Changing the slug produces the [CategorySlugChangedMessage](ctp:api:type:CategorySlugChangedMessage).
*
 */
type CategoryChangeSlugAction struct {
	// New value to set. Must not be empty.
	// A Category can have the same slug for different [Locales](ctp:api:type:Locale), but it must be unique across the [Project](ctp:api:type:Project).
	// Valid slugs must match the pattern `^[A-Za-z0-9_-]{2,256}+$`.
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
	// Value to remove. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// Value to remove. Either `assetId` or `assetKey` is required.
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// Defines the [Type](ctp:api:type:Type) that extends the Asset with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Asset.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Asset.
	Fields *FieldContainer `json:"fields,omitempty"`
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set. If empty, any existing value will be removed.
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

/**
*	Set or remove the `key` of an [Asset](ctp:api:type:Asset).
*
 */
type CategorySetAssetKeyAction struct {
	// Value to set.
	AssetId string `json:"assetId"`
	// Value to set. If empty, any existing value will be removed.
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// Must not be empty. At least one entry is required.
	Sources []AssetSource `json:"sources"`
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
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetId *string `json:"assetId,omitempty"`
	// New value to set. Either `assetId` or `assetKey` is required.
	AssetKey *string `json:"assetKey,omitempty"`
	// Keywords for categorizing and organizing Assets.
	Tags []string `json:"tags"`
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
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
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
	// Defines the [Type](ctp:api:type:Type) that extends the Category with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the Category.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the Category.
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
	// Value to set. If empty, any existing value will be removed.
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

/**
*	This update action sets a new ID that can be used as an additional identifier for external systems like Customer Relationship Management (CRM) or Enterprise Resource Planning (ERP).
*
 */
type CategorySetExternalIdAction struct {
	// Value to set. If empty, any existing value will be removed.
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
	// Value to set. If empty, any existing value will be removed.
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
	// Value to set.
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
	// Value to set.
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
	// Value to set.
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
