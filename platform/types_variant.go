package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

/**
*	A Variant represents a single Variant with both current (published) and staged (draft) data.
*	Variants are writable resources that support create, update, and delete operations.
*
 */
type Variant struct {
	// Unique identifier of the Variant.
	ID string `json:"id"`
	// Current version of the Variant.
	Version int `json:"version"`
	// Date and time (UTC) the Variant was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// Date and time (UTC) the Variant was last updated.
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// IDs and references that last modified the Variant.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// IDs and references that created the Variant.
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the Variant.
	// This is different from [Product](ctp:api:type:Product) `key`.
	Key *string `json:"key,omitempty"`
	// A unique, sequential identifier of the Variant within the parent Product.
	VariantId int `json:"variantId"`
	// Reference to the parent [Product](ctp:api:type:Product) the Variant belongs to.
	Product ProductReference `json:"product"`
	// Whether the Variant is published, `false` if it is unpublished.
	Published bool `json:"published"`
	// The current data of the Variant.
	Current VariantData `json:"current"`
	// The staged data of the Variant.
	// Only present if there are staged changes that differ from the current data.
	Staged *VariantData `json:"staged,omitempty"`
	// Warnings about processing of a request.
	// Appears in response to requests with response status code `202 Accepted`.
	Warnings []WarningObject `json:"warnings"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *Variant) UnmarshalJSON(data []byte) error {
	type Alias Variant
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
func (obj Variant) MarshalJSON() ([]byte, error) {
	type Alias Variant
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
*	The request body for a bulk update of Variants.
*
 */
type VariantBulkUpdate struct {
	// List of Variants to update. Each item must identify the Variant by `id` or `key`.
	Items []VariantBulkUpdateItem `json:"items"`
	// The update action to apply to each Variant. Only [Publish](ctp:api:type:VariantPublishAction) and [Unpublish](ctp:api:type:VariantUnpublishAction) actions are supported.
	// An empty array results in a no-op: all items are returned as successful with no changes applied.
	Actions []VariantUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantBulkUpdate) UnmarshalJSON(data []byte) error {
	type Alias VariantBulkUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorVariantUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Identifies a single Variant to be updated in a bulk update request.
*	Provide either `id` or `key` to identify the Variant.
*
 */
type VariantBulkUpdateItem struct {
	// Unique identifier of the Variant to update. Either `id` or `key` must be provided.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the Variant to update. Either `id` or `key` must be provided.
	Key *string `json:"key,omitempty"`
	// Expected version of the Variant. Required when the `versionControl` query parameter is `On` (the default).
	// If the version does not match, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error is returned for that item.
	// Not required when `versionControl=Off`.
	Version *int `json:"version,omitempty"`
}

/**
*	A lightweight reference containing the `id` and updated `version` of a Variant after a successful bulk update.
*
 */
type VariantBulkUpdateResource struct {
	// Unique identifier of the updated Variant.
	ID string `json:"id"`
	// The new version of the Variant after the update.
	Version int `json:"version"`
}

/**
*	The response body for a bulk update of Variants. Contains the individual result for each Variant in the request.
*
 */
type VariantBulkUpdateResponse struct {
	// Number of Variants that were successfully updated.
	SuccessCount int `json:"successCount"`
	// Number of Variants that failed to update.
	FailureCount int `json:"failureCount"`
	// Results for each Variant in the request, in the same order as the `items` in the request body.
	Results []VariantBulkUpdateResult `json:"results"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantBulkUpdateResponse) UnmarshalJSON(data []byte) error {
	type Alias VariantBulkUpdateResponse
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Results {
		var err error
		obj.Results[i], err = mapDiscriminatorVariantBulkUpdateResult(obj.Results[i])
		if err != nil {
			return err
		}
	}

	return nil
}

/**
*	Represents the result of a single Variant update within a bulk update request.
*	Use `status` to determine whether the update succeeded or failed.
*
 */
type VariantBulkUpdateResult interface{}

func mapDiscriminatorVariantBulkUpdateResult(input interface{}) (VariantBulkUpdateResult, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["status"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'status'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "Fail":
		obj := VariantBulkUpdateFailResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		for i := range obj.Errors {
			var err error
			obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
			if err != nil {
				return nil, err
			}
		}
		return obj, nil
	case "Success":
		obj := VariantBulkUpdateSuccessResult{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Result for a Variant that failed to update in a bulk update request.
*
 */
type VariantBulkUpdateFailResult struct {
	// Unique identifier of the Variant from the request. Present when the Variant was identified by `id`.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the Variant from the request. Present when the Variant was identified by `key`.
	Key *string `json:"key,omitempty"`
	// Errors that caused the update to fail for this Variant.
	Errors []ErrorObject `json:"errors"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantBulkUpdateFailResult) UnmarshalJSON(data []byte) error {
	type Alias VariantBulkUpdateFailResult
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Errors {
		var err error
		obj.Errors[i], err = mapDiscriminatorErrorObject(obj.Errors[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantBulkUpdateFailResult) MarshalJSON() ([]byte, error) {
	type Alias VariantBulkUpdateFailResult
	return json.Marshal(struct {
		Action string `json:"status"`
		*Alias
	}{Action: "Fail", Alias: (*Alias)(&obj)})
}

/**
*	Result for a Variant that was successfully updated in a bulk update request.
*
 */
type VariantBulkUpdateSuccessResult struct {
	// Unique identifier of the Variant from the request. Present when the Variant was identified by `id`.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the Variant from the request. Present when the Variant was identified by `key`.
	Key *string `json:"key,omitempty"`
	// A lightweight reference containing the `id` and updated `version` of the Variant.
	Resource VariantBulkUpdateResource `json:"resource"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantBulkUpdateSuccessResult) MarshalJSON() ([]byte, error) {
	type Alias VariantBulkUpdateSuccessResult
	return json.Marshal(struct {
		Action string `json:"status"`
		*Alias
	}{Action: "Success", Alias: (*Alias)(&obj)})
}

/**
*	Contains the actual data for a Variant in either current (published) or staged (draft) state.
*
 */
type VariantData struct {
	// User-defined unique SKU of the Variant.
	Sku *string `json:"sku,omitempty"`
	// Images of the Variant.
	Images []Image `json:"images"`
	// Attributes of the Variant according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Attributes []Attribute `json:"attributes"`
	// Media assets of the Variant.
	Assets []Asset `json:"assets"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantData) MarshalJSON() ([]byte, error) {
	type Alias VariantData
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

/**
*	Draft for creating a new Variant. When a Variant is created, only the current version is set. The staged version is only set when the Variant is updated with different current and staged data.
*
 */
type VariantDraft struct {
	// User-defined unique identifier for the Variant.
	// This field is optional, but we strongly recommend setting it.
	Key *string `json:"key,omitempty"`
	// User-defined unique SKU of the Variant.
	Sku *string `json:"sku,omitempty"`
	// ResourceIdentifier of the [Product](ctp:api:type:Product) the Variant belongs to.
	Product ProductResourceIdentifier `json:"product"`
	// Indicates whether the Variant is published.
	// If `true`, the Variant will be immediately available in the current published state.
	Publish *bool `json:"publish,omitempty"`
	// Images for the Variant.
	Images []Image `json:"images"`
	// Variant Attributes according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Attributes []Attribute `json:"attributes"`
	// Media assets for the Variant.
	Assets []AssetDraft `json:"assets"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantDraft) MarshalJSON() ([]byte, error) {
	type Alias VariantDraft
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

/**
*	[PagedQueryResult](/../api/general-concepts#pagedqueryresult) with results containing an array of [Variant](ctp:api:type:Variant).
*
 */
type VariantPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// Unless absolutely necessary, use the query parameter `withTotal=false` to improve performance by deactivating the calculation of this field.
	// When the results are filtered with a [Query Predicate](ctp:api:type:QueryPredicate), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// [Variants](ctp:api:type:Variant) matching the query.
	Results []Variant `json:"results"`
}

/**
*	A lightweight, read-only projection of a single Variant with embedded Product data.
*	Variant Projections are automatically created and updated when Variants or their parent Products change.
*
 */
type VariantProjection struct {
	// Unique identifier of the Variant within its parent Product.
	ID string `json:"id"`
	// Current version of the Variant Projection.
	Version int `json:"version"`
	// Date and time (UTC) the Variant Projection was initially created.
	CreatedAt time.Time `json:"createdAt"`
	// `true` for the staged (draft) projection, `false` for the current (published) projection.
	Staged bool `json:"staged"`
	// The `id` of the [Variant](ctp:api:type:Variant).
	VariantId int `json:"variantId"`
	// [Reference](ctp:api:type:Reference) to the parent [Product](ctp:api:type:Product).
	Product ProductReference `json:"product"`
	// Name of the parent [Product](ctp:api:type:Product).
	Name LocalizedString `json:"name"`
	// Slug of the parent [Product](ctp:api:type:Product).
	Slug LocalizedString `json:"slug"`
	// Description of the parent [Product](ctp:api:type:Product).
	Description *LocalizedString `json:"description,omitempty"`
	// [Categories](ctp:api:type:Category) assigned to the parent [Product](ctp:api:type:Product).
	Categories []CategoryReference `json:"categories"`
	// Order of the parent [Product](ctp:api:type:Product) in [Categories](ctp:api:type:Category).
	CategoryOrderHints *CategoryOrderHints `json:"categoryOrderHints,omitempty"`
	// User-defined unique identifier of the [Variant](ctp:api:type:Variant).
	Key *string `json:"key,omitempty"`
	// SKU of the [Variant](ctp:api:type:Variant).
	Sku *string `json:"sku,omitempty"`
	// Images of the [Variant](ctp:api:type:Variant).
	Images []Image `json:"images"`
	// Assets of the [Variant](ctp:api:type:Variant).
	Assets []Asset `json:"assets"`
	// Attributes of the [Variant](ctp:api:type:Variant), including product-level Attributes merged at projection time.
	Attributes []Attribute `json:"attributes"`
	// The selected price based on the [price selection](/../api/projects/productProjections#prices) query parameters.
	// Only present when price selection parameters are provided.
	Price *Price `json:"price,omitempty"`
	// Whether this Variant is the default Variant of its [Product](ctp:api:type:Product) (see [Product.defaultVariant](/projects/products#product)).
	Default bool `json:"default"`
}

type VariantProjectionPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int `json:"limit"`
	// Actual number of results returned.
	Count int `json:"count"`
	// Total number of results matching the query.
	// This number is an estimation that is not [strongly consistent](/../api/general-concepts#strong-consistency).
	// This field is returned by default.
	// For improved performance, calculating this field can be deactivated by using the query parameter `withTotal=false`.
	// When the results are filtered with a [Query Predicate](/../api/predicates/query), `total` is subject to a [limit](/../api/limits#queries).
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset int `json:"offset"`
	// [VariantProjections](ctp:api:type:VariantProjection) matching the query.
	Results []VariantProjection `json:"results"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [Variant](ctp:api:type:Variant).
*
 */
type VariantReference struct {
	// Unique identifier of the referenced [Variant](ctp:api:type:Variant).
	ID string `json:"id"`
	// Contains the representation of the expanded Variant. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for Variants.
	Obj *Variant `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantReference) MarshalJSON() ([]byte, error) {
	type Alias VariantReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "variant", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [Variant](ctp:api:type:Variant). Either `id` or `key` is required. If both are set, an [InvalidJsonInput](/../api/errors#invalidjsoninput) error is returned.
*
 */
type VariantResourceIdentifier struct {
	// Unique identifier of the referenced [Variant](ctp:api:type:Variant). Required if `key` is absent.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [Variant](ctp:api:type:Variant). Required if `id` is absent.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias VariantResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "variant", Alias: (*Alias)(&obj)})
}

type VariantUpdate struct {
	// Expected version of the Variant on which the changes should be applied. If the expected version does not match the actual version, a [ConcurrentModification](ctp:api:type:ConcurrentModificationError) error will be returned.
	Version int `json:"version"`
	// Update actions to be performed on the Variant.
	Actions []VariantUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *VariantUpdate) UnmarshalJSON(data []byte) error {
	type Alias VariantUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorVariantUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}

	return nil
}

type VariantUpdateAction interface{}

func mapDiscriminatorVariantUpdateAction(input interface{}) (VariantUpdateAction, error) {
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
		obj := VariantAddAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "addExternalImage":
		obj := VariantAddExternalImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetName":
		obj := VariantChangeAssetNameAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeAssetOrder":
		obj := VariantChangeAssetOrderAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "moveImageToPosition":
		obj := VariantMoveImageToPositionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "publish":
		obj := VariantPublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeAsset":
		obj := VariantRemoveAssetAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeImage":
		obj := VariantRemoveImageAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeStagedChanges":
		obj := VariantRemoveStagedChangesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomField":
		obj := VariantSetAssetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetCustomType":
		obj := VariantSetAssetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetDescription":
		obj := VariantSetAssetDescriptionAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetKey":
		obj := VariantSetAssetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetSources":
		obj := VariantSetAssetSourcesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssetTags":
		obj := VariantSetAssetTagsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAssets":
		obj := VariantSetAssetsAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttribute":
		obj := VariantSetAttributeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setAttributes":
		obj := VariantSetAttributesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImageLabel":
		obj := VariantSetImageLabelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setImages":
		obj := VariantSetImagesAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := VariantSetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSku":
		obj := VariantSetSkuAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "unpublish":
		obj := VariantUnpublishAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

/**
*	Adds an asset to existing variant's assets.
*
 */
type VariantAddAssetAction struct {
	// Whether only the staged `assets` are updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// Value to append.
	Asset AssetDraft `json:"asset"`
	// Position in `assets` where the Asset should be put. When specified, the value must be between `0` and the total number of Assets minus `1`.
	Position *int `json:"position,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantAddAssetAction) MarshalJSON() ([]byte, error) {
	type Alias VariantAddAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addAsset", Alias: (*Alias)(&obj)})
}

/**
*	Adds an image to existing variant's images.
*	Generates the [VariantImageAdded](ctp:api:type:VariantImageAddedMessage) Message.
*
 */
type VariantAddExternalImageAction struct {
	// Value to add to `images`.
	Image Image `json:"image"`
	// Whether only the staged `images` is updated. If `false`, both the current and staged `images` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantAddExternalImageAction) MarshalJSON() ([]byte, error) {
	type Alias VariantAddExternalImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addExternalImage", Alias: (*Alias)(&obj)})
}

/**
*	Changes the name of an existing Asset on a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantChangeAssetNameAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
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
func (obj VariantChangeAssetNameAction) MarshalJSON() ([]byte, error) {
	type Alias VariantChangeAssetNameAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetName", Alias: (*Alias)(&obj)})
}

/**
*	Changes the order of Assets on a Variant.
*
 */
type VariantChangeAssetOrderAction struct {
	// Whether only the staged `assets` is updated. If `false`, both the current and staged `assets` are updated.
	Staged *bool `json:"staged,omitempty"`
	// All existing Asset `id`s of the Variant in the desired new order.
	AssetOrder []string `json:"assetOrder"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantChangeAssetOrderAction) MarshalJSON() ([]byte, error) {
	type Alias VariantChangeAssetOrderAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeAssetOrder", Alias: (*Alias)(&obj)})
}

/**
*	Moves a Variant's image to a new position.
*
 */
type VariantMoveImageToPositionAction struct {
	// The URL of the image to update.
	ImageUrl string `json:"imageUrl"`
	// Position in `images` where the image should be moved. Must be between `0` and the total number of images minus `1`.
	Position int `json:"position"`
	// If `true`, only the staged `images` is updated. If `false`, both the current and staged `images` is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantMoveImageToPositionAction) MarshalJSON() ([]byte, error) {
	type Alias VariantMoveImageToPositionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "moveImageToPosition", Alias: (*Alias)(&obj)})
}

/**
*	Publishes the Variant by copying the staged data to the current data and setting the `published` flag to `true`. Staged will be empty.
*	The parent Product must be published before a Variant can be published.
*	Generates the [VariantPublished](ctp:api:type:VariantPublishedMessage) Message.
*
 */
type VariantPublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantPublishAction) MarshalJSON() ([]byte, error) {
	type Alias VariantPublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "publish", Alias: (*Alias)(&obj)})
}

/**
*	Removes an asset from an existing Variant's assets. The asset to remove must be specified using either `assetId` or `assetKey`.
*
 */
type VariantRemoveAssetAction struct {
	// Whether only the staged Asset is removed. If `false`, both the current and staged Asset is removed.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to remove.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to remove.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantRemoveAssetAction) MarshalJSON() ([]byte, error) {
	type Alias VariantRemoveAssetAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeAsset", Alias: (*Alias)(&obj)})
}

/**
*	Removes a Variant's image.
*
 */
type VariantRemoveImageAction struct {
	// The URL of the image to remove.
	ImageUrl string `json:"imageUrl"`
	// Whether only the staged image is removed. If `false`, both the current and staged image is removed.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantRemoveImageAction) MarshalJSON() ([]byte, error) {
	type Alias VariantRemoveImageAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeImage", Alias: (*Alias)(&obj)})
}

/**
*	Removes all staged changes from the Variant, only the current data remains.
*	Generates the [VariantStagedChangesRemoved](ctp:api:type:VariantStagedChangesRemovedMessage) Message.
*
 */
type VariantRemoveStagedChangesAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantRemoveStagedChangesAction) MarshalJSON() ([]byte, error) {
	type Alias VariantRemoveStagedChangesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeStagedChanges", Alias: (*Alias)(&obj)})
}

/**
*	Sets or removes a Custom Field on an existing Asset of a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantSetAssetCustomFieldAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
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
func (obj VariantSetAssetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomField", Alias: (*Alias)(&obj)})
}

/**
*	Sets or removes the Custom Type on an existing Asset of a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantSetAssetCustomTypeAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
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
func (obj VariantSetAssetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetCustomType", Alias: (*Alias)(&obj)})
}

/**
*	Sets the description of an existing Asset on a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantSetAssetDescriptionAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId *string `json:"assetId,omitempty"`
	// The `key` of the Asset to update.
	AssetKey *string `json:"assetKey,omitempty"`
	// Value to set. If omitted, any existing value is removed.
	Description *LocalizedString `json:"description,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetAssetDescriptionAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetDescriptionAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetDescription", Alias: (*Alias)(&obj)})
}

/**
*	Sets the key of an existing Asset on a Variant.
*
 */
type VariantSetAssetKeyAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
	Staged *bool `json:"staged,omitempty"`
	// The `id` of the Asset to update.
	AssetId string `json:"assetId"`
	// Value to set. If omitted, any existing value is removed.
	AssetKey *string `json:"assetKey,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetAssetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetKey", Alias: (*Alias)(&obj)})
}

/**
*	Sets the sources of an existing Asset on a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantSetAssetSourcesAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
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
func (obj VariantSetAssetSourcesAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetSourcesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssetSources", Alias: (*Alias)(&obj)})
}

/**
*	Sets the tags of an existing Asset on a Variant. The Asset to update must be specified using either `assetId` or `assetKey`.
*
 */
type VariantSetAssetTagsAction struct {
	// Whether only the staged Asset is updated. If `false`, both the current and staged Asset is updated.
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
func (obj VariantSetAssetTagsAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetTagsAction
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
*	Sets the assets of the Variant. Existing assets will be replaced with the new ones.
*
 */
type VariantSetAssetsAction struct {
	// Assets to set for the Variant.
	Assets []AssetDraft `json:"assets"`
	// Whether only the staged assets are set. If `false`, both the current and staged assets are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetAssetsAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAssetsAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAssets", Alias: (*Alias)(&obj)})
}

/**
*	Set a specific Variant attribute.
*
 */
type VariantSetAttributeAction struct {
	// Name of the Attribute to set.
	Name string `json:"name"`
	// Value to set for the Attribute. If omitted, any existing value is removed.
	//
	// The [AttributeType](ctp:api:type:AttributeType) determines the format of the Attribute `value` to be provided:
	//
	// - For [Enum Type](ctp:api:type:AttributeEnumType) and [Localized Enum Type](ctp:api:type:AttributeLocalizedEnumType), use the `key` of the [Plain Enum Value](ctp:api:type:AttributePlainEnumValue) or [Localized Enum Value](ctp:api:type:AttributeLocalizedEnumValue) object or the complete object as `value`.
	// - For [Localizable Text Type](ctp:api:type:AttributeLocalizableTextType), use the [LocalizedString](ctp:api:type:LocalizedString) object as `value`.
	// - For [Money Type](ctp:api:type:AttributeMoneyType) Attributes, use the [Money](ctp:api:type:Money) object as `value`.
	// - For [Set Type](ctp:api:type:AttributeSetType) Attributes, use the entire `set` object  as `value`.
	// - For [Nested Type](ctp:api:type:AttributeNestedType) Attributes, use the list of values of all Attributes of the nested Variant as `value`.
	// - For [Reference Type](ctp:api:type:AttributeReferenceType) Attributes, use the [Reference](ctp:api:type:Reference) object as `value`.
	Value interface{} `json:"value,omitempty"`
	// Whether only the staged Attribute is set. If `false`, both the current and staged Attributes are set.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetAttributeAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAttributeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttribute", Alias: (*Alias)(&obj)})
}

/**
*	Sets the attributes of the Variant. Existing attributes will be replaced with the new ones.
*
 */
type VariantSetAttributesAction struct {
	// Attributes to set for the Variant according to the respective [AttributeDefinition](ctp:api:type:AttributeDefinition).
	Attributes []Attribute `json:"attributes"`
	// Whether only the staged attributes are updated. If `false`, both the current and staged attributes are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetAttributesAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetAttributesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setAttributes", Alias: (*Alias)(&obj)})
}

/**
*	Sets the label of a Variant's image.
*
 */
type VariantSetImageLabelAction struct {
	// The URL of the image to set the label.
	ImageUrl string `json:"imageUrl"`
	// Value to set. If omitted, any existing value is removed.
	Label *string `json:"label,omitempty"`
	// If `true`, only the staged image is updated. If `false`, both the current and staged image is updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetImageLabelAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetImageLabelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImageLabel", Alias: (*Alias)(&obj)})
}

/**
*	Sets the images of the Variant. Existing images will be replaced with the new ones.
*	Generates the [VariantImagesSet](ctp:api:type:VariantImagesSetMessage) Message.
*
 */
type VariantSetImagesAction struct {
	// Images to set for the Variant.
	Images []Image `json:"images"`
	// Whether only the staged images are updated. If `false`, both the current and staged images are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetImagesAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetImagesAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setImages", Alias: (*Alias)(&obj)})
}

/**
*	Sets the key of the Variant.
*	Generates the [VariantKeySet](ctp:api:type:VariantKeySetMessage) Message.
*
 */
type VariantSetKeyAction struct {
	// Value to set. Must be unique.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

/**
*	Sets the SKU of the Variant.
*	SKU cannot be changed or removed if it is associated with an [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type VariantSetSkuAction struct {
	// Value to set. Must be unique.
	Sku *string `json:"sku,omitempty"`
	// Whether only the staged `sku` is updated. If `false`, both the current and staged `sku` are updated.
	Staged *bool `json:"staged,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantSetSkuAction) MarshalJSON() ([]byte, error) {
	type Alias VariantSetSkuAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSku", Alias: (*Alias)(&obj)})
}

/**
*	Unpublishes the Variant by setting the `published` flag to `false`.
*	Generates the [VariantUnpublished](ctp:api:type:VariantUnpublishedMessage) Message.
*
 */
type VariantUnpublishAction struct {
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj VariantUnpublishAction) MarshalJSON() ([]byte, error) {
	type Alias VariantUnpublishAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "unpublish", Alias: (*Alias)(&obj)})
}
